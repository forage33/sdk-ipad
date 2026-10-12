package sdk_ipad

// callback_dispatcher.go 实现回调消息分发器。
//
// CallbackDispatcher 负责接收协议服务推送的回调消息，根据消息类型
// 路由到已注册的 Hook 链中进行处理。支持优先级排序、并发控制、
// 原子统计和死信队列等特性。
//
// 架构关系: ShardManager → SharedWorkerPool → CallbackDispatcher → HookRegistry

import (
	"sync"
	"sync/atomic"
	"time"
)

// CallbackDispatcher 回调消息分发器。
//
// 核心组件，负责将回调消息分发到对应的 Hook 处理链。
// 内部依赖 ShardManager 进行账号路由，依赖 SharedWorkerPool 进行并发调度。
//
// 字段:
//   - shardManager: 分片管理器，用于获取/创建账号上下文
//   - workerPool: 共享 Worker 池，用于并发执行任务
//   - hookRegistry: Hook 注册表，按消息类型索引
//   - config: 分发器配置
//   - stats: 运行时统计（原子操作）
//   - sdk: SDK 实例引用，传递给 HookContext
type CallbackDispatcher struct {
	shardManager *ShardManager
	workerPool   *SharedWorkerPool
	hookRegistry *HookRegistry
	config       DispatcherConfig
	stats        DispatcherStats
	mu           sync.Mutex
	isRunning    atomic.Bool
	stopOnce     sync.Once
	stopChan     chan struct{}
	sdk          *SdkIPad // SDK引用，用于传递给HookContext
}

// DispatcherConfig 分发器配置。
//
// 注意：以下配置项目前为预留设计，分发链路（Submit → Worker 池）暂未消费，
// 实际并发与容量由 SharedWorkerPool 配置（Min/MaxWorkers、QueueCapacity）承担；
// 回调可靠性由 CallbackHTTPHandler 传输层重试（3 次指数退避）兜底。
//
// 字段:
//   - MaxConcurrent: 最大并发处理数，默认 100（预留，未消费）
//   - BatchSize: 批处理大小，默认 100（预留，未消费）
//   - Timeout: 单条消息处理超时，默认 5s（预留：仅透传至 Task.Timeout，执行层未启用超时）
//   - RetryCount: 失败重试次数，默认 3（预留，未消费）
//   - RetryInterval: 重试间隔，默认 100ms（预留，未消费）
//   - DeadLetterQueue: 是否启用死信队列，默认 true（预留，未消费）
type DispatcherConfig struct {
	MaxConcurrent   int           // 最大并发处理数，默认100（预留，未消费）
	BatchSize       int           // 批处理大小，默认100（预留，未消费）
	Timeout         time.Duration // 处理超时，默认5s（预留：透传 Task.Timeout，未启用）
	RetryCount      int           // 重试次数，默认3（预留，未消费）
	RetryInterval   time.Duration // 重试间隔，默认100ms（预留，未消费）
	DeadLetterQueue bool          // 是否启用死信队列，默认true（预留，未消费）
}

// DefaultDispatcherConfig 返回默认分发器配置。
//
// 返回: DispatcherConfig 包含推荐默认值的配置
func DefaultDispatcherConfig() DispatcherConfig {
	return DispatcherConfig{
		MaxConcurrent:   100,
		BatchSize:       100,
		Timeout:         5 * time.Second,
		RetryCount:      3,
		RetryInterval:   100 * time.Millisecond,
		DeadLetterQueue: true,
	}
}

// DispatcherStats 分发器运行时统计。
//
// 所有数值字段均通过原子操作更新，支持并发安全读取。
//
// 字段:
//   - TotalReceived: 累计接收消息数
//   - TotalProcessed: 累计成功处理数
//   - TotalFailed: 累计失败数
//   - TotalRetried: 累计重试数（预留：当前无重试机制写入，恒为 0）
//   - AvgProcessTime: 平均处理耗时（纳秒，移动平均）
//   - CurrentQueueSize: 当前队列中的消息数
//   - MaxQueueSize: 历史队列峰值
//   - LastReceivedTime: 最后一次收到消息的时间戳（UnixNano）
type DispatcherStats struct {
	TotalReceived    int64
	TotalProcessed   int64
	TotalFailed      int64
	TotalRetried     int64 // 预留：无重试机制写入，恒为 0
	AvgProcessTime   int64 // 纳秒，使用 atomic 操作
	CurrentQueueSize int32
	MaxQueueSize     int32
	LastReceivedTime int64 // UnixNano，使用 atomic 操作
}

