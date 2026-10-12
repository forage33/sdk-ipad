// Package sdk_ipad 提供企业微信 iPad 协议 SDK 的 Go 封装。
//
// 本包通过 HTTP 调用后端协议服务，实现企业微信 iPad 端的各项功能操作，
// 同时提供回调消息分发、分片管理、Worker 池等基础设施。
//
// 一、主动调用 API（编号与 docs/api 域一致，按方法名前缀可快速定位文件）：
//
//  01. 初始化       - api_wx_init.go          InItWXWork/SetCallBackURL/GetRunClient 等
//  02. 登录         - api_login.go            GetQrCode/CheckCode/AutomaticLogin 等
//  03. 账号操作     - api_account.go          头像/资料/二维码/封禁/企业信息
//  04. 开放平台     - 未实现
//  05. 标签         - api_label.go            *Label* 企业标签增删改查
//  06. CDN          - api_cdn.go              Cdn*/Upload*/Download*/*BigFile* 上传下载
//  07. 联系人操作   - api_contact_add.go      搜索/添加/申请/共享（加好友链路）
//     api_contact_manage.go   资料/列表/设置/拉黑/删除
//  08. 朋友圈       - api_moments.go          *Sns*/PostPyq
//  09. 控制台群发   - api_broadcast.go        ConfirmGroupSend
//  10. 聊天标签     - api_chat_tag.go         ModSessionTag
//  11. 客服管理     - 未实现
//  12. 商家收款     - 未实现
//  13. 发送消息     - api_message_send.go     Send* 全部发送接口
//     api_message_action.go   RevokeMsg/MarkAsRead/SpeechToTextEntity/SyncAllData
//  14. 群操作       - api_group.go            *Room*/*Group* 创建/邀请/公告/二维码/欢迎语等
//
// 二、协议服务回调消息链路：
//
//	收包     callback_base.go（CallBackMsg 信封 + 回调类型码常量）
//	传输     callback_http.go（CallbackHTTPHandler：解析 → 分发 → 失败异步重试）
//	分发     callback_dispatcher.go / dispatcher_adapter.go（Hook 注册、分片路由、Worker 池）
//	契约     callback_payload_msg.go（消息域 As* 访问器）/ callback_payload_account.go（登录/群/联系人/断连域）
//	消息类型 chat_msgtype.go（msgtype 常量：ChatMsgText/Voice/Video 等）
//
// 回调域覆盖：登录/消息/群/联系人/控制台任务/标签变动/聊天标签/断开。
package sdk_ipad

import "go.uber.org/zap"

// SdkIPad 企业微信 iPad 协议 SDK 主入口。
// 封装了 HTTP 客户端、分片管理器和全局配置，对外提供各业务模块的 API 调用方法。
type SdkIPad struct {
	httpClient   HTTPClient    // HTTP 客户端
	shardManager *ShardManager // 分片管理器
	config       *SdkConfig    // SDK 配置
	logger       *zap.Logger   // 日志器，由 SdkConfig.Logger 注入，未注入时回退全局 zap.L()
}

// SdkConfig SDK 全局配置，包含分片、Worker 池和优先级队列的配置项。
type SdkConfig struct {
	ShardConfig   ShardConfig         // 分片配置
	WorkerPool    WorkerPoolConfig    // Worker 池配置
	PriorityQueue PriorityQueueConfig // 优先级队列配置（预留：调度未实现，见 PriorityQueueConfig 注释）

	// Logger 日志器（zap Field 风格），由宿主应用注入；为 nil 时回退全局 zap.L()。
	// SDK 不依赖任何宿主项目的日志封装，以保证本包可独立编译复用。
	Logger *zap.Logger
}

// DefaultSdkConfig 返回 SDK 默认全局配置。
func DefaultSdkConfig() *SdkConfig {
	return &SdkConfig{
		ShardConfig:   DefaultShardConfig(),
		WorkerPool:    DefaultWorkerPoolConfig(),
		PriorityQueue: DefaultPriorityQueueConfig(),
	}
}

// NewSdkIPad 使用默认配置创建 SDK 实例。
//
// 返回: *SdkIPad 实例
func NewSdkIPad() *SdkIPad {
	return NewSdkIPadWithConfig(DefaultSdkConfig())
}

// NewSdkIPadWithConfig 使用自定义配置创建 SDK 实例。
//
// 参数:
//   - config: SDK 配置，为 nil 时使用默认配置
//
// 返回: *SdkIPad 实例
func NewSdkIPadWithConfig(config *SdkConfig) *SdkIPad {
	if config == nil {
		config = DefaultSdkConfig()
	}

	logger := config.Logger
	if logger == nil {
		logger = zap.L()
	}

	client := NewDefaultHTTPClient(NewOptimizedHTTPClient())
	shardManager := NewShardManager(config.ShardConfig)

	return &SdkIPad{
		httpClient:   client,
		shardManager: shardManager,
		config:       config,
		logger:       logger,
	}
}

