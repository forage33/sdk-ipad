package sdk_ipad

import "testing"

// hook_manager_test.go Hook 管理器（example_hooks.go）的功能测试。

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
