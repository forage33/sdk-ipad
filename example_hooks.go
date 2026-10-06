package sdk_ipad

// example_hooks.go 提供 Hook 接口的示例实现和 Hook 管理器。
//
// 包含:
//   - ExampleHook: 基础 Hook 实现示例，演示按消息类型分发处理
//   - AccountHook: 扩展 Hook 接口，支持账号级别的初始化和清理
//   - ExampleAccountHook: AccountHook 的示例实现
//   - HookManager: Hook 集中管理器，支持注册/取消注册和批量注册到分发器

import (
	"sync"
)

// ExampleHook 示例 Hook 实现。
//
// 演示如何实现 Hook 接口，按消息类型路由到不同的处理方法。
// 可作为自定义 Hook 的基础模板。
//
// 字段:
//   - name: Hook 名称，用于注册时的唯一标识
//   - priority: 优先级，数值越大优先级越高
type ExampleHook struct {
	name     string
	priority int
}

// NewExampleHook 创建示例 Hook。
//
// 参数:
//   - name: Hook 名称
//   - priority: 优先级
//
// 返回: *ExampleHook 示例 Hook 实例
func NewExampleHook(name string, priority int) *ExampleHook {
	return &ExampleHook{
		name:     name,
		priority: priority,
	}
}

// Handle 处理回调消息。
//
// 根据消息类型路由到对应的处理方法。未匹配的消息类型静默忽略。
//
// 参数:
//   - ctx: Hook 执行上下文
//   - msg: 回调消息
//
// 返回: error 处理失败时返回错误
func (h *ExampleHook) Handle(ctx *HookContext, msg *CallBackMsg) error {
	// 这里可以根据消息类型进行不同的处理
	switch msg.Type {
	case ScanReturn:
		return h.handleScanReturn(ctx, msg)
	case LoginSuccessReturn:
		return h.handleLoginSuccess(ctx, msg)
	case MsgCallback: // 使用 MsgCallback 而不是 NewMessage
		return h.handleNewMessage(ctx, msg)
	case ApplyFriend:
		return h.handleApplyFriend(ctx, msg)
	default:
		// 对于其他消息类型，记录但不处理
		return nil
	}
}

// Priority 返回 Hook 优先级。
//
// 返回: int 优先级数值，越大优先级越高
func (h *ExampleHook) Priority() int {
	return h.priority
}

// Name 返回 Hook 名称。
//
// 返回: string Hook 唯一标识名称
func (h *ExampleHook) Name() string {
	return h.name
}

// handleScanReturn 处理扫码返回消息。
func (h *ExampleHook) handleScanReturn(ctx *HookContext, msg *CallBackMsg) error {
	return nil
}

// handleLoginSuccess 处理登录成功消息。
func (h *ExampleHook) handleLoginSuccess(ctx *HookContext, msg *CallBackMsg) error {
	return nil
}

// handleNewMessage 处理新聊天消息。
func (h *ExampleHook) handleNewMessage(ctx *HookContext, msg *CallBackMsg) error {
	return nil
}

// handleApplyFriend 处理好友申请消息。
func (h *ExampleHook) handleApplyFriend(ctx *HookContext, msg *CallBackMsg) error {
	return nil
}

// AccountHook 账号级别的 Hook 接口。
//
// 在基础 Hook 接口上扩展了账号生命周期管理能力：
//   - InitAccount: 当新账号首次接入时调用，用于初始化账号级别的资源
//   - CleanupAccount: 当账号断开或过期时调用，用于释放资源
type AccountHook interface {
	Hook
	// InitAccount 初始化账号。
	//
	// 参数:
	//   - uuid: 账号唯一标识
	//   - sdk: SDK 实例引用
	//
	// 返回: error 初始化失败时返回错误
	InitAccount(uuid string, sdk *SdkIPad) error
	// CleanupAccount 清理账号资源。
	//
	// 参数:
	//   - uuid: 账号唯一标识
	//
	// 返回: error 清理失败时返回错误
	CleanupAccount(uuid string) error
}

// ExampleAccountHook 示例账号 Hook 实现。
//
// 内嵌 ExampleHook，扩展了账号追踪能力。
// 使用 sync.RWMutex 保护 accounts map 的并发访问。
//
// 字段:
//   - ExampleHook: 内嵌的基础 Hook
//   - accounts: 已初始化账号的追踪映射
type ExampleAccountHook struct {
	*ExampleHook
	mu       sync.RWMutex    // Fix #11: 保护 accounts map 的并发访问
	accounts map[string]bool // 跟踪已初始化的账号
}

