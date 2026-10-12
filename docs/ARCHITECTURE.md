# sdk-ipad 架构文档

## 1. 概述

sdk-ipad 是 Ocean SCRM 系统的核心包，负责：
- 封装企业微信 iPad 协议的 HTTP API 调用（14个业务模块）
- 处理协议服务器的多账号高并发回调消息（8个回调模块）

新架构采用 **分片管理 + 共享 Worker 池 + Hook 机制** 的设计，解决原有单管道架构在多账号（1000+）高并发回调时的性能瓶颈和内存扩展性问题。

### 核心设计原则

1. **内存高效**：10000 账号场景下内存 < 2GB
2. **性能隔离**：热点账号优先处理，不影响其他账号
3. **线性扩展**：分片架构支持水平扩展
4. **向后兼容**：无缝迁移，业务逻辑无需修改

---

## 2. 架构组件

### 2.1 SdkIPad（SDK 核心）

- **文件**：`sdk.go`
- **职责**：SDK 入口，持有 HTTP 客户端和分片管理器
- **配置**：`SdkConfig`（含 ShardConfig、WorkerPoolConfig、PriorityQueueConfig）

```go
type SdkIPad struct {
    httpClient   HTTPClient
    shardManager *ShardManager
    config       *SdkConfig
}
```

### 2.2 ShardManager（分片管理器）

- **文件**：`shard_manager.go`
- **职责**：按账号 UUID 哈希分片，管理轻量账号上下文（~100 字节/账号）
- **特点**：支持懒加载、TTL 自动清理、可优雅停止

```go
type ShardManager struct {
    shards     []*Shard          // 固定数量的分片
    shardCount int
    hashFunc   func(string) int  // 默认 FNV-32a
    stopChan   chan struct{}      // 用于停止 cleanup goroutine
    config     ShardConfig
}
```

### 2.3 SharedWorkerPool（共享 Worker 池）

- **文件**：`shared_pool.go`
- **职责**：全局任务处理池，动态扩缩容
- **特点**：按队列深度自动调整 Worker 数量，原子统计，幂等停止

```go
type SharedWorkerPool struct {
    config    WorkerPoolConfig
    taskChan  chan *Task
    workers   []*worker
    mu        sync.Mutex
    isRunning atomic.Bool
    stopOnce  sync.Once
    stopChan  chan struct{}
    stats     PoolStats
}
```

### 2.4 CallbackDispatcher（回调分发器）

- **文件**：`callback_dispatcher.go`
- **职责**：按消息类型分发到注册的 Hook 处理器
- **特点**：消息优先级、原子统计、线程安全 Hook 注册

```go
type CallbackDispatcher struct {
    shardManager *ShardManager
    workerPool   *SharedWorkerPool
    hookRegistry *HookRegistry
    config       DispatcherConfig
    stats        DispatcherStats   // 全原子字段
    isRunning    atomic.Bool
    stopOnce     sync.Once
    sdk          *SdkIPad
}
```

### 2.5 DispatcherAdapter（分发器适配器）

- **文件**：`dispatcher_adapter.go`
- **职责**：连接现有业务代码和新架构的适配层
- **特点**：提供全局单例、管理完整生命周期（含 WorkerPool）

```go
type DispatcherAdapter struct {
    dispatcher *CallbackDispatcher
    workerPool *SharedWorkerPool
    sdk        *SdkIPad
}
```

### 2.6 Hook 机制

- **文件**：`example_hooks.go`、`callback_dispatcher.go`
- **职责**：业务逻辑处理接口
- **特点**：按优先级执行、线程安全注册/注销、支持账号级上下文

```go
type Hook interface {
    Handle(ctx *HookContext, msg *msg_struct.CallBackMsg) error
    Priority() int
    Name() string
}

type HookContext struct {
    AccountCtx *AccountContext
    SDK        *SdkIPad
    StartTime  time.Time
}
```

---

## 3. 消息处理流程

### 3.1 流程图

