package sdk_ipad

// shared_pool.go 实现共享 Worker 池。
//
// SharedWorkerPool 是一个支持动态扩缩容的 goroutine 工作池，
// 用于并发处理回调消息任务。支持优先级队列配置、自动扩缩容、
// 提交超时和原子统计。
//
// 架构: Submit(Task) → taskChan → worker goroutine → TaskHandler.Handle

import (
	"sync"
	"sync/atomic"
	"time"
)

// WorkerPoolConfig Worker 池配置。
//
// 字段:
//   - MinWorkers: 最小 worker 数，默认 4（CPU 核心数）
//   - MaxWorkers: 最大 worker 数，默认 16（MinWorkers×4）
//   - QueueCapacity: 任务队列容量，默认 10000
//   - WorkerIdleTime: worker 空闲超时，超时后可被缩容回收，默认 30s
//   - ScaleUpInterval: 扩缩容检查间隔，默认 1s
type WorkerPoolConfig struct {
	MinWorkers      int           // 最小worker数，默认CPU核心数
	MaxWorkers      int           // 最大worker数，默认CPU核心数×4
	QueueCapacity   int           // 任务队列容量，默认10000
	WorkerIdleTime  time.Duration // worker空闲时间，默认30秒
	ScaleUpInterval time.Duration // 扩容检查间隔，默认1秒
}

// DefaultWorkerPoolConfig 返回默认 Worker 池配置。
//
// 返回: WorkerPoolConfig 包含推荐默认值的配置
func DefaultWorkerPoolConfig() WorkerPoolConfig {
	return WorkerPoolConfig{
		MinWorkers:      4,
		MaxWorkers:      16,
		QueueCapacity:   10000,
		WorkerIdleTime:  30 * time.Second,
		ScaleUpInterval: time.Second,
	}
}

// PriorityQueueConfig 优先级队列配置。
//
// 字段:
//   - Priorities: 优先级数量，默认 3
//   - CapacityPerPri: 每个优先级队列容量，默认 10000
//   - HighPriorityCap: 高优先级队列容量，默认 20000
type PriorityQueueConfig struct {
	Priorities      int // 优先级数量，默认3
	CapacityPerPri  int // 每个优先级队列容量，默认10000
	HighPriorityCap int // 高优先级队列容量，默认20000
}

// DefaultPriorityQueueConfig 返回默认优先级队列配置。
//
// 返回: PriorityQueueConfig 包含推荐默认值的配置
func DefaultPriorityQueueConfig() PriorityQueueConfig {
	return PriorityQueueConfig{
		Priorities:      3,
		CapacityPerPri:  10000,
		HighPriorityCap: 20000,
	}
}

// Task 任务结构体。
//
// 封装提交到 Worker 池的单个工作单元。
//
// 字段:
//   - AccountUUID: 关联的账号 UUID
//   - Msg: 原始消息（interface{}）
//   - Priority: 优先级（0-低，1-中，2-高）
//   - Handler: 任务处理器
//   - CreatedAt: 任务创建时间
//   - Timeout: 任务处理超时时间
type Task struct {
	AccountUUID string
	Msg         interface{}
	Priority    int           // 0-低，1-中，2-高
	Handler     TaskHandler   // 任务处理器
	CreatedAt   time.Time     // 创建时间
	Timeout     time.Duration // 超时时间
}

// TaskHandler 任务处理器接口。
//
// 由具体的业务处理器实现，Worker 取到任务后调用 Handle 方法执行。
type TaskHandler interface {
	Handle(task *Task) error
}

// SharedWorkerPool 共享 Worker 池。
//
// 管理一组 worker goroutine，通过 buffered channel 分发任务。
// 支持动态扩缩容：队列积压时自动扩容，空闲时自动缩容。
//
// 字段:
//   - config: Worker 池配置
//   - taskChan: 任务分发 channel
//   - workers: 当前活跃的 worker 列表
//   - stats: 运行时统计（原子操作）
type SharedWorkerPool struct {
	config    WorkerPoolConfig
	taskChan  chan *Task // 任务通道
	workers   []*worker  // worker列表
	mu        sync.Mutex
	isRunning atomic.Bool
	stopOnce  sync.Once
	stopChan  chan struct{}
	stats     PoolStats
}

// PoolStats Worker 池运行时统计。
//
// 所有数值字段均通过原子操作更新。
//
// 字段:
//   - TotalTasks: 累计提交的任务数
//   - ProcessedTasks: 累计成功处理数
//   - FailedTasks: 累计失败数
//   - ActiveWorkers: 当前活跃 worker 数
//   - QueueLength: 当前队列中的任务数
//   - AvgProcessTime: 平均处理耗时（纳秒，移动平均）
type PoolStats struct {
	TotalTasks     int64
	ProcessedTasks int64
	FailedTasks    int64
	ActiveWorkers  int32
	QueueLength    int32
	AvgProcessTime int64 // 纳秒，使用 atomic 操作
}

