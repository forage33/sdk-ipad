package sdk_ipad

import (
	"testing"
	"time"
)

// shard_manager_test.go 分片管理器（shard_manager.go）的功能与性能测试。

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
