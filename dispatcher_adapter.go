package sdk_ipad

// dispatcher_adapter.go 实现分发器适配器。
//
// DispatcherAdapter 是一个便捷封装层，将 CallbackDispatcher、
// SharedWorkerPool 和 SdkIPad 组合在一起，提供统一的回调处理接口。
// 适用于需要快速集成回调处理能力的场景。
//
// 提供全局单例 DefaultAdapter，可通过 GetDefaultAdapter 获取，
// HandleCallbackGlobal 函数可直接用于 /callback 接口路由。

import (
	"sync"
)

// DispatcherAdapter 分发器适配器。
//
// 组合 CallbackDispatcher、SharedWorkerPool 和 SdkIPad，
// 对外提供统一的回调处理、Hook 管理和统计查询接口。
//
// 字段:
//   - dispatcher: 回调消息分发器
//   - workerPool: 共享 Worker 池
//   - sdk: SDK 实例引用
type DispatcherAdapter struct {
	dispatcher *CallbackDispatcher
	workerPool *SharedWorkerPool
	sdk        *SdkIPad
}

// NewDispatcherAdapter 创建分发器适配器。
//
// 内部自动创建 SharedWorkerPool 和 CallbackDispatcher 实例。
// 如果 sdk 为 nil，则自动创建一个默认 SdkIPad 实例。
//
// 参数:
//   - sdk: SDK 实例（可为 nil，将自动创建默认实例）
//   - config: 分发器配置
//
// 返回: *DispatcherAdapter 初始化完成的适配器实例
func NewDispatcherAdapter(sdk *SdkIPad, config DispatcherConfig) *DispatcherAdapter {
	if sdk == nil {
		sdk = NewSdkIPad()
	}

	// 创建Worker池
	workerPool := NewSharedWorkerPool(sdk.config.WorkerPool)

	// 创建分发器
	dispatcher := NewCallbackDispatcher(sdk, sdk.GetShardManager(), workerPool, config)

	return &DispatcherAdapter{
		dispatcher: dispatcher,
		workerPool: workerPool,
		sdk:        sdk,
	}
}

// Start 启动适配器。
//
// 启动内部的 CallbackDispatcher（Worker 池在创建时已自动启动）。
//
// 返回: error 启动失败时返回错误
func (da *DispatcherAdapter) Start() error {
	// 启动Worker池已经在创建时自动启动
	// 启动分发器
	return da.dispatcher.Start()
}

// Stop 停止适配器。
//
// 依次停止 CallbackDispatcher 和 SharedWorkerPool。
func (da *DispatcherAdapter) Stop() {
	da.dispatcher.Stop()
	da.workerPool.Stop()
}

// HandleCallback 处理回调消息（主要接口）。
//
// 将回调消息委托给内部 CallbackDispatcher 处理。
//
// 参数:
//   - msg: 回调消息
//
// 返回: error 处理失败时返回错误
func (da *DispatcherAdapter) HandleCallback(msg *CallBackMsg) error {
	return da.dispatcher.HandleCallback(msg)
}

// RegisterHook 注册 Hook 到指定消息类型。
//
// 参数:
//   - msgType: 消息类型（对应 msg_ipad.go 中的常量）
//   - hook: 实现 Hook 接口的实例
//
// 返回: error 注册失败时返回错误
func (da *DispatcherAdapter) RegisterHook(msgType int, hook Hook) error {
	return da.dispatcher.RegisterHook(msgType, hook)
}

// UnregisterHook 取消注册指定消息类型下的 Hook。
//
// 参数:
//   - msgType: 消息类型
//   - hookName: Hook 名称
//
// 返回: error 取消注册失败时返回错误
func (da *DispatcherAdapter) UnregisterHook(msgType int, hookName string) error {
	return da.dispatcher.UnregisterHook(msgType, hookName)
}

// GetStats 获取分发器统计信息。
//
// 返回: DispatcherStats 分发器运行时统计快照
func (da *DispatcherAdapter) GetStats() DispatcherStats {
	return da.dispatcher.GetStats()
}

// GetShardStats 获取指定分片的统计信息。
//
// 参数:
//   - shardID: 分片 ID
//
// 返回: ShardStats 分片统计信息
func (da *DispatcherAdapter) GetShardStats(shardID int) ShardStats {
	return da.sdk.GetShardStats(shardID)
}

// GetAllShardStats 获取所有分片的统计信息。
//
// 返回: []ShardStats 所有分片统计信息列表
func (da *DispatcherAdapter) GetAllShardStats() []ShardStats {
	return da.sdk.GetAllShardStats()
}

// GetWorkerPoolStats 获取 Worker 池统计信息。
//
// 返回: PoolStats Worker 池运行时统计快照
func (da *DispatcherAdapter) GetWorkerPoolStats() PoolStats {
	return da.workerPool.GetStats()
}

// 默认适配器单例。
var (
	defaultAdapter     *DispatcherAdapter
	defaultAdapterOnce sync.Once
)

// GetDefaultAdapter 获取默认全局适配器（单例模式）。
//
// 首次调用时自动创建默认 SdkIPad 实例和适配器，并启动服务。
// 后续调用返回同一实例。线程安全。
//
// 返回: *DispatcherAdapter 全局默认适配器
func GetDefaultAdapter() *DispatcherAdapter {
	defaultAdapterOnce.Do(func() {
		sdk := NewSdkIPad()
		defaultAdapter = NewDispatcherAdapter(sdk, DefaultDispatcherConfig())
		defaultAdapter.Start()
	})
	return defaultAdapter
}

// HandleCallbackGlobal 全局回调处理函数。
//
// 通过默认全局适配器处理回调消息，可直接用于 /callback 接口路由。
//
// 参数:
//   - msg: 回调消息
//
// 返回: error 处理失败时返回错误
func HandleCallbackGlobal(msg *CallBackMsg) error {
	adapter := GetDefaultAdapter()
	return adapter.HandleCallback(msg)
}