// HookRegistry Hook 注册表。
//
// 维护消息类型到 Hook 列表的映射关系，支持并发读写。
type HookRegistry struct {
	hooks map[int][]Hook // 消息类型 -> Hook列表
	mu    sync.RWMutex
}

// Hook 回调处理钩子接口。
//
// 实现此接口可注册到 CallbackDispatcher 中处理指定类型的回调消息。
// 同一消息类型可注册多个 Hook，按 Priority 从高到低依次执行。
//
// 方法:
//   - Handle: 处理回调消息，返回 error 表示处理失败（不中断后续 Hook 执行）
//   - Priority: 返回优先级，数值越大优先级越高
//   - Name: 返回 Hook 名称，用于注册/取消注册时的唯一标识
type Hook interface {
	Handle(ctx *HookContext, msg *CallBackMsg) error
	Priority() int // 优先级，数值越大优先级越高
	Name() string
}

// HookContext Hook 执行上下文。
//
// 在 Hook.Handle 调用时传入，提供账号上下文和 SDK 引用。
//
// 字段:
//   - AccountCtx: 当前账号上下文（由 ShardManager 提供）
//   - SDK: SDK 实例引用，可用于在 Hook 中调用 API
//   - StartTime: 本次处理的开始时间
type HookContext struct {
	AccountCtx *AccountContext
	SDK        *SdkIPad
	StartTime  time.Time
}

// NewCallbackDispatcher 创建回调消息分发器。
//
// 参数:
//   - sdk: SDK 实例引用
//   - shardManager: 分片管理器
//   - workerPool: 共享 Worker 池
//   - config: 分发器配置，零值字段将使用默认值
//
// 返回: *CallbackDispatcher 初始化完成的分发器实例
func NewCallbackDispatcher(sdk *SdkIPad, shardManager *ShardManager, workerPool *SharedWorkerPool, config DispatcherConfig) *CallbackDispatcher {
	if config.MaxConcurrent <= 0 {
		config.MaxConcurrent = 100
	}
	if config.BatchSize <= 0 {
		config.BatchSize = 100
	}
	if config.Timeout <= 0 {
		config.Timeout = 5 * time.Second
	}

	return &CallbackDispatcher{
		shardManager: shardManager,
		workerPool:   workerPool,
		hookRegistry: &HookRegistry{
			hooks: make(map[int][]Hook),
		},
		config:   config,
		stats:    DispatcherStats{},
		stopChan: make(chan struct{}),
		sdk:      sdk,
	}
}

// Start 启动分发器。
//
// 将分发器标记为运行状态，使其可以接收并处理回调消息。
// 重复调用是安全的（幂等）。
//
// 返回: error 始终为 nil
func (d *CallbackDispatcher) Start() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.isRunning.Load() {
		return nil
	}

	d.isRunning.Store(true)
	return nil
}

// Stop 停止分发器。
//
// 将分发器标记为停止状态，后续提交的消息将被拒绝。
// 使用 sync.Once 保证只执行一次。
func (d *CallbackDispatcher) Stop() {
	d.stopOnce.Do(func() {
		d.isRunning.Store(false)
		close(d.stopChan)
	})
}

// RegisterHook 注册 Hook 到指定消息类型。
//
// 同一消息类型的 Hook 按 Priority 从高到低排序。
//
// 参数:
//   - msgType: 消息类型（对应 callback_base.go 中的常量）
//   - hook: 实现 Hook 接口的实例
//
// 返回: error 始终为 nil
func (d *CallbackDispatcher) RegisterHook(msgType int, hook Hook) error {
	d.hookRegistry.mu.Lock()
	defer d.hookRegistry.mu.Unlock()

	hooks := d.hookRegistry.hooks[msgType]

	// 按优先级插入（从高到低）
	insertPos := 0
	for i, h := range hooks {
		if hook.Priority() > h.Priority() {
			insertPos = i
			break
		}
		insertPos = i + 1
	}

	// 插入到正确位置
	if insertPos == 0 {
		d.hookRegistry.hooks[msgType] = append([]Hook{hook}, hooks...)
	} else if insertPos >= len(hooks) {
		d.hookRegistry.hooks[msgType] = append(hooks, hook)
	} else {
		d.hookRegistry.hooks[msgType] = append(hooks[:insertPos], append([]Hook{hook}, hooks[insertPos:]...)...)
	}

	return nil
}