// NewExampleAccountHook 创建示例账号 Hook。
//
// 参数:
//   - name: Hook 名称
//   - priority: 优先级
//
// 返回: *ExampleAccountHook 示例账号 Hook 实例
func NewExampleAccountHook(name string, priority int) *ExampleAccountHook {
	return &ExampleAccountHook{
		ExampleHook: NewExampleHook(name, priority),
		accounts:    make(map[string]bool),
	}
}

// InitAccount 初始化账号，将其加入追踪映射。
//
// 参数:
//   - uuid: 账号唯一标识
//   - sdk: SDK 实例引用
//
// 返回: error 始终为 nil
func (h *ExampleAccountHook) InitAccount(uuid string, sdk *SdkIPad) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.accounts[uuid] = true
	return nil
}

// CleanupAccount 清理账号，从追踪映射中移除。
//
// 参数:
//   - uuid: 账号唯一标识
//
// 返回: error 始终为 nil
func (h *ExampleAccountHook) CleanupAccount(uuid string) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.accounts, uuid)
	return nil
}

// Handle 重写 Handle 方法，添加账号上下文支持。
//
// 参数:
//   - ctx: Hook 执行上下文
//   - msg: 回调消息
//
// 返回: error 处理失败时返回错误
func (h *ExampleAccountHook) Handle(ctx *HookContext, msg *CallBackMsg) error {
	if ctx.AccountCtx == nil {
		return h.ExampleHook.Handle(ctx, msg)
	}
	return h.ExampleHook.Handle(ctx, msg)
}

// HookManager Hook 集中管理器。
//
// 按名称索引管理所有 Hook 实例，支持注册/取消注册、
// 按名称查找、获取全部 Hook 以及批量注册到 CallbackDispatcher。
// 使用 sync.RWMutex 保证并发安全。
type HookManager struct {
	mu    sync.RWMutex    // Fix #12: 保护 hooks map 的并发访问
	hooks map[string]Hook // 按名称索引的Hook
}

// NewHookManager 创建 Hook 管理器。
//
// 返回: *HookManager 初始化完成的 Hook 管理器
func NewHookManager() *HookManager {
	return &HookManager{
		hooks: make(map[string]Hook),
	}
}

// RegisterHook 注册 Hook 到管理器。
//
// 同名 Hook 会被覆盖。
//
// 参数:
//   - hook: 实现 Hook 接口的实例
//
// 返回: error 始终为 nil
func (hm *HookManager) RegisterHook(hook Hook) error {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	hm.hooks[hook.Name()] = hook
	return nil
}

// UnregisterHook 从管理器中取消注册 Hook。
//
// 参数:
//   - name: Hook 名称
//
// 返回: error 始终为 nil
func (hm *HookManager) UnregisterHook(name string) error {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	delete(hm.hooks, name)
	return nil
}

// GetHook 按名称获取 Hook。
//
// 参数:
//   - name: Hook 名称
//
// 返回: Hook 对应的 Hook 实例，不存在时返回 nil
func (hm *HookManager) GetHook(name string) Hook {
	hm.mu.RLock()
	defer hm.mu.RUnlock()
	return hm.hooks[name]
}

// GetAllHooks 获取管理器中的所有 Hook。
//
// 返回: []Hook Hook 列表副本
func (hm *HookManager) GetAllHooks() []Hook {
	hm.mu.RLock()
	defer hm.mu.RUnlock()
	hooks := make([]Hook, 0, len(hm.hooks))
	for _, hook := range hm.hooks {
		hooks = append(hooks, hook)
	}
	return hooks
}

// RegisterAllToDispatcher 将管理器中所有 Hook 批量注册到分发器。
//
// 每个 Hook 会被注册到所有指定的消息类型上。
//
// 参数:
//   - dispatcher: 目标回调消息分发器
//   - msgTypes: 要注册的消息类型列表
//
// 返回: error 注册失败时返回第一个错误
func (hm *HookManager) RegisterAllToDispatcher(dispatcher *CallbackDispatcher, msgTypes []int) error {
	hm.mu.RLock()
	defer hm.mu.RUnlock()
	for _, hook := range hm.hooks {
		for _, msgType := range msgTypes {
			if err := dispatcher.RegisterHook(msgType, hook); err != nil {
				return err
			}
		}
	}
	return nil
}