```
企业微信服务器
     │  HTTP 回调
     ▼
/callback 接口（解析 JSON → CallBackMsg）
     │
     ▼
DispatcherAdapter.HandleCallback()
     │
     ▼
CallbackDispatcher.HandleCallback()
     │  ┌─ 确定消息优先级（0/1/2）
     │  └─ 创建 Task
     ▼
SharedWorkerPool.Submit(task)
     │
     ▼
Worker 取出 Task → callbackTaskHandler.Handle()
     │  ┌─ ShardManager.GetOrCreateAccount(uuid)
     │  └─ 获取 AccountContext
     ▼
CallbackDispatcher.processCallback()
     │  ┌─ 查找 HookRegistry[msg.Type]
     │  └─ 复制 hooks 切片（避免迭代竞态）
     ▼
按优先级依次执行 Hook.Handle()
     │
     ▼
业务处理完成，更新统计
```

### 3.2 优先级策略

| 优先级 | 消息类型 | 说明 |
|--------|----------|------|
| **高 (2)** | `ScanReturn`, `ConfirmReturn`, `LoginSuccessReturn`, `Exception`, `LoginInOtherDevice` | 登录/断开，需即时处理 |
| **中 (1)** | `ApplyFriend`, `NewGroup`, `RemoveGroupMember`, `GroupDisBand` | 重要业务，时效性较强 |
| **低 (0)** | 其他所有消息类型 | 普通消息，可适当延迟 |

---

## 4. 调用入口

sdk-ipad 提供两个主要入口，适用于不同场景：

### 4.1 DispatcherAdapter 入口（推荐）

适合大多数现有业务场景，提供开箱即用的全局单例。

```go
// 方式1：全局便捷函数（最常用）
err = sdk_ipad.HandleCallbackGlobal(&msg)

// 方式2：获取默认适配器
adapter := sdk_ipad.GetDefaultAdapter()
err = adapter.HandleCallback(&msg)

// 方式3：自定义配置
config := sdk_ipad.DefaultDispatcherConfig()
config.MaxConcurrent = 200
sdk := sdk_ipad.NewSdkIPad()
adapter := sdk_ipad.NewDispatcherAdapter(sdk, config)
adapter.Start()
defer adapter.Stop()  // Stop 会同时停止 Dispatcher 和 WorkerPool
```

**优点**：简单、安全、完整（内置所有组件）、带监控统计
**限制**：配置相对固定、所有业务共享同一 Worker 池

### 4.2 SdkIPad 入口（高级）

适合需要自定义配置、资源隔离或精细控制的场景。

```go
// 自定义配置
config := &sdk_ipad.SdkConfig{
    ShardConfig: sdk_ipad.ShardConfig{
        ShardCount:  512,
        MaxAccounts: 20000,
    },
    WorkerPool: sdk_ipad.WorkerPoolConfig{
        MinWorkers:    16,
        MaxWorkers:    64,
        QueueCapacity: 50000,
    },
}
sdk := sdk_ipad.NewSdkIPadWithConfig(config)

// 直接使用 SDK 处理回调
err = sdk.HandleCallback(&msg, func(ctx *sdk_ipad.AccountContext, msg *msg_struct.CallBackMsg) error {
    return processMessage(ctx, msg)
})
```

**优点**：灵活、可隔离资源、支持自定义算法
**限制**：需手动管理组件生命周期

### 4.3 选择指南

```
现有代码迁移 / 标准业务 / 单应用 → DispatcherAdapter
多租户隔离 / 性能调优 / 自定义策略 → SdkIPad
```

升级路径：`HandleCallbackGlobal → 自定义 Adapter → 独立 SdkIPad → 完全手动组装`

---

## 5. Hook 注册与使用

### 5.1 实现自定义 Hook

```go
type MyBusinessHook struct {
    name     string
    priority int
}

func (h *MyBusinessHook) Handle(ctx *sdk_ipad.HookContext, msg *msg_struct.CallBackMsg) error {
    // ctx.SDK 可访问 SDK 实例
    // ctx.AccountCtx 可访问账号上下文
    switch msg.Type {
    case msg_struct.MsgCallback:
        return h.handleMessage(ctx, msg)
    default:
        return nil
    }
}

func (h *MyBusinessHook) Priority() int { return h.priority }
func (h *MyBusinessHook) Name() string  { return h.name }
```

### 5.2 注册 Hook

```go
adapter := sdk_ipad.GetDefaultAdapter()
hook := &MyBusinessHook{name: "biz-hook", priority: 1}
adapter.RegisterHook(msg_struct.MsgCallback, hook)
adapter.RegisterHook(msg_struct.ApplyFriend, hook)
```

