# sdk-ipad — 企业微信 iPad 协议 SDK

独立 Go 模块（`module github.com/forage33/sdk-ipad`），封装与企微 iPad 协议服务的全部交互：**主动调用 API**、**回调分发机制**、**回调数据契约**、**HTTP 传输适配**。

仅依赖标准库 + `go.uber.org/zap` + `github.com/tidwall/gjson`，不依赖任何宿主项目代码，可独立编译、测试、打 tag 发版。仓库地址：<https://github.com/forage33/sdk-ipad>。

架构设计（分片管理、共享 Worker 池、Hook 机制、消息处理流程、并发安全与生命周期）详见 [ARCHITECTURE.md](ARCHITECTURE.md)。

## 快速开始

### 1. 创建 SDK 实例

```go
cfg := &sdk_ipad.SdkConfig{
    ShardConfig:   sdk_ipad.DefaultShardConfig(),   // 分片（万号承载建议调大，见 DefaultDispatcherConfig 注释）
    WorkerPool:    sdk_ipad.DefaultWorkerPoolConfig(),
    PriorityQueue: sdk_ipad.DefaultPriorityQueueConfig(),
    Logger:        myLogger, // 可选，*zap.Logger；nil 时回退全局 zap.L()
}
sdk := sdk_ipad.NewSdkIPadWithConfig(cfg)
```

### 2. 主动调用协议接口

所有 API 方法第一个参数均为协议服务地址 `host`（如 `http://8.148.223.15:8083`），第二个参数为实例 `uuid`：

```go
// 初始化实例 → 设置回调地址
resp, err := sdk.InItWXWork(host, vid, "", "", 0, "", "", "ipad", 0)
_, err = sdk.SetCallBackURL(host, uuid, sdk_ipad.CallbackConfig{CallbackType: sdk_ipad.CallbackTypeHTTP, URL: callbackURL})

// 发送文本消息
_, err = sdk.SendTextMsg(host, uuid, sendUserID, "hello", false /* isRoom */)
```

### 3. 接收回调

协议服务将回调 POST 到宿主暴露的 HTTP 端点。SDK 提供与框架解耦的标准 handler（信封解析 → 分发提交 → 失败异步重试，始终应答 200）：

```go
// gin 场景挂载
engine.POST("/callback", gin.WrapF(
    sdk_ipad.CallbackHTTPHandler(adapter.HandleCallback, sdk_ipad.DefaultCallbackHandlerOptions()),
))
```

业务侧通过 Hook 声明关心的回调类型（完整示例见 `example_hooks.go`；端到端测试见 `dispatcher_test.go`）：

```go
hook := myHook // 实现 sdk_ipad.Hook 接口
hook.RegisterHandler(sdk_ipad.MsgCallback, handleMsgCallback)   // 102000 聊天消息
hook.RegisterHandler(sdk_ipad.LoginSuccessReturn, handleLogin)  // 104001 登录成功
adapter.RegisterHook(hook)
```

处理函数从**契约层**取强类型数据，不接触原始 JSON 字段名：

```go
func handleMsgCallback(ctx *sdk_ipad.HookContext, msg *sdk_ipad.CallBackMsg) error {
    env := msg.Envelope()              // 通用信封：msgtype/sender/receiver/is_room/...
    switch env.MsgType {
    case sdk_ipad.ChatMsgText:
        text := msg.AsTextMsg()        // 强类型 payload
    case sdk_ipad.ChatMsgVoice:
        voice := msg.AsVoiceMsg()      // voice_id/aes_key/voice_time 协议字段已收敛
    }
    return nil
}
```

## 文件结构与各文件作用

### 入口与配置

| 文件 | 作用 |
|---|---|
| `sdk.go` | 包文档与 SDK 主入口：`SdkIPad` 主结构体、`SdkConfig` 全局配置（分片/Worker 池/优先级队列（预留）/Logger 注入）、构造函数、账号上下文访问、回调入口 `HandleCallback`，附协议服务已知错误码参考表 |
| `sdk_interface.go` | `SdkInterface` 接口定义：服务层使用的 SDK 方法集合（`SdkIPad` 隐式实现），供调用方 mock/替换 |

