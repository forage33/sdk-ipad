package sdk_ipad

import "testing"

// dispatcher_adapter_test.go 分发器适配器（dispatcher_adapter.go）的功能测试。

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
