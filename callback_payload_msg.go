// callback_payload_msg.go 消息回调（MsgCallback 102000）的数据契约层。
//
// 本文件是协议服务下发消息回调的强类型 payload 定义与类型化访问器，
// 字段名与协议文档 docs/callback/02-message.md 一一对应；协议字段知识
// （字段名、双名兼容、多级回退）全部收敛于此，业务层不再直接接触
// 原始 JSON 字符串字段。
//
// 解析语义约定（与历史业务实现严格等价，迁移时逐处核对过）：
//   - 消息类回调字段解析使用 gjson 宽容语义：字段缺失/类型不匹配时
//     返回零值而非报错（与原业务层逐字段 gjson.Get 行为一致）；
//   - 数字字段的字符串视图（如撤回回调的 sender/receiver）保持
//     gjson Result.String() 语义：数字返回十进制串、缺失返回空串，
//     业务侧依赖空串判断字段缺失时行为不变。

package sdk_ipad

import "github.com/tidwall/gjson"

// ---------------------------------------------------------------------------
// 通用信封
// ---------------------------------------------------------------------------

// MessageEnvelope 聊天消息回调（102000）的通用信封字段，所有子类型共享。
type MessageEnvelope struct {
	MsgType            int64  // msgtype 原生消息类型（ChatMsg* 常量）
	Sender             int64  // sender 发送人 VID
	Receiver           int64  // receiver 接收人 VID（群聊/系统类回调可能为 0）
	IsRoom             int64  // is_room 是否群聊：1 是 0 否
	RoomConversationID string // room_conversation_id 群 ID
	SendTime           int64  // send_time 发送时间（秒级）
	MsgID              int64  // msg_id 原生消息 ID（账号本地自增序列，跨号不一致，不可作全局唯一键）
	ServerID           int64  // server_id 服务器 ID
}

// Envelope 解析消息回调的通用信封字段（gjson 宽容语义）。
func (m *CallBackMsg) Envelope() MessageEnvelope {
	j := gjson.Parse(m.Json)
	return MessageEnvelope{
		MsgType:            j.Get("msgtype").Int(),
		Sender:             j.Get("sender").Int(),
		Receiver:           j.Get("receiver").Int(),
		IsRoom:             j.Get("is_room").Int(),
		RoomConversationID: j.Get("room_conversation_id").String(),
		SendTime:           j.Get("send_time").Int(),
		MsgID:              j.Get("msg_id").Int(),
		ServerID:           j.Get("server_id").Int(),
	}
}

// ---------------------------------------------------------------------------
// 撤回（ChatMsgRevoke）
// ---------------------------------------------------------------------------

// RevokePayload 撤回回调的定位字段。
//
// 协议约定（docs/callback/02-message.md §2.20）撤回回调不携带原消息 id：
// revoke_time 为原消息发送时间、revoke_msgtype 为原消息类型、msg_id 是
// 撤回动作自身的 id。Sender/Receiver 为字符串视图（gjson String 语义：
// 缺失/为 null 时为空串），与历史业务实现的空串缺失判断保持一致。
type RevokePayload struct {
	IsRoom        int64  // is_room 群聊撤回不支持定位原消息
	RevokeTime    int64  // revoke_time 原消息发送时间（秒，±2s 容差定位）
	RevokeMsgType int64  // revoke_msgtype 原消息类型
	Sender        string // sender（缺失时为空串）
	Receiver      string // receiver（缺失时为空串）
}

// AsRevokeMsg 解析撤回回调定位字段。
func (m *CallBackMsg) AsRevokeMsg() RevokePayload {
	j := gjson.Parse(m.Json)
	return RevokePayload{
		IsRoom:        j.Get("is_room").Int(),
		RevokeTime:    j.Get("revoke_time").Int(),
		RevokeMsgType: j.Get("revoke_msgtype").Int(),
		Sender:        j.Get("sender").String(),
		Receiver:      j.Get("receiver").String(),
	}
}

// ---------------------------------------------------------------------------
// 各消息类型 payload
// ---------------------------------------------------------------------------

// SystemMsgPayload 系统消息（ChatMsgSystem）。
type SystemMsgPayload struct {
	Content    string // msg 系统消息内容
	BanJumpURL string // extraContent.extraMessageLink.jumpUrl 封禁详情跳转链接（封号类系统消息携带）
}

// AsSystemMsg 解析系统消息。
func (m *CallBackMsg) AsSystemMsg() SystemMsgPayload {
	j := gjson.Parse(m.Json)
	return SystemMsgPayload{
		Content:    j.Get("msg").String(),
		BanJumpURL: j.Get("extraContent.extraMessageLink.jumpUrl").String(),
	}
}

