package sdk_ipad

import (
	"sync"
	"testing"
	"time"
)

// TestShardManager 测试分片管理器
func TestShardManager(t *testing.T) {
	config := DefaultShardConfig()
	config.ShardCount = 4 // 测试用较小的分片数
	config.AccountTTL = 1 * time.Minute

	sm := NewShardManager(config)

	// 测试获取或创建账号上下文
	account1 := sm.GetOrCreateAccount("test-uuid-1")
	if account1 == nil {
		t.Fatal("GetOrCreateAccount failed")
	}
	if account1.uuid != "test-uuid-1" {
		t.Errorf("Expected uuid test-uuid-1, got %s", account1.uuid)
	}

	// 测试获取现有账号
	account2 := sm.GetOrCreateAccount("test-uuid-1")
	if account2 != account1 {
		t.Error("GetOrCreateAccount should return same instance for same uuid")
	}

	// 测试获取分片统计
	stats := sm.GetShardStats(0)
	if stats.TotalAccounts < 0 {
		t.Error("Invalid shard stats")
	}

	// 测试更新账号统计
	sm.UpdateAccountStats("test-uuid-1", true, 100*time.Millisecond)

	// 测试处理回调
	msg := &CallBackMsg{
		UUID: "test-uuid-1",
		Type: ScanReturn,
	}

	err := sm.HandleCallback(msg, func(ctx *AccountContext, msg *CallBackMsg) error {
		if ctx.uuid != "test-uuid-1" {
			t.Errorf("Expected uuid test-uuid-1, got %s", ctx.uuid)
		}
		return nil
	})

	if err != nil {
		t.Errorf("HandleCallback failed: %v", err)
	}
}

// TestSharedWorkerPool 测试共享Worker池
func TestSharedWorkerPool(t *testing.T) {
	config := DefaultWorkerPoolConfig()
	config.MinWorkers = 2
	config.MaxWorkers = 4

	pool := NewSharedWorkerPool(config)
	defer pool.Stop()

	// 测试提交任务
	task := &Task{
		AccountUUID: "test-uuid",
		Msg:         "test message",
		Priority:    1,
		Handler:     &testTaskHandler{},
		CreatedAt:   time.Now(),
		Timeout:     5 * time.Second,
	}

	err := pool.Submit(task)
	if err != nil {
		t.Errorf("Submit task failed: %v", err)
	}

	// 等待任务处理
	time.Sleep(100 * time.Millisecond)

	// 获取统计
	stats := pool.GetStats()
	if stats.TotalTasks < 1 {
		t.Error("Task should be processed")
	}
}

// TestCallbackDispatcher 测试回调分发器
func TestCallbackDispatcher(t *testing.T) {
	// 创建测试组件
	shardConfig := DefaultShardConfig()
	shardConfig.ShardCount = 4

	workerPoolConfig := DefaultWorkerPoolConfig()
	workerPoolConfig.MinWorkers = 2
	workerPoolConfig.MaxWorkers = 4

	dispatcherConfig := DefaultDispatcherConfig()

	sm := NewShardManager(shardConfig)
	pool := NewSharedWorkerPool(workerPoolConfig)

	dispatcher := NewCallbackDispatcher(nil, sm, pool, dispatcherConfig)

	// 注册示例Hook
	hook := NewExampleHook("test-hook", 1)
	err := dispatcher.RegisterHook(ScanReturn, hook)
	if err != nil {
		t.Errorf("RegisterHook failed: %v", err)
	}

	// 启动分发器
	err = dispatcher.Start()
	if err != nil {
		t.Errorf("Start dispatcher failed: %v", err)
	}
	defer dispatcher.Stop()

	// 创建测试消息
	msg := &CallBackMsg{
		UUID: "test-uuid-1",
		Type: ScanReturn,
	}

	// 处理回调
	err = dispatcher.HandleCallback(msg)
	if err != nil {
		t.Errorf("HandleCallback failed: %v", err)
	}

	// 等待处理完成
	time.Sleep(200 * time.Millisecond)

	// 获取统计
	stats := dispatcher.GetStats()
	if stats.TotalReceived < 1 {
		t.Error("Message should be received")
	}
}

