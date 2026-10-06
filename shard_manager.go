package sdk_ipad

// shard_manager.go 实现分片管理器。
//
// ShardManager 通过一致性哈希将账号分配到固定数量的分片中，
// 每个分片独立维护其下属账号的上下文信息。支持自动过期清理、
// 容量限制和统计信息采集。
//
// 架构: UUID → hash → Shard[N] → AccountContext

import (
	"hash/fnv"
	"sync"
	"time"
)

// ShardManager 分片管理器。
//
// 将大量账号按哈希值分散到多个 Shard 中，降低锁竞争。
// 内部启动定期清理协程，自动移除过期账号。
//
// 字段:
//   - shards: 固定数量的分片数组
//   - shardCount: 分片总数
//   - hashFunc: UUID → int 的哈希函数
//   - config: 分片配置
type ShardManager struct {
	shards     []*Shard         // 固定数量的分片
	shardCount int              // 分片数
	hashFunc   func(string) int // 哈希函数
	stopChan   chan struct{}    // 用于停止 cleanup goroutine

	// 配置
	config ShardConfig
}

// ShardConfig 分片配置。
//
// 字段:
//   - ShardCount: 分片数量，默认 256
//   - MaxAccounts: 全局最大账号数，默认 5000
//   - AccountTTL: 账号过期时间，默认 24h
//   - CleanupInterval: 清理任务执行间隔，默认 1h
type ShardConfig struct {
	ShardCount      int           // 分片数量，默认256
	MaxAccounts     int           // 最大账号数，默认5000
	AccountTTL      time.Duration // 账号TTL，默认24h
	CleanupInterval time.Duration // 清理间隔，默认1h
}

// DefaultShardConfig 返回默认分片配置。
//
// 返回: ShardConfig 包含推荐默认值的配置
func DefaultShardConfig() ShardConfig {
	return ShardConfig{
		ShardCount:      256,
		MaxAccounts:     5000,
		AccountTTL:      24 * time.Hour,
		CleanupInterval: time.Hour,
	}
}

// Shard 单个分片。
//
// 每个分片独立持有一把读写锁，管理一组账号上下文。
//
// 字段:
//   - id: 分片 ID
//   - accounts: UUID → AccountContext 映射
//   - lastActive: UUID → 最后活跃时间映射
//   - stats: 分片统计信息
type Shard struct {
	id         int
	accounts   map[string]*AccountContext // 轻量账号上下文
	mu         sync.RWMutex
	lastActive map[string]time.Time // 活跃时间
	stats      ShardStats
}

// ShardStats 分片统计信息。
//
// 字段:
//   - TotalMessages: 累计处理消息数
//   - TotalAccounts: 当前账号总数
//   - ActiveAccounts: 最近 5 分钟内活跃的账号数
//   - LastCleanup: 最后一次清理时间
type ShardStats struct {
	TotalMessages  int64
	TotalAccounts  int
	ActiveAccounts int
	LastCleanup    time.Time
}

// AccountContext 轻量账号上下文。
//
// 记录单个运行实例的状态和统计信息，由 ShardManager 管理生命周期。
//
// 字段:
//   - uuid: 实例唯一标识
//   - shardID: 所属分片 ID
//   - priority: 当前动态优先级
//   - stats: 账号统计信息
//   - lastMsgTime: 最后收到消息的时间
//   - createdAt: 账号上下文创建时间
type AccountContext struct {
	uuid        string
	shardID     int
	priority    int          // 当前优先级（动态调整）
	stats       AccountStats // 统计信息
	lastMsgTime time.Time    // 最后消息时间
	createdAt   time.Time    // 创建时间
}

// AccountStats 账号统计信息。
//
// 字段:
//   - TotalMessages: 累计消息数
//   - SuccessCount: 成功处理数
//   - ErrorCount: 失败处理数
//   - AvgLatency: 平均处理延迟（移动平均）
type AccountStats struct {
	TotalMessages int64
	SuccessCount  int64
	ErrorCount    int64
	AvgLatency    time.Duration
}

// NewShardManager 创建分片管理器。
//
// 初始化指定数量的分片，并启动后台定期清理协程。
// 如果 ShardCount <= 0，则使用默认值 256。
//
// 参数:
//   - config: 分片配置
//
// 返回: *ShardManager 初始化完成的分片管理器
func NewShardManager(config ShardConfig) *ShardManager {
	if config.ShardCount <= 0 {
		config.ShardCount = 256
	}

	shards := make([]*Shard, config.ShardCount)
	for i := 0; i < config.ShardCount; i++ {
		shards[i] = &Shard{
			id:         i,
			accounts:   make(map[string]*AccountContext),
			lastActive: make(map[string]time.Time),
			stats:      ShardStats{},
		}
	}

	sm := &ShardManager{
		shards:     shards,
		shardCount: config.ShardCount,
		hashFunc:   defaultHashFunc,
		config:     config,
		stopChan:   make(chan struct{}),
	}

	// 启动定期清理任务
	go sm.startCleanupTask()

	return sm
}

