package sdk_ipad

import (
	"testing"
	"time"
)

// shared_pool_test.go 共享 Worker 池（shared_pool.go）的功能与性能测试。

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

// testTaskHandler 测试任务处理器
type testTaskHandler struct{}

func (h *testTaskHandler) Handle(task *Task) error {
	// 简单的测试处理
	return nil
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