### 主动调用 API（每接口一文件）

命名规则：`api_<域>_<方法名snake_case>.go`——如 `SendTextMsg` → `api_message_send_text_msg.go`、`RevokeMsg` → `api_message_revoke_msg.go`；域编号与 `docs/api` 一致；跨接口共享类型集中在 `api_common_types.go`。

| 域 | 文件前缀 | 接口 |
|---|---|---|
| 01 初始化 | `api_wx_init_*.go` | 实例初始化（含代理配置）、回调地址设置（HTTP/RABBITMQ）、运行实例查询、连接关闭 等 7 个 |
| 02 登录 | `api_login_*.go` | 二维码获取、验证码校验、自动登录、退出登录、二次验证 5 个 |
| 03 账号操作 | `api_account_*.go` | 头像设置、资料修改、个人二维码、封禁详情、对外名片、企业信息 8 个 |
| 04 开放平台 | `api_open_platform_*.go` | 企业用户/开放平台 id 互转、群 id/chatid 互转 5 个 |
| 05 标签 | `api_label_*.go` | 标签列表、增删改、标签与联系人关联 6 个 |
| 06 CDN | `api_cdn_*.go` | 网络图片/视频/文件上传、大文件上传下载、微信文件下载、CDN 初始化与鉴权 12 个 |
| 07 联系人操作 | `api_contact_*.go` | 手机号搜索、添加好友链路、内外部联系人列表、删除、备注、置顶/免打扰/拉黑 23 个 |
| 08 朋友圈 | `api_moments_*.go` | 列表/详情、发布、删除、评论、点赞 7 个 |
| 09 控制台群发 | `api_broadcast_*.go` | 群发消息确认 1 个 |
| 10 聊天标签 | `api_chat_tag_*.go` | 会话标签增删改 1 个 |
| 13 发送消息 | `api_message_*.go` | 各类消息发送、撤回、已读标记、语音转文字、消息同步、引用消息 21 个 |
| 14 群操作 | `api_group_*.go` | 群列表、成员管理、建群、群公告、群名、邀请、二维码、欢迎语、退群、解散 25 个 |

### 回调机制（接收、分发、并发）

| 文件 | 作用 |
|---|---|
| `callback_base.go` | 回调基础层：`CallBackMsg` 信封（type/uuid/json）、回调类型码常量（100xxx 登录 / 102xxx 消息 / 115xxx 群 / 116xxx 联系人 / 130xxx 标签）、`SendMsg` 内部发送结构 |
| `chat_msgtype.go` | 聊天消息 `msgtype` 常量（ChatMsgText/Voice/Video/Revoke...，消息分发的分支依据） |
| `callback_dispatcher.go` | `CallbackDispatcher` 回调分发器：Hook 注册（按消息类型，同类型按 Hook 优先级排序执行）、统计（DispatcherStats）、消息消费处理（消息优先级为预留，未参与调度） |
| `dispatcher_adapter.go` | `DispatcherAdapter` 便捷封装：组装 Dispatcher + ShardManager + WorkerPool，业务侧只需 `HandleCallback(msg)` 一个入口 |
| `shard_manager.go` | `ShardManager` 分片管理：一致性哈希将账号分片到固定分片，账号上下文（AccountContext）生命周期管理、过期清理、回调按账号路由 |
| `shared_pool.go` | `SharedWorkerPool` 动态扩缩容 goroutine 工作池：Min/MaxWorkers、队列容量、优先级队列（PriorityQueueConfig，预留未实现） |
| `callback_http.go` | `CallbackHTTPHandler` 传输层适配：标准 `net/http` handler（JSON 信封解析 → 分发 → 失败异步重试 3 次指数退避 → 始终应答 200），与 Web 框架解耦（gin 用 `gin.WrapF` 挂载） |
| `example_hooks.go` | Hook 接口的示例实现与 Hook 管理器，业务侧 Hook 编写模板 |

### 回调契约层（协议字段知识的唯一归口）