### 5.3 HookManager（批量管理）

```go
hm := sdk_ipad.NewHookManager()
hm.RegisterHook(hook1)
hm.RegisterHook(hook2)
hm.RegisterAllToDispatcher(dispatcher, []int{msg_struct.MsgCallback, msg_struct.ApplyFriend})
```

### 5.4 AccountHook（账号级生命周期）

```go
type AccountHook interface {
    Hook
    InitAccount(uuid string, sdk *SdkIPad) error
    CleanupAccount(uuid string) error
}
```

---

## 6. 配置参数

### 6.1 ShardConfig

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `ShardCount` | 256 | 分片数量 |
| `MaxAccounts` | 5000 | 最大账号数 |
| `AccountTTL` | 24h | 账号 TTL |
| `CleanupInterval` | 1h | 清理间隔 |

### 6.2 WorkerPoolConfig

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `MinWorkers` | 4 | 最小 Worker 数 |
| `MaxWorkers` | 16 | 最大 Worker 数 |
| `QueueCapacity` | 10000 | 任务队列容量 |
| `WorkerIdleTime` | 30s | Worker 空闲超时 |
| `ScaleUpInterval` | 1s | 扩容检查间隔 |

### 6.3 DispatcherConfig

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `MaxConcurrent` | 100 | 最大并发处理数 |
| `Timeout` | 5s | 处理超时 |
| `RetryCount` | 3 | 重试次数 |
| `RetryInterval` | 100ms | 重试间隔 |
| `DeadLetterQueue` | true | 是否启用死信队列 |

### 6.4 场景配置建议

```go
// 小型（< 1000 账号）：使用默认配置
adapter := sdk_ipad.GetDefaultAdapter()

// 中型（1000-5000 账号）
config.MaxConcurrent = 200
workerConfig.MinWorkers = 8
workerConfig.MaxWorkers = 32

// 大型（5000+ 账号）
shardConfig.ShardCount = 512
shardConfig.AccountTTL = 12 * time.Hour
workerConfig.MinWorkers = 16
workerConfig.MaxWorkers = 64
workerConfig.QueueCapacity = 50000
```

---

## 7. 监控与统计

### 7.1 获取统计

```go
adapter := sdk_ipad.GetDefaultAdapter()

// 分发器统计
stats := adapter.GetStats()
// stats.TotalReceived / TotalProcessed / TotalFailed（原子安全）

// 分片统计
shardStats := adapter.GetShardStats(0)
allStats := adapter.GetAllShardStats()

// Worker 池统计
poolStats := adapter.GetWorkerPoolStats()
```

### 7.2 关键监控指标

| 指标 | 说明 | 告警阈值 |
|------|------|----------|
| 队列深度 | WorkerPool 队列长度 | > 80% 容量 |
| 处理延迟 | P99 消息处理延迟 | > 500ms |
| 错误率 | 处理失败比例 | > 1% |
| 内存使用 | 分片管理器内存 | > 2GB |
| 活跃账号 | 5 分钟内有消息的账号数 | 按业务需求 |

---

## 8. 并发安全设计

本架构在以下关键点保证线程安全：

| 组件 | 机制 | 说明 |
|------|------|------|
| ShardManager | 每分片独立 RWMutex | 分片间无锁竞争 |
| ShardManager.cleanupShard | 拆分 Locked 版本 | 避免重入死锁 |
| SharedWorkerPool | atomic.Bool + sync.Once | 幂等 Stop，无 double-close |
| Worker.stopChan | sync.Once 保护 close | 防止 removeIdleWorker 与 Stop 冲突 |
| CallbackDispatcher | atomic.Bool + sync.Once | Start/Stop 无竞态 |
| DispatcherStats | 全 int64 + atomic 操作 | HandleCallback 并发安全 |
| HookRegistry | RWMutex + copy-on-read | 注册/迭代无竞态 |
| HookManager | sync.RWMutex | 并发注册/查询安全 |
| ExampleAccountHook | sync.RWMutex 保护 map | 并发 Init/Cleanup 安全 |

---

## 9. 生命周期管理

