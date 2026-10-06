// msg_types.go 聊天消息类型常量定义。
// 定义企业微信回调中聊天消息的 msgtype 字段值。

package sdk_ipad

// 企业微信回调聊天消息类型（msgtype 字段）。
const (
	ChatMsgSystem         = 1011 // 系统消息
	ChatMsgText           = 0    // 文本消息
	ChatMsgTextAlt        = 2    // 文本消息（备用类型）
	ChatMsgGif            = 104  // GIF 表情消息
	ChatMsgGifDynamic     = 29   // 动态 GIF 表情
	ChatMsgMiniProgram    = 78   // 小程序消息
	ChatMsgLink           = 13   // 链接消息
	ChatMsgVideoChannel   = 141  // 视频号消息
	ChatMsgCard           = 41   // 名片消息
	ChatMsgLocation       = 6    // 位置消息
	ChatMsgExternalImage  = 101  // 外部联系人图片消息（个微发的）
	ChatMsgRedPacket      = 26   // 红包消息
	ChatMsgInternalImage  = 14   // 内部联系人图片消息（企微用户）
	ChatMsgExternalFile   = 102  // 外部联系人文件消息
	ChatMsgVoice          = 16   // 语音消息
	ChatMsgInternalFile   = 15   // 内部文件消息
	ChatMsgInternalVideo  = 23   // 内部视频消息
	ChatMsgExternalVideo  = 103  // 外部视频消息
	ChatMsgVoiceVideoCall = 503  // 语音/视频通话消息
	ChatMsgCallCancel     = 40   // 语音或视频通话取消
	ChatMsgInternalMerge  = 123  // 内部联系人合并消息
	ChatMsgRevoke         = 2063 // 撤回消息
	ChatMsgMerge          = 4    // 合并消息
	ChatMsgChainRelay     = 213  // 接龙消息
)