// Stop 停止分片管理器。
//
// 关闭后台清理协程。调用后不应继续使用该实例。
func (sm *ShardManager) Stop() {
	close(sm.stopChan)
}

// defaultHashFunc 默认哈希函数。
//
// 使用 FNV-1a 算法将 UUID 转换为 int。
// 注意：不在此处取模，由调用方根据 shardCount 取模。
//
// 参数:
//   - uuid: 实例唯一标识
//
// 返回: int 哈希值
func defaultHashFunc(uuid string) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(uuid)) // hash.Write 永不返回错误，显式忽略以满足 gosec G104
	return int(h.Sum32())
}

// GetOrCreateAccount 获取或创建账号上下文。
//
// 优先使用读锁查找，未命中时升级为写锁并创建。
// 如果分片已满，会触发强制清理过期账号以腾出空间。
//
// 参数:
//   - uuid: 实例唯一标识
//
// 返回: *AccountContext 账号上下文（始终非 nil）
func (sm *ShardManager) GetOrCreateAccount(uuid string) *AccountContext {
	shardID := sm.hashFunc(uuid) % sm.shardCount
	shard := sm.shards[shardID]

	// 先用读锁尝试获取
	shard.mu.RLock()
	ctx, exists := shard.accounts[uuid]
	shard.mu.RUnlock()

	if exists {
		// Fix #2: 更新 lastActive 需要写锁保护
		shard.mu.Lock()
		shard.lastActive[uuid] = time.Now()
		shard.mu.Unlock()
		return ctx
	}

	// 创建新的账号上下文
	shard.mu.Lock()
	defer shard.mu.Unlock()

	// 双重检查
	if ctx, exists = shard.accounts[uuid]; exists {
		shard.lastActive[uuid] = time.Now()
		return ctx
	}

	// 检查分片容量
	if len(shard.accounts) >= sm.config.MaxAccounts/sm.shardCount {
		// Fix #1: 调用 cleanupShardLocked（假设锁已持有），避免死锁
		sm.cleanupShardLocked(shard, true)
	}

	ctx = &AccountContext{
		uuid:        uuid,
		shardID:     shardID,
		priority:    0,
		stats:       AccountStats{},
		lastMsgTime: time.Now(),
		createdAt:   time.Now(),
	}

	shard.accounts[uuid] = ctx
	shard.lastActive[uuid] = time.Now()
	shard.stats.TotalAccounts++

	return ctx
}

// GetAccount 获取账号上下文（只读）。
//
// 如果指定 UUID 的账号不存在，返回 nil。
//
// 参数:
//   - uuid: 实例唯一标识
//
// 返回: *AccountContext 账号上下文，不存在时返回 nil
func (sm *ShardManager) GetAccount(uuid string) *AccountContext {
	shardID := sm.hashFunc(uuid) % sm.shardCount
	shard := sm.shards[shardID]

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	return shard.accounts[uuid]
}

// UpdateAccountStats 更新账号统计信息。
//
// 原子更新消息计数和平均延迟（移动平均，新值占 20%）。
// 如果账号不存在则静默忽略。
//
// 参数:
//   - uuid: 实例唯一标识
//   - success: 本次处理是否成功
//   - latency: 本次处理延迟
func (sm *ShardManager) UpdateAccountStats(uuid string, success bool, latency time.Duration) {
	shardID := sm.hashFunc(uuid) % sm.shardCount
	shard := sm.shards[shardID]

	shard.mu.Lock()
	defer shard.mu.Unlock()

	ctx, exists := shard.accounts[uuid]
	if !exists {
		return
	}

	ctx.stats.TotalMessages++
	if success {
		ctx.stats.SuccessCount++
	} else {
		ctx.stats.ErrorCount++
	}

	// 更新平均延迟（移动平均）
	if ctx.stats.AvgLatency == 0 {
		ctx.stats.AvgLatency = latency
	} else {
		// 简单移动平均：新值占20%
		ctx.stats.AvgLatency = time.Duration(0.8*float64(ctx.stats.AvgLatency) + 0.2*float64(latency))
	}

	ctx.lastMsgTime = time.Now()
	shard.lastActive[uuid] = ctx.lastMsgTime
	shard.stats.TotalMessages++
}