// GetShardManager 获取分片管理器实例。
//
// 返回: *ShardManager 分片管理器
func (s *SdkIPad) GetShardManager() *ShardManager {
	return s.shardManager
}

// GetAccountContext 根据 UUID 获取账号上下文，不存在时返回 nil。
//
// 参数:
//   - uuid: 实例唯一标识
//
// 返回: *AccountContext 账号上下文，不存在时为 nil
func (s *SdkIPad) GetAccountContext(uuid string) *AccountContext {
	return s.shardManager.GetAccount(uuid)
}

// GetOrCreateAccountContext 根据 UUID 获取或创建账号上下文。
// 如果该 UUID 对应的上下文不存在，则自动创建。
//
// 参数:
//   - uuid: 实例唯一标识
//
// 返回: *AccountContext 账号上下文
func (s *SdkIPad) GetOrCreateAccountContext(uuid string) *AccountContext {
	return s.shardManager.GetOrCreateAccount(uuid)
}

// GetShardStats 获取指定分片的统计信息。
//
// 参数:
//   - shardID: 分片 ID
//
// 返回: ShardStats 分片统计数据
func (s *SdkIPad) GetShardStats(shardID int) ShardStats {
	return s.shardManager.GetShardStats(shardID)
}

// GetAllShardStats 获取所有分片的统计信息。
//
// 返回: []ShardStats 各分片统计数据切片
func (s *SdkIPad) GetAllShardStats() []ShardStats {
	return s.shardManager.GetAllStats()
}

// HandleCallback 处理协议服务器推送的回调消息。
// 根据消息中的 UUID 路由到对应的分片，并调用 handler 处理。
//
// 参数:
//   - msg: 回调消息体
//   - handler: 消息处理函数，接收账号上下文和消息
//
// 返回: error 处理过程中的错误
func (s *SdkIPad) HandleCallback(msg *CallBackMsg, handler func(*AccountContext, *CallBackMsg) error) error {
	return s.shardManager.HandleCallback(msg, handler)
}

// 协议服务已知错误码参考表（仅供开发参考，非代码常量）
//
//	CODE_86001(86001, "非法群ID")
//	CODE_86216(86216, "非法群成员(请确认账号是否在群内)")
//	CODE_2(2, "当前客户端版本较低，请升级到最新版本后再试")
//	SUCCESS(0, "ok")
//	ERR_1(-1, "")
//	ERR_9(-9, "")
//	ERR_11(-11, "操作失败请重试")
//	ERR_104(-104, "请求失败：无群邀请开关权限")
//	ERR_202(-202, "请求失败：请在手机上操作复现后反馈管理员")
//	ERR_401(-401, "登录过期")
//	ERR_1007(-1007, "")
//	ERR_2001(-2001, "参数错误:(添加群成员为好友时群ID错误)")
//	ERR_2003(-2003, "参数错误:(群ID为0)")
//	ERR_2007(-2007, "扫码授权过期，请重新登录")
//	ERR_2009(-2009, "帐号环境出现异常，请重新登录使用")
//	ERR_2013(-2013, "目标好友非单向好友")
//	ERR_2014(-2014, "请求失败：对方无外部联系人权限")
//	ERR_2016(-2016, "请求失败：添加好友前请先完成实名认证")
//	ERR_2024(-2024, "请求失败：无外部联系人权限")
//	ERR_2038(-2038, "你的客户数已达上限，无法添加客户")
//	ERR_2054(-2054, "添加请求已过期，添加失败")
//	ERR_2068(-2068, "未认证企业仅可添加200位外部联系人")
//	ERR_3004(-3004, "请求失败：已不在群会话中")
//	ERR_3020(-3020, "请求失败：群不存在")
//	ERR_4018(-4018, "对方已离职，无法发送消息")
//	ERR_4027(-4027, "消息撤回失败(微信或群内含微信号只能撤回2分钟内的消息)")
//	ERR_4048(-4048, "请求失败：消息发送过于频繁，请稍后再试")
//	ERR_11001(-11001, "登录态过期，请重新登录")
//	ERR_11025(-11025, "新设备扫码授权验证")
//	ERR_12022(-12022, "验证码错误")
//	ERR_18000039(-18000039, "添加好友过于频繁，请稍后再重试")
//	ERR_18000043(-18000043, "对方不允许通过群聊添加好友")
//	ERR_99999998(-99999998, "官方服务端错误，请稍后重试")
//	ERR_99999999(-99999999, "发送消息失败，服务器错误，请重试")
//
// 完整错误码列表请以协议服务文档为准。