// UnregisterHook 取消注册指定消息类型下的 Hook。
//
// 参数:
//   - msgType: 消息类型
//   - hookName: 要取消注册的 Hook 名称
//
// 返回: error 始终为 nil（未找到时静默忽略）
func (d *CallbackDispatcher) UnregisterHook(msgType int, hookName string) error {
	d.hookRegistry.mu.Lock()
	defer d.hookRegistry.mu.Unlock()

	hooks, exists := d.hookRegistry.hooks[msgType]
	if !exists {
		return nil
	}

	for i, hook := range hooks {
		if hook.Name() == hookName {
			d.hookRegistry.hooks[msgType] = append(hooks[:i], hooks[i+1:]...)
			if len(d.hookRegistry.hooks[msgType]) == 0 {
				delete(d.hookRegistry.hooks, msgType)
			}
			return nil
		}
	}

	return nil
}

// HandleCallback 处理回调消息（主要入口点）。
//
// 接收回调消息，更新统计信息，根据消息类型确定优先级，
// 创建任务并提交到 SharedWorkerPool 异步执行。
//
// 参数:
//   - msg: 回调消息
//
// 返回: error 分发器未运行或提交失败时返回错误
func (d *CallbackDispatcher) HandleCallback(msg *CallBackMsg) error {
	if msg == nil {
		return ErrInvalidMessageType
	}

	if !d.isRunning.Load() {
		return ErrDispatcherNotRunning
	}

	// Fix #4: 使用原子操作更新统计
	atomic.AddInt64(&d.stats.TotalReceived, 1)
	atomic.StoreInt64(&d.stats.LastReceivedTime, time.Now().UnixNano())
	queueSize := atomic.AddInt32(&d.stats.CurrentQueueSize, 1)
	defer atomic.AddInt32(&d.stats.CurrentQueueSize, -1)

	// 原子更新最大队列大小
	for {
		oldMax := atomic.LoadInt32(&d.stats.MaxQueueSize)
		if queueSize <= oldMax {
			break
		}
		if atomic.CompareAndSwapInt32(&d.stats.MaxQueueSize, oldMax, queueSize) {
			break
		}
	}

	// 获取消息优先级
	priority := d.getMessagePriority(msg.Type)

	// 创建任务
	task := &Task{
		AccountUUID: msg.UUID,
		Msg:         msg,
		Priority:    priority,
		Handler:     &callbackTaskHandler{dispatcher: d},
		CreatedAt:   time.Now(),
		Timeout:     d.config.Timeout,
	}

	// 提交到Worker池
	startTime := time.Now()
	err := d.workerPool.Submit(task)
	processTime := time.Since(startTime)

	if err != nil {
		atomic.AddInt64(&d.stats.TotalFailed, 1)
		return err
	}

	// 原子更新平均处理时间（移动平均）
	for {
		oldAvg := atomic.LoadInt64(&d.stats.AvgProcessTime)
		var newAvg int64
		if oldAvg == 0 {
			newAvg = int64(processTime)
		} else {
			newAvg = int64(0.8*float64(oldAvg) + 0.2*float64(processTime))
		}
		if atomic.CompareAndSwapInt64(&d.stats.AvgProcessTime, oldAvg, newAvg) {
			break
		}
	}

	return nil
}

// getMessagePriority 根据消息类型返回处理优先级。
//
// 注意：优先级为预留能力——当前仅写入 Task.Priority 供调试参考，
// 任务队列为单一 FIFO，未参与调度排序（见 PriorityQueueConfig）。
//
// 优先级分为三档:
//   - 2（高）: 登录回调、异常断开等需要立即响应的消息
//   - 1（中）: 好友申请、群创建/解散等重要业务消息
//   - 0（低）: 其他普通消息
//
// 参数:
//   - msgType: 消息类型常量
//
// 返回: int 优先级（0/1/2）
func (d *CallbackDispatcher) getMessagePriority(msgType int) int {
	// 高优先级：登录回调、异常断开
	switch msgType {
	case ScanReturn, // 扫码返回
		ConfirmReturn,      // 确认扫码
		LoginSuccessReturn, // 登录成功
		Exception,          // 异常断开
		LoginInOtherDevice: // 其他设备登录
		return 2 // 高优先级

	// 中优先级：重要业务消息
	case ApplyFriend, // 好友申请
		NewGroup,          // 新建群
		RemoveGroupMember, // 移除群成员
		GroupDisBand:      // 群解散
		return 1 // 中优先级

	// 低优先级：普通消息
	default:
		return 0 // 低优先级
	}
}