// GetShardStats 获取指定分片的统计信息。
//
// 统计活跃账号数（最近 5 分钟内有活动的账号）。
// 如果 shardID 越界，返回空的 ShardStats。
//
// 参数:
//   - shardID: 分片 ID（0 ~ shardCount-1）
//
// 返回: ShardStats 分片统计信息
func (sm *ShardManager) GetShardStats(shardID int) ShardStats {
	if shardID < 0 || shardID >= sm.shardCount {
		return ShardStats{}
	}

	shard := sm.shards[shardID]
	shard.mu.RLock()
	defer shard.mu.RUnlock()

	// 计算活跃账号数
	activeAccounts := 0
	now := time.Now()
	for _, lastActive := range shard.lastActive {
		if now.Sub(lastActive) < 5*time.Minute {
			activeAccounts++
		}
	}

	stats := shard.stats
	stats.ActiveAccounts = activeAccounts

	return stats
}

// GetAllStats 获取所有分片的统计信息。
//
// 返回: []ShardStats 所有分片统计信息列表，索引即分片 ID
func (sm *ShardManager) GetAllStats() []ShardStats {
	stats := make([]ShardStats, sm.shardCount)
	for i := 0; i < sm.shardCount; i++ {
		stats[i] = sm.GetShardStats(i)
	}
	return stats
}

// cleanupShard 清理分片中的过期账号（自行获取写锁）。
//
// 参数:
//   - shard: 目标分片
//   - force: 是否强制清理（使用更短的 TTL: 1 小时）
func (sm *ShardManager) cleanupShard(shard *Shard, force bool) {
	shard.mu.Lock()
	defer shard.mu.Unlock()

	sm.cleanupShardLocked(shard, force)
}

// cleanupShardLocked 清理分片中的过期账号（调用方必须持有 shard.mu 写锁）。
//
// 遍历 lastActive 映射，删除超过 TTL 的账号上下文。
// force 模式使用 1 小时作为过期阈值，用于容量不足时的紧急清理。
//
// 参数:
//   - shard: 目标分片
//   - force: 是否强制清理
func (sm *ShardManager) cleanupShardLocked(shard *Shard, force bool) {
	now := time.Now()
	expiredTime := now.Add(-sm.config.AccountTTL)

	// 如果是强制清理，使用更短的TTL
	if force {
		expiredTime = now.Add(-time.Hour)
	}

	removed := 0
	for uuid, lastActive := range shard.lastActive {
		if lastActive.Before(expiredTime) {
			delete(shard.accounts, uuid)
			delete(shard.lastActive, uuid)
			removed++
		}
	}

	if removed > 0 {
		shard.stats.TotalAccounts -= removed
		shard.stats.LastCleanup = now
	}
}

// startCleanupTask 启动定期清理后台任务。
//
// 按 config.CleanupInterval 间隔遍历所有分片执行过期清理。
// 通过 stopChan 信号终止。
func (sm *ShardManager) startCleanupTask() {
	ticker := time.NewTicker(sm.config.CleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, shard := range sm.shards {
				sm.cleanupShard(shard, false)
			}
		case <-sm.stopChan:
			return
		}
	}
}

// HandleCallback 处理回调消息（路由到对应分片）。
//
// 根据 UUID 哈希找到对应分片，获取/创建账号上下文，
// 执行 handler 回调函数，并更新账号统计。
//
// 参数:
//   - msg: 回调消息
//   - handler: 处理函数，接收账号上下文和回调消息
//
// 返回: error UUID 为空返回 ErrInvalidUUID，或 handler 返回的错误
func (sm *ShardManager) HandleCallback(msg *CallBackMsg, handler func(*AccountContext, *CallBackMsg) error) error {
	if msg.UUID == "" {
		return ErrInvalidUUID
	}

	ctx := sm.GetOrCreateAccount(msg.UUID)
	startTime := time.Now()

	err := handler(ctx, msg)
	latency := time.Since(startTime)

	sm.UpdateAccountStats(msg.UUID, err == nil, latency)

	return err
}

// ErrInvalidUUID UUID 无效错误。
var ErrInvalidUUID = &Error{Code: -1, Message: "invalid uuid"}

// Error 自定义错误类型。
//
// 包含错误码和错误消息，用于 SDK 内部统一错误返回。
type Error struct {
	Code    int
	Message string
}

// Error 实现 error 接口。
func (e *Error) Error() string {
	return e.Message
}