// TestDispatcherAdapter 测试分发器适配器。
// 使用独立实例而非全局单例：单例被 Stop 后 sync.Once 不会重建，
// -count 多轮运行时会因跨轮单例状态而失败。
func TestDispatcherAdapter(t *testing.T) {
	adapter := NewDispatcherAdapter(NewSdkIPad(), DefaultDispatcherConfig())
	if err := adapter.Start(); err != nil {
		t.Fatalf("adapter.Start failed: %v", err)
	}
	defer adapter.Stop()

	// 注册测试Hook
	hook := NewExampleHook("adapter-test-hook", 1)
	err := adapter.RegisterHook(LoginSuccessReturn, hook)
	if err != nil {
		t.Errorf("RegisterHook failed: %v", err)
	}

	// 创建测试消息
	msg := &CallBackMsg{
		UUID: "test-uuid-adapter",
		Type: LoginSuccessReturn,
	}

	// 处理回调
	err = adapter.HandleCallback(msg)
	if err != nil {
		t.Errorf("HandleCallback failed: %v", err)
	}

	// 获取分片统计
	shardStats := adapter.GetShardStats(0)
	if shardStats.TotalAccounts < 0 {
		t.Error("Invalid shard stats")
	}
}

// TestHookManager 测试Hook管理器
func TestHookManager(t *testing.T) {
	hm := NewHookManager()

	// 创建并注册多个Hook
	hook1 := NewExampleHook("hook-1", 1)
	hook2 := NewExampleHook("hook-2", 2)

	err := hm.RegisterHook(hook1)
	if err != nil {
		t.Errorf("RegisterHook 1 failed: %v", err)
	}

	err = hm.RegisterHook(hook2)
	if err != nil {
		t.Errorf("RegisterHook 2 failed: %v", err)
	}

	// 获取Hook
	retrievedHook := hm.GetHook("hook-1")
	if retrievedHook == nil {
		t.Error("GetHook failed")
	}

	// 获取所有Hook
	allHooks := hm.GetAllHooks()
	if len(allHooks) != 2 {
		t.Errorf("Expected 2 hooks, got %d", len(allHooks))
	}

	// 取消注册Hook
	err = hm.UnregisterHook("hook-1")
	if err != nil {
		t.Errorf("UnregisterHook failed: %v", err)
	}

	allHooks = hm.GetAllHooks()
	if len(allHooks) != 1 {
		t.Errorf("Expected 1 hook after unregister, got %d", len(allHooks))
	}
}