// worker 单个工作协程。
//
// 从 pool.taskChan 获取任务并执行，空闲超时后可被缩容回收。
type worker struct {
	id        int
	pool      *SharedWorkerPool
	taskChan  chan *Task
	stopChan  chan struct{}
	stopOnce  sync.Once // Fix #7: 用 sync.Once 保证只 close 一次
	lastWork  time.Time
	isWorking atomic.Bool
}

// NewSharedWorkerPool 创建共享 Worker 池。
//
// 初始化 MinWorkers 个 worker goroutine，并启动后台扩缩容监控。
// 零值配置字段将使用默认值。
//
// 参数:
//   - config: Worker 池配置
//
// 返回: *SharedWorkerPool 初始化完成且已启动的 Worker 池
func NewSharedWorkerPool(config WorkerPoolConfig) *SharedWorkerPool {
	if config.MinWorkers <= 0 {
		config.MinWorkers = 4
	}
	if config.MaxWorkers < config.MinWorkers {
		config.MaxWorkers = config.MinWorkers * 4
	}
	if config.QueueCapacity <= 0 {
		config.QueueCapacity = 10000
	}

	pool := &SharedWorkerPool{
		config:   config,
		taskChan: make(chan *Task, config.QueueCapacity),
		workers:  make([]*worker, 0, config.MaxWorkers),
		stopChan: make(chan struct{}),
		stats:    PoolStats{},
	}

	// 初始化最小worker数
	for i := 0; i < config.MinWorkers; i++ {
		pool.addWorker()
	}

	pool.isRunning.Store(true)
	go pool.monitorAndScale()

	return pool
}

// addWorker 添加一个 worker goroutine。
//
// 如果已达到 MaxWorkers 上限则不添加。线程安全。
func (p *SharedWorkerPool) addWorker() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.workers) >= p.config.MaxWorkers {
		return
	}

	id := len(p.workers)
	w := &worker{
		id:       id,
		pool:     p,
		taskChan: make(chan *Task, 1),
		stopChan: make(chan struct{}),
		lastWork: time.Now(),
	}

	p.workers = append(p.workers, w)
	go w.start()

	atomic.AddInt32(&p.stats.ActiveWorkers, 1)
}

// removeIdleWorker 移除最空闲的 worker。
//
// 查找空闲时间最长且超过 WorkerIdleTime 的 worker 进行回收。
// 不会低于 MinWorkers 数量。线程安全。
func (p *SharedWorkerPool) removeIdleWorker() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.workers) <= p.config.MinWorkers {
		return
	}

	// 查找最空闲的worker
	idleIndex := -1
	maxIdleTime := time.Duration(0)
	now := time.Now()

	for i, w := range p.workers {
		if w.isWorking.Load() {
			continue
		}

		idleTime := now.Sub(w.lastWork)
		if idleTime > maxIdleTime && idleTime > p.config.WorkerIdleTime {
			maxIdleTime = idleTime
			idleIndex = i
		}
	}

	if idleIndex == -1 {
		return
	}

	// Fix #7: 使用 sync.Once 停止 worker，避免 double-close
	p.workers[idleIndex].stop()

	// 从切片中移除
	p.workers = append(p.workers[:idleIndex], p.workers[idleIndex+1:]...)
	atomic.AddInt32(&p.stats.ActiveWorkers, -1)
}

// Submit 提交任务到 Worker 池。
//
// 将任务写入 taskChan，如果 500ms 内无法写入则返回超时错误。
// Worker 池未运行时返回 ErrPoolNotRunning。
//
// 参数:
//   - task: 待执行的任务
//
// 返回: error 池未运行返回 ErrPoolNotRunning，超时返回 ErrSubmitTimeout
func (p *SharedWorkerPool) Submit(task *Task) error {
	if !p.isRunning.Load() {
		return ErrPoolNotRunning
	}

	atomic.AddInt64(&p.stats.TotalTasks, 1)

	select {
	case p.taskChan <- task:
		atomic.AddInt32(&p.stats.QueueLength, 1)
		return nil
	case <-time.After(500 * time.Millisecond):
		return ErrSubmitTimeout
	}
}