```go
// 启动
sdk := sdk_ipad.NewSdkIPad()
adapter := sdk_ipad.NewDispatcherAdapter(sdk, config)
adapter.Start()

// 运行期间
adapter.HandleCallback(&msg)

// 停止（级联关闭）
adapter.Stop()
//   └─ CallbackDispatcher.Stop()  (sync.Once 幂等)
//   └─ SharedWorkerPool.Stop()    (sync.Once 幂等，安全关闭所有 Worker)

// ShardManager 独立停止（如需要）
sdk.GetShardManager().Stop()  // 终止 cleanup goroutine
```

---

## 10. 迁移指南

### 10.1 快速迁移（一行修改）

```go
// 修改前
h.service.CallbackMsgChan <- &msg

// 修改后
err = sdk_ipad.HandleCallbackGlobal(&msg)
```

### 10.2 渐进式迁移策略

1. **影子模式**：部署新架构，并行记录新旧处理结果对比
2. **小流量切换**：1% → 5% → 20% → 50% 逐步切换
3. **全面切换**：100% 流量，监控 24 小时
4. **清理优化**：移除旧代码，根据负载调优

### 10.3 回滚

```go
// 将 HandleCallbackGlobal 替换回原有 Channel 写入即可
h.service.CallbackMsgChan <- &msg
```

---

## 11. 故障排查

| 问题 | 排查方向 |
|------|----------|
| 消息延迟高 | 检查 WorkerPool 队列深度；增加 MaxWorkers；检查 Hook 是否阻塞 |
| 内存过高 | 缩短 AccountTTL；检查 Hook 内存泄漏；监控分片负载均衡 |
| 消息丢失 | 检查死信队列；验证重试机制；检查超时配置 |
| 分片不均衡 | 检查哈希函数分布；监控各分片负载 |

---

## 12. 性能基准

| 场景 | 账号数 | 消息频率 | 内存使用 | P99 延迟 |
|------|--------|----------|----------|----------|
| 基准 | 1,000 | 100 msg/s | ~200MB | < 50ms |
| 中等 | 5,000 | 500 msg/s | ~800MB | < 100ms |
| 高负载 | 10,000 | 1,000 msg/s | ~1.5GB | < 200ms |
| 压力 | 20,000 | 2,000 msg/s | ~2.8GB | < 500ms |

内存估算：`(账号数 × 100B) + (分片数 × 1KB) + (队列容量 × 8B × 优先级数) + (Worker数 × 2MB)`

---

## 13. 文件清单

| 文件 | 职责 |
|------|------|
| `sdk.go` | SDK 入口、SdkConfig、SdkIPad 核心结构、包级接口定位指引 |
| `sdk_interface.go` | SdkInterface 接口（业务侧使用的方法集合，供 mock/替换） |
| `shard_manager.go` | 分片管理器、AccountContext、TTL 清理 |
| `shared_pool.go` | 共享 Worker 池、动态扩缩容、Task 接口 |
| `callback_dispatcher.go` | 回调分发器、Hook 接口/注册表、HookContext |
| `dispatcher_adapter.go` | 适配器层、全局单例、HandleCallbackGlobal |
| `callback_base.go` | 回调基础层：CallBackMsg 信封、回调类型码常量、SendMsg |
| `callback_http.go` | 回调 HTTP 传输层适配（CallbackHTTPHandler） |
| `callback_payload_msg.go` | 消息回调（102000）契约层：信封与 12 类消息访问器 |
| `callback_payload_account.go` | 登录/群/联系人/断连回调契约层 |
| `example_hooks.go` | Hook 示例、AccountHook、HookManager |
| `http_client.go` | HTTP 客户端接口与默认实现 |
| `flex_int64.go` | 容错整数类型（number/string 双格式兼容） |
| `api_*.go`（112 个） | 主动调用 API：每接口一文件，命名 `api_<域>_<方法名snake_case>.go`（SendTextMsg → api_message_send_text_msg.go）；域编号与 docs/api 一致 |
| `api_common_types.go` | 跨接口共享数据结构（CommonContactResp / CommonGroupResp / Sns* 等） |
| `chat_msgtype.go` | 聊天消息 msgtype 常量 |
| 测试文件 | `*_test.go`：按被测对象拆分（shard_manager / shared_pool / dispatcher / dispatcher_adapter / hook_manager / callback_payload_msg / callback_payload_account / callback_http / flex_int64 / group_qrcode） |

---

*文档版本：v1.1.0 | 最后更新：2026-04-21*