| 文件 | 作用 |
|---|---|
| `callback_payload_msg.go` | 消息回调（102000）契约：`MessageEnvelope` 通用信封（8 字段）、`RevokePayload` 撤回定位、12 类消息 payload（系统/文本/GIF/小程序/链接/视频号/位置/红包/名片/媒体文件/视频/语音）与 `As*` 宽容访问器（gjson 语义，缺失返零值）；视频 `aes_key`/`aeskey` 双名兼容在此收敛 |
| `callback_payload_account.go` | 其余域契约：登录域 5 类 payload（`AsScanReturn`/`AsLoginSuccess` 等，**严格** `json.Unmarshal` 语义返回 error）；群域（`RoomID`/`AsNewRoom`/`GroupMemberList`/`AsGroupNameChange` 三级回退）；联系人域（`AsFriendChangeList`/`AsInnerLabelChangeList`/`AsContactEventMsg`）；断连域 `DisconnectVid`（string/number 双型兼容） |

> 解析语义约定：登录域历史上走 `json.Unmarshal` 判错丢弃，故保持严格语义；群/联系人/断连/消息域历史上走 gjson 逐字段取值，故保持宽容零值语义。**新增访问器必须与既有语义域一致。**

### 基础设施

| 文件 | 作用 |
|---|---|
| `http_client.go` | HTTP 客户端封装：JSON POST、网络错误自动重试、多协议服务器高并发调优的 `optimizedTransport`（连接池参数注释含容量推算依据） |
| `flex_int64.go` | `FlexInt64` 容错类型：协议整数字段可能返回 number 或 string，统一兼容 |

### 测试与示例

| 文件 | 作用 |
|---|---|
| `dispatcher_test.go` | 分发器用例（Hook 注册/消费、优先级队列标签） |
| `dispatcher_adapter_test.go` | 适配器用例（独立实例，不触碰全局单例） |
| `hook_manager_test.go` | Hook 管理器用例 |
| `shard_manager_test.go` | 分片管理器用例 + 基准 |
| `shared_pool_test.go` | Worker 池用例 + 基准 |
| `callback_payload_msg_test.go` | 消息回调契约层用例（宽容解析/双名兼容/字符串视图） |
| `callback_payload_account_test.go` | 登录/群/联系人/断连域契约用例（严格语义/三级回退/双型兼容） |
| `callback_http_test.go` | 回调 HTTP 传输层用例（应答协议/异步重试） |
| `flex_int64_test.go` | `FlexInt64` 双格式解析用例 |
| `group_qrcode_test.go` | 群二维码接口解析用例 |

### 文档

| 路径 | 内容 |
|---|---|
| `ARCHITECTURE.md` | 架构说明：分片管理、共享 Worker 池、Hook 机制、消息处理流程、并发安全与生命周期（仓库根目录） |

> 接口文档（API 域）与回调文档不随仓库发布，需要请联系作者（QQ：1801533361）获取。

## 配置参考（SdkConfig）

| 字段 | 说明 |
|---|---|
| `ShardConfig` | 分片数/最大账号数/账号 TTL/清理间隔（ShardManager） |
| `WorkerPool` | Worker 池 Min/Max、队列容量、扩容检查间隔（SharedWorkerPool） |
| `PriorityQueue` | 优先级档位数、每档容量、高优容量（回调分发；**预留**：调度未实现，当前队列为单一 FIFO） |
| `Logger` | `*zap.Logger`，宿主注入（Field 风格）；nil 时回退 `zap.L()`。SDK 不依赖宿主日志封装 |

## 独立仓库与发版

- 本仓库独立携带 `go.mod`/`go.sum`，在仓库目录内单独执行 `go build ./...`、`go test ./...`、`golangci-lint run ./...`。
- 引用方式：宿主工程 `go get github.com/forage33/sdk-ipad@v0.1.0`（联调期可用 `replace` 指向本地副本）。
- 独立发版：打语义化 tag（如 `v0.1.0`）并推送，宿主按 tag 升级版本即可，import 路径不变。
- 协议字段演进时**只改本仓库**（优先契约层文件）；协议文档由作者内部维护，不随仓库发布。

## 联系作者

- QQ：1801533361
- 需要 API 接口文档、回调文档或接入支持，请通过 QQ 联系作者发送。
