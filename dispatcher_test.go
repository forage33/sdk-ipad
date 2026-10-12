package sdk_ipad

import (
	"sync"
	"testing"
	"time"
)

// dispatcher_test.go 回调分发器（callback_dispatcher.go）的功能测试。

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