// monitorAndScale 后台扩缩容监控协程。
//
// 按 ScaleUpInterval 间隔检查队列积压和 worker 空闲情况：
//   - 队列长度 > 活跃 worker 数 × 2 → 扩容
//   - worker 数 > MinWorkers 且有空闲超时的 worker → 缩容
func (p *SharedWorkerPool) monitorAndScale() {
	ticker := time.NewTicker(p.config.ScaleUpInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if !p.isRunning.Load() {
				return
			}

			// 检查是否需要扩容：使用 channel 实际长度
			queueLen := len(p.taskChan)
			activeWorkers := int(atomic.LoadInt32(&p.stats.ActiveWorkers))

			// 如果队列长度超过worker数×2，且还有扩容空间，则扩容
			if queueLen > activeWorkers*2 {
				p.mu.Lock()
				needScale := len(p.workers) < p.config.MaxWorkers
				p.mu.Unlock()
				if needScale {
					p.addWorker()
				}
			}

			// 检查是否需要缩容
			p.mu.Lock()
			needShrink := len(p.workers) > p.config.MinWorkers
			p.mu.Unlock()
			if needShrink {
				p.removeIdleWorker()
			}

		case <-p.stopChan:
			return
		}
	}
}

// Stop 停止 Worker 池。
//
// 先标记停止状态拒绝新任务，等待队列中任务清空（最多 10s），再关闭 worker。
// 使用 sync.Once 保证只执行一次。
func (p *SharedWorkerPool) Stop() {
	p.stopOnce.Do(func() {
		p.isRunning.Store(false)

		// 优雅 drain：等待队列清空或超时 10s
		timer := time.NewTimer(10 * time.Second)
		defer timer.Stop()
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
	drainLoop:
		for {
			if len(p.taskChan) == 0 {
				break drainLoop
			}
			select {
			case <-timer.C:
				break drainLoop
			case <-ticker.C:
			}
		}

		close(p.stopChan)

		p.mu.Lock()
		defer p.mu.Unlock()

		// Fix #7: 使用每个 worker 的 stop 方法安全关闭
		for _, w := range p.workers {
			w.stop()
		}
		p.workers = nil
	})
}

// GetStats 获取 Worker 池运行时统计（线程安全快照）。
//
// 返回: PoolStats 统计信息副本
func (p *SharedWorkerPool) GetStats() PoolStats {
	return PoolStats{
		TotalTasks:     atomic.LoadInt64(&p.stats.TotalTasks),
		ProcessedTasks: atomic.LoadInt64(&p.stats.ProcessedTasks),
		FailedTasks:    atomic.LoadInt64(&p.stats.FailedTasks),
		ActiveWorkers:  atomic.LoadInt32(&p.stats.ActiveWorkers),
		QueueLength:    atomic.LoadInt32(&p.stats.QueueLength),
		AvgProcessTime: atomic.LoadInt64(&p.stats.AvgProcessTime),
	}
}

// stop 安全停止 worker goroutine。
//
// 使用 sync.Once 保证 stopChan 只关闭一次。
func (w *worker) stop() {
	w.stopOnce.Do(func() {
		close(w.stopChan)
	})
}

// start worker 主循环。
//
// 从 pool.taskChan 获取任务，调用 Handler.Handle 执行，
// 并通过原子操作更新统计（处理数、失败数、平均耗时）。
// 收到 stopChan 信号后退出。
func (w *worker) start() {
	for {
		select {
		case task := <-w.pool.taskChan:
			atomic.AddInt32(&w.pool.stats.QueueLength, -1)
			w.isWorking.Store(true)
			w.lastWork = time.Now()

			startTime := time.Now()
			err := task.Handler.Handle(task)
			processTime := time.Since(startTime)

			if err != nil {
				atomic.AddInt64(&w.pool.stats.FailedTasks, 1)
			} else {
				atomic.AddInt64(&w.pool.stats.ProcessedTasks, 1)
			}

			// Fix #6: 使用 CAS 循环原子更新平均处理时间
			for {
				oldAvg := atomic.LoadInt64(&w.pool.stats.AvgProcessTime)
				var newAvg int64
				if oldAvg == 0 {
					newAvg = int64(processTime)
				} else {
					newAvg = int64(0.8*float64(oldAvg) + 0.2*float64(processTime))
				}
				if atomic.CompareAndSwapInt64(&w.pool.stats.AvgProcessTime, oldAvg, newAvg) {
					break
				}
			}

			w.isWorking.Store(false)

		case <-w.stopChan:
			return
		}
	}
}

// Worker 池相关错误定义。
var (
	ErrPoolNotRunning = &Error{Code: -1001, Message: "worker pool not running"} // Worker 池未运行
	ErrQueueFull      = &Error{Code: -1002, Message: "task queue is full"}      // 任务队列已满
	ErrSubmitTimeout  = &Error{Code: -1003, Message: "submit task timeout"}     // 提交任务超时
)