// TestPriorityQueue 测试优先级队列功能
func TestPriorityQueue(t *testing.T) {
	// 创建分发器进行优先级测试
	shardConfig := DefaultShardConfig()
	shardConfig.ShardCount = 2

	workerPoolConfig := DefaultWorkerPoolConfig()
	workerPoolConfig.MinWorkers = 1
	workerPoolConfig.MaxWorkers = 2

	dispatcherConfig := DefaultDispatcherConfig()

	sm := NewShardManager(shardConfig)
	pool := NewSharedWorkerPool(workerPoolConfig)

	dispatcher := NewCallbackDispatcher(nil, sm, pool, dispatcherConfig)
	if err := dispatcher.Start(); err != nil {
		t.Fatalf("dispatcher.Start failed: %v", err)
	}
	defer dispatcher.Stop()

	// 创建不同优先级的消息
	messages := []struct {
		uuid     string
		msgType  int
		priority int // 预期优先级
	}{
		{"uuid-1", ScanReturn, 2},      // 高优先级
		{"uuid-2", ApplyFriend, 1},     // 中优先级
		{"uuid-3", GroupNameChange, 0}, // 低优先级
	}

	// 注册测试Hook以验证处理顺序（mutex 计数收集，
	// 替代 channel 收集：close/send 在异步分发下必然竞态，-race 可复现）
	hook := &testPriorityHook{}
	if err := dispatcher.RegisterHook(ScanReturn, hook); err != nil {
		t.Fatalf("RegisterHook(ScanReturn) failed: %v", err)
	}
	if err := dispatcher.RegisterHook(ApplyFriend, hook); err != nil {
		t.Fatalf("RegisterHook(ApplyFriend) failed: %v", err)
	}
	if err := dispatcher.RegisterHook(GroupNameChange, hook); err != nil {
		t.Fatalf("RegisterHook(GroupNameChange) failed: %v", err)
	}

	// 发送消息
	for _, msg := range messages {
		msgObj := &CallBackMsg{
			UUID: msg.uuid,
			Type: msg.msgType,
		}
		err := dispatcher.HandleCallback(msgObj)
		if err != nil {
			t.Errorf("HandleCallback failed for %s: %v", msg.uuid, err)
		}
	}

	// 等待所有消息处理完成（轮询计数而非固定 Sleep：
	// 固定等待会与异步分发时序竞态，-race 模式可复现）
	deadline := time.Now().Add(2 * time.Second)
	for hook.Count() < len(messages) && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}

	// 理论上高优先级消息应该先处理，但由于并发性，顺序可能不严格
	// 这里主要验证所有消息都被处理了
	if count := hook.Count(); count != len(messages) {
		t.Errorf("Expected %d messages processed, got %d (handled: %v)",
			len(messages), count, hook.Handled())
	}
}

// testTaskHandler 测试任务处理器
type testTaskHandler struct{}

func (h *testTaskHandler) Handle(task *Task) error {
	// 简单的测试处理
	return nil
}

// testPriorityHook 测试优先级Hook：mutex 保护记录已处理消息，
// 通过 Count/Handled 线程安全读取（异步 worker 调用 Handle）。
type testPriorityHook struct {
	mu      sync.Mutex
	records []struct {
		uuid     string
		priority int
	}
}

func (h *testPriorityHook) Handle(ctx *HookContext, msg *CallBackMsg) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.records = append(h.records, struct {
		uuid     string
		priority int
	}{
		uuid:     msg.UUID,
		priority: ctx.AccountCtx.priority,
	})
	return nil
}

// Count 已处理消息数（线程安全）。
func (h *testPriorityHook) Count() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	return len(h.records)
}

// Handled 已处理消息的 uuid 列表（线程安全，调试输出用）。
func (h *testPriorityHook) Handled() []string {
	h.mu.Lock()
	defer h.mu.Unlock()
	uuids := make([]string, 0, len(h.records))
	for _, r := range h.records {
		uuids = append(uuids, r.uuid)
	}
	return uuids
}

func (h *testPriorityHook) Priority() int {
	return 1
}

func (h *testPriorityHook) Name() string {
	return "test-priority-hook"
}

// BenchmarkShardManager 分片管理器性能测试
func BenchmarkShardManager(b *testing.B) {
	config := DefaultShardConfig()
	config.ShardCount = 256
	sm := NewShardManager(config)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		uuid := string(byte(i%256)) + "-test-uuid"
		sm.GetOrCreateAccount(uuid)
	}
}

// BenchmarkSharedWorkerPool Worker池性能测试
func BenchmarkSharedWorkerPool(b *testing.B) {
	config := DefaultWorkerPoolConfig()
	config.MinWorkers = 4
	config.MaxWorkers = 16

	pool := NewSharedWorkerPool(config)
	defer pool.Stop()

	handler := &testTaskHandler{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		task := &Task{
			AccountUUID: "benchmark-uuid",
			Msg:         "benchmark message",
			Priority:    0,
			Handler:     handler,
			CreatedAt:   time.Now(),
			Timeout:     5 * time.Second,
		}
		_ = pool.Submit(task)
	}
}