// TextMsgPayload 文本消息（ChatMsgText / ChatMsgTextAlt）。
type TextMsgPayload struct {
	Content string // content 文本内容
}

// AsTextMsg 解析文本消息。
func (m *CallBackMsg) AsTextMsg() TextMsgPayload {
	return TextMsgPayload{Content: gjson.Get(m.Json, "content").String()}
}

// GifMsgPayload 表情消息（ChatMsgGif / ChatMsgGifDynamic）。
type GifMsgPayload struct {
	EmotionType string // EmotionType "EMOTION_DYNAMIC" 为动态表情，其他为静态
	URL         string // url 动图地址（ChatMsgGif）
	StaticURL   string // static_url 静态图地址（ChatMsgGifDynamic）
}

// AsGifMsg 解析表情消息。
func (m *CallBackMsg) AsGifMsg() GifMsgPayload {
	j := gjson.Parse(m.Json)
	return GifMsgPayload{
		EmotionType: j.Get("EmotionType").String(),
		URL:         j.Get("url").String(),
		StaticURL:   j.Get("static_url").String(),
	}
}

// MiniProgramMsgPayload 小程序消息（ChatMsgMiniProgram）。
//
// 携带完整素材字段，与协议发送接口 SendAppMsg 参数一一对应，
// 供前端卡片渲染与转发复用。
type MiniProgramMsgPayload struct {
	Title        string // title
	Desc         string // desc
	AppName      string // appName
	WeappIconURL string // weappIconUrl
	Pagepath     string // pagepath
	Username     string // username
	AppID        string // appid
	ThumbFileID  string // thumbFileId 封面 CDN key
	ThumbMD5     string // thumbMD5
	ThumbAESKey  string // thumbAESKey
	FileSize     int64  // size
}

// AsMiniProgramMsg 解析小程序消息。
func (m *CallBackMsg) AsMiniProgramMsg() MiniProgramMsgPayload {
	j := gjson.Parse(m.Json)
	return MiniProgramMsgPayload{
		Title:        j.Get("title").String(),
		Desc:         j.Get("desc").String(),
		AppName:      j.Get("appName").String(),
		WeappIconURL: j.Get("weappIconUrl").String(),
		Pagepath:     j.Get("pagepath").String(),
		Username:     j.Get("username").String(),
		AppID:        j.Get("appid").String(),
		ThumbFileID:  j.Get("thumbFileId").String(),
		ThumbMD5:     j.Get("thumbMD5").String(),
		ThumbAESKey:  j.Get("thumbAESKey").String(),
		FileSize:     j.Get("size").Int(),
	}
}

// LinkMsgPayload 链接消息（ChatMsgLink）。
type LinkMsgPayload struct {
	HeadImg string // headImg
	Title   string // title
	Desc    string // des（协议字段名即 des）
	URL     string // url
}

// AsLinkMsg 解析链接消息。
func (m *CallBackMsg) AsLinkMsg() LinkMsgPayload {
	j := gjson.Parse(m.Json)
	return LinkMsgPayload{
		HeadImg: j.Get("headImg").String(),
		Title:   j.Get("title").String(),
		Desc:    j.Get("des").String(),
		URL:     j.Get("url").String(),
	}
}

// ChannelMsgPayload 视频号消息（ChatMsgVideoChannel）。
type ChannelMsgPayload struct {
	Nickname string // nickname 视频号作者
	Desc     string // desc 内容描述
}

// AsChannelMsg 解析视频号消息。
func (m *CallBackMsg) AsChannelMsg() ChannelMsgPayload {
	j := gjson.Parse(m.Json)
	return ChannelMsgPayload{
		Nickname: j.Get("nickname").String(),
		Desc:     j.Get("desc").String(),
	}
}

// LocationMsgPayload 位置消息（ChatMsgLocation）。
type LocationMsgPayload struct {
	DetailedAddress string // detailed_address 详细地址
}

// AsLocationMsg 解析位置消息。
func (m *CallBackMsg) AsLocationMsg() LocationMsgPayload {
	return LocationMsgPayload{DetailedAddress: gjson.Get(m.Json, "detailed_address").String()}
}

// RedPacketMsgPayload 红包消息（ChatMsgRedPacket）。
type RedPacketMsgPayload struct {
	Desc string // desc 红包文案
}

// AsRedPacketMsg 解析红包消息。
func (m *CallBackMsg) AsRedPacketMsg() RedPacketMsgPayload {
	return RedPacketMsgPayload{Desc: gjson.Get(m.Json, "desc").String()}
}

// CardMsgPayload 名片消息（ChatMsgCard）。
type CardMsgPayload struct {
	Username       string // username 被推荐人 id
	Nickname       string // nickname
	HeadImg        string // headImg
	CorpID         string // corp_id
	EnterpriseName string // enterpriseName
}