// processCallback 处理单个回调消息。
//
// 从 HookRegistry 中获取对应消息类型的 Hook 列表，
// 创建 HookContext 并按优先级依次执行。单个 Hook 失败不影响后续执行。
//
// 参数:
//   - accountCtx: 账号上下文
//   - msg: 回调消息
//
// 返回: error 最后一个失败的 Hook 错误（如有），无 Hook 注册时走默认处理
func (d *CallbackDispatcher) processCallback(accountCtx *AccountContext, msg *CallBackMsg) error {
	// Fix #9: 在 RLock 保护下复制 hooks 切片，避免迭代期间的竞态
	d.hookRegistry.mu.RLock()
	original, exists := d.hookRegistry.hooks[msg.Type]
	if !exists {
		d.hookRegistry.mu.RUnlock()
		return d.defaultHandleCallback(accountCtx, msg)
	}
	// 复制切片引用（slice header），注册时创建新切片所以这里安全
	hooks := make([]Hook, len(original))
	copy(hooks, original)
	d.hookRegistry.mu.RUnlock()

	// 创建Hook上下文
	hookCtx := &HookContext{
		AccountCtx: accountCtx,
		SDK:        d.sdk,
		StartTime:  time.Now(),
	}

	// 按顺序执行Hook
	var lastErr error
	for _, hook := range hooks {
		if err := hook.Handle(hookCtx, msg); err != nil {
			lastErr = err
			// 记录错误，但继续执行下一个Hook
		}
	}

	return lastErr
}

// defaultHandleCallback 默认回调处理（无 Hook 注册时的兜底逻辑）。
//
// 参数:
//   - accountCtx: 账号上下文
//   - msg: 回调消息
//
// 返回: error 始终为 nil
func (d *CallbackDispatcher) defaultHandleCallback(accountCtx *AccountContext, msg *CallBackMsg) error {
	return nil
}

// GetStats 获取分发器运行时统计（线程安全快照）。
//
// 所有字段通过原子操作读取，返回一致性快照。
//
// 返回: DispatcherStats 统计信息副本
func (d *CallbackDispatcher) GetStats() DispatcherStats {
	return DispatcherStats{
		TotalReceived:    atomic.LoadInt64(&d.stats.TotalReceived),
		TotalProcessed:   atomic.LoadInt64(&d.stats.TotalProcessed),
		TotalFailed:      atomic.LoadInt64(&d.stats.TotalFailed),
		TotalRetried:     atomic.LoadInt64(&d.stats.TotalRetried),
		AvgProcessTime:   atomic.LoadInt64(&d.stats.AvgProcessTime),
		CurrentQueueSize: atomic.LoadInt32(&d.stats.CurrentQueueSize),
		MaxQueueSize:     atomic.LoadInt32(&d.stats.MaxQueueSize),
		LastReceivedTime: atomic.LoadInt64(&d.stats.LastReceivedTime),
	}
}

// callbackTaskHandler 回调任务处理器。
//
// 实现 TaskHandler 接口，作为 SharedWorkerPool 中执行回调处理的载体。
type callbackTaskHandler struct {
	dispatcher *CallbackDispatcher
}

// Handle 执行回调任务。
//
// 从任务中提取 CallBackMsg，通过 ShardManager 获取账号上下文，
// 然后委托给 CallbackDispatcher.processCallback 处理。
//
// 参数:
//   - task: Worker 池分配的任务
//
// 返回: error 消息类型不匹配、账号未找到或处理失败时返回错误
func (h *callbackTaskHandler) Handle(task *Task) error {
	msg, ok := task.Msg.(*CallBackMsg)
	if !ok {
		return ErrInvalidMessageType
	}

	// 获取账号上下文
	accountCtx := h.dispatcher.shardManager.GetOrCreateAccount(msg.UUID)
	if accountCtx == nil {
		return ErrAccountNotFound
	}

	// 处理回调
	err := h.dispatcher.processCallback(accountCtx, msg)
	if err != nil {
		atomic.AddInt64(&h.dispatcher.stats.TotalFailed, 1)
	} else {
		atomic.AddInt64(&h.dispatcher.stats.TotalProcessed, 1)
	}

	return err
}

// 分发器相关错误定义。
var (
	ErrDispatcherNotRunning = &Error{Code: -2001, Message: "dispatcher not running"} // 分发器未运行
	ErrInvalidMessageType   = &Error{Code: -2002, Message: "invalid message type"}   // 消息类型无效
	ErrAccountNotFound      = &Error{Code: -2003, Message: "account not found"}      // 账号未找到
	ErrHookExecutionFailed  = &Error{Code: -2004, Message: "hook execution failed"}  // Hook 执行失败
)