// AsCardMsg 解析名片消息。
func (m *CallBackMsg) AsCardMsg() CardMsgPayload {
	j := gjson.Parse(m.Json)
	return CardMsgPayload{
		Username:       j.Get("username").String(),
		Nickname:       j.Get("nickname").String(),
		HeadImg:        j.Get("headImg").String(),
		CorpID:         j.Get("corp_id").String(),
		EnterpriseName: j.Get("enterpriseName").String(),
	}
}

// MediaFilePayload 图片/文件类消息共用的 CDN 下载字段
// （ChatMsgExternalImage / ChatMsgExternalFile）。
type MediaFilePayload struct {
	FileName         string // file_name（可能缺失，业务侧需兜底命名）
	FileID           string // file_id（缺失时无法幂等去重）
	OpenIMCDNAuthKey string // openim_cdn_authkey
	AESKey           string // aeskey（图片/文件回调的加密字段名）
	MD5              string // md5（图片回调携带）
	FileSize         int64  // file_size
}

// AsMediaFile 解析图片/文件消息的 CDN 下载字段。
func (m *CallBackMsg) AsMediaFile() MediaFilePayload {
	j := gjson.Parse(m.Json)
	return MediaFilePayload{
		FileName:         j.Get("file_name").String(),
		FileID:           j.Get("file_id").String(),
		OpenIMCDNAuthKey: j.Get("openim_cdn_authkey").String(),
		AESKey:           j.Get("aeskey").String(),
		MD5:              j.Get("md5").String(),
		FileSize:         j.Get("file_size").Int(),
	}
}

// VideoMsgPayload 外部视频消息（ChatMsgExternalVideo）。
//
// 加密字段双名兼容：视频回调为 aes_key，与图片/文件回调的 aeskey 不一致
// （见 docs/callback/02-message.md 2.10），读错字段会导致协议服务
// DownloadWeChatFile 返回 errcode=500；AESKey 已在 SDK 内完成双名收敛。
type VideoMsgPayload struct {
	FileName         string // file_name（可能缺失）
	FileID           string // file_id
	OpenIMCDNAuthKey string // openim_cdn_authkey
	AESKey           string // aeskey 优先，缺失时回退 aes_key
	FileSize         int64  // file_size
	PreviewImgURL    string // preview_img_url 封面外链（微信 CDN 签名，浏览器直连 400，需转存）
	LDAESKey         string // openim_cdn_ld_aeskey 封面低清加密 key（auth 复用 openim_cdn_authkey）
	LDSize           int64  // openim_cdn_ld_size
	LDWidth          int64  // openim_cdn_ld_width
	LDHeight         int64  // openim_cdn_ld_height
	Duration         int64  // video_duration 时长（秒）
}

// AsVideoMsg 解析外部视频消息（含封面字段）。
func (m *CallBackMsg) AsVideoMsg() VideoMsgPayload {
	j := gjson.Parse(m.Json)
	aesKey := j.Get("aeskey").String()
	if aesKey == "" {
		aesKey = j.Get("aes_key").String()
	}
	return VideoMsgPayload{
		FileName:         j.Get("file_name").String(),
		FileID:           j.Get("file_id").String(),
		OpenIMCDNAuthKey: j.Get("openim_cdn_authkey").String(),
		AESKey:           aesKey,
		FileSize:         j.Get("file_size").Int(),
		PreviewImgURL:    j.Get("preview_img_url").String(),
		LDAESKey:         j.Get("openim_cdn_ld_aeskey").String(),
		LDSize:           j.Get("openim_cdn_ld_size").Int(),
		LDWidth:          j.Get("openim_cdn_ld_width").Int(),
		LDHeight:         j.Get("openim_cdn_ld_height").Int(),
		Duration:         j.Get("video_duration").Int(),
	}
}

// VoiceMsgPayload 语音消息（ChatMsgVoice）。
//
// 加密字段为 aes_key（与视频一致，区别于图片/文件的 aeskey，见 2.14）。
type VoiceMsgPayload struct {
	VoiceID   string // voice_id 语音文件 id（缺失时无法下载）
	VoiceSize int64  // voice_size
	AESKey    string // aes_key
	Duration  int64  // voice_time 时长（秒）
}

// AsVoiceMsg 解析语音消息。
func (m *CallBackMsg) AsVoiceMsg() VoiceMsgPayload {
	j := gjson.Parse(m.Json)
	return VoiceMsgPayload{
		VoiceID:   j.Get("voice_id").String(),
		VoiceSize: j.Get("voice_size").Int(),
		AESKey:    j.Get("aes_key").String(),
		Duration:  j.Get("voice_time").Int(),
	}
}
