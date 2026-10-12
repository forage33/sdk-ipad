package sdk_ipad

import "testing"

// callback_payload_msg_test.go 消息回调契约层（callback_payload_msg.go）的语义契约测试。
//
// 覆盖的语义铁律（修改契约层前须保持测试通过）：
//   - 宽容解析：字段缺失/类型不匹配返回零值，不报错（对齐历史 gjson 行为）
//   - 数字字段的字符串视图保持 gjson String 语义：数字返回十进制串、
//     缺失返回空串（不可用 FormatInt 替代——null 时 "0" ≠ "" 会改变缺失判断）
//   - 视频加密字段 aeskey/aes_key 双名兼容（aeskey 优先，缺失回退 aes_key）

// TestEnvelope 通用信封 8 字段解析与缺失零值（宽容语义）。
func TestEnvelope(t *testing.T) {
	msg := &CallBackMsg{Json: `{
		"msgtype": 1011,
		"sender": 1688851000000001,
		"receiver": 1688851000000002,
		"is_room": 1,
		"room_conversation_id": "R:10782666091770616",
		"send_time": 1755600000,
		"msg_id": 8848,
		"server_id": 9527
	}`}
	env := msg.Envelope()
	if env.MsgType != 1011 {
		t.Errorf("MsgType = %d, want 1011", env.MsgType)
	}
	if env.Sender != 1688851000000001 {
		t.Errorf("Sender = %d, want 1688851000000001", env.Sender)
	}
	if env.Receiver != 1688851000000002 {
		t.Errorf("Receiver = %d, want 1688851000000002", env.Receiver)
	}
	if env.IsRoom != 1 {
		t.Errorf("IsRoom = %d, want 1", env.IsRoom)
	}
	if env.RoomConversationID != "R:10782666091770616" {
		t.Errorf("RoomConversationID = %q, want R:10782666091770616", env.RoomConversationID)
	}
	if env.SendTime != 1755600000 {
		t.Errorf("SendTime = %d, want 1755600000", env.SendTime)
	}
	if env.MsgID != 8848 {
		t.Errorf("MsgID = %d, want 8848", env.MsgID)
	}
	if env.ServerID != 9527 {
		t.Errorf("ServerID = %d, want 9527", env.ServerID)
	}

	// 全缺失 → 全零值（宽容语义）
	empty := (&CallBackMsg{Json: `{}`}).Envelope()
	if empty != (MessageEnvelope{}) {
		t.Errorf("缺失字段应为零值，got %+v", empty)
	}
}

// TestAsRevokeMsg 撤回回调：revoke_time/revoke_msgtype 定位字段与
// sender/receiver 字符串视图（gjson String 语义）。
func TestAsRevokeMsg(t *testing.T) {
	msg := &CallBackMsg{Json: `{
		"is_room": 0,
		"revoke_time": 1755600123,
		"revoke_msgtype": 0,
		"sender": 1688851000000001,
		"receiver": 1688851000000002
	}`}
	p := msg.AsRevokeMsg()
	if p.IsRoom != 0 || p.RevokeTime != 1755600123 || p.RevokeMsgType != 0 {
		t.Errorf("定位字段解析错误: %+v", p)
	}
	// 数字 → 十进制串
	if p.Sender != "1688851000000001" || p.Receiver != "1688851000000002" {
		t.Errorf("Sender/Receiver 字符串视图错误: %+v", p)
	}

	// 缺失 sender/receiver → 空串（业务侧以空串判定字段缺失，不可变为 "0"）
	missing := (&CallBackMsg{Json: `{"is_room":1,"revoke_time":1755600123}`}).AsRevokeMsg()
	if missing.Sender != "" || missing.Receiver != "" {
		t.Errorf("缺失 sender/receiver 应为空串: %+v", missing)
	}
}

// TestAsVideoMsgAESKeyDualName 视频加密字段双名兼容（aeskey 优先 → aes_key 回退）。
func TestAsVideoMsgAESKeyDualName(t *testing.T) {
	cases := []struct {
		name string
		json string
		want string
	}{
		{"aeskey 优先", `{"aeskey":"key-lower","aes_key":"key-upper"}`, "key-lower"},
		{"缺失 aeskey 回退 aes_key", `{"aes_key":"key-upper"}`, "key-upper"},
		{"均缺失为空串", `{}`, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := (&CallBackMsg{Json: tc.json}).AsVideoMsg().AESKey
			if got != tc.want {
				t.Errorf("AESKey = %q, want %q", got, tc.want)
			}
		})
	}
}

// TestAsVideoMsgFullFields 视频消息完整字段解析（含封面低清字段）。
func TestAsVideoMsgFullFields(t *testing.T) {
	msg := &CallBackMsg{Json: `{
		"file_name": "v.mp4",
		"file_id": "file-1",
		"openim_cdn_authkey": "auth-1",
		"aes_key": "aes-1",
		"file_size": 2048,
		"preview_img_url": "http://wx.example/preview",
		"openim_cdn_ld_aeskey": "ld-aes",
		"openim_cdn_ld_size": 512,
		"openim_cdn_ld_width": 200,
		"openim_cdn_ld_height": 100,
		"video_duration": 15
	}`}
	p := msg.AsVideoMsg()
	if p.FileName != "v.mp4" || p.FileID != "file-1" || p.OpenIMCDNAuthKey != "auth-1" ||
		p.AESKey != "aes-1" || p.FileSize != 2048 {
		t.Errorf("基础字段解析错误: %+v", p)
	}
	if p.PreviewImgURL != "http://wx.example/preview" || p.LDAESKey != "ld-aes" ||
		p.LDSize != 512 || p.LDWidth != 200 || p.LDHeight != 100 || p.Duration != 15 {
		t.Errorf("封面/时长字段解析错误: %+v", p)
	}
}

// TestAsVoiceMsg 语音消息：aes_key 单名（区别于图片/文件的 aeskey）。
func TestAsVoiceMsg(t *testing.T) {
	msg := &CallBackMsg{Json: `{"voice_id":"voice-1","voice_size":4096,"aes_key":"aes-v","voice_time":8}`}
	p := msg.AsVoiceMsg()
	if p.VoiceID != "voice-1" || p.VoiceSize != 4096 || p.AESKey != "aes-v" || p.Duration != 8 {
		t.Errorf("语音字段解析错误: %+v", p)
	}
	// 缺失全部 → 零值
	if zero := (&CallBackMsg{Json: `{}`}).AsVoiceMsg(); zero != (VoiceMsgPayload{}) {
		t.Errorf("缺失字段应为零值: %+v", zero)
	}
}

// TestAsMediaFile 图片/文件消息：aeskey 单名 + md5。
func TestAsMediaFile(t *testing.T) {
	msg := &CallBackMsg{Json: `{
		"file_name": "doc.pdf",
		"file_id": "file-2",
		"openim_cdn_authkey": "auth-2",
		"aeskey": "aes-2",
		"md5": "d41d8cd98f00b204e9800998ecf8427e",
		"file_size": 1024
	}`}
	p := msg.AsMediaFile()
	if p.FileName != "doc.pdf" || p.FileID != "file-2" || p.OpenIMCDNAuthKey != "auth-2" ||
		p.AESKey != "aes-2" || p.MD5 != "d41d8cd98f00b204e9800998ecf8427e" || p.FileSize != 1024 {
		t.Errorf("图片/文件字段解析错误: %+v", p)
	}
}

// TestAsSystemMsg 系统消息：内容与封禁跳转链接（嵌套路径 extraContent.extraMessageLink.jumpUrl）。
func TestAsSystemMsg(t *testing.T) {
	msg := &CallBackMsg{Json: `{
		"msg": "你的账号已被限制",
		"extraContent": {"extraMessageLink": {"jumpUrl": "https://work.weixin.qq.com/ban"}}
	}`}
	p := msg.AsSystemMsg()
	if p.Content != "你的账号已被限制" {
		t.Errorf("Content = %q", p.Content)
	}
	if p.BanJumpURL != "https://work.weixin.qq.com/ban" {
		t.Errorf("BanJumpURL = %q", p.BanJumpURL)
	}

	// 普通系统消息无跳转链接字段 → 空串
	plain := (&CallBackMsg{Json: `{"msg":"普通提示"}`}).AsSystemMsg()
	if plain.BanJumpURL != "" {
		t.Errorf("无跳转链接时应为空串, got %q", plain.BanJumpURL)
	}
}

// TestAsCardMsg 名片消息：被推荐人昵称/头像/企业名（服务端名片落库依赖字段）。
func TestAsCardMsg(t *testing.T) {
	msg := &CallBackMsg{Json: `{
		"username": "1688851000000009",
		"nickname": "被推荐人",
		"headImg": "http://wx.example/head.png",
		"corp_id": "1970324836974591",
		"enterpriseName": "某某公司"
	}`}
	p := msg.AsCardMsg()
	if p.Username != "1688851000000009" || p.Nickname != "被推荐人" ||
		p.HeadImg != "http://wx.example/head.png" || p.CorpID != "1970324836974591" ||
		p.EnterpriseName != "某某公司" {
		t.Errorf("名片字段解析错误: %+v", p)
	}
}

// TestAsLinkMsg 链接消息：描述字段协议名即 des（非 desc）。
func TestAsLinkMsg(t *testing.T) {
	msg := &CallBackMsg{Json: `{
		"headImg": "http://wx.example/link.png",
		"title": "标题",
		"des": "描述文本",
		"url": "https://example.com/a"
	}`}
	p := msg.AsLinkMsg()
	if p.HeadImg != "http://wx.example/link.png" || p.Title != "标题" ||
		p.Desc != "描述文本" || p.URL != "https://example.com/a" {
		t.Errorf("链接字段解析错误: %+v", p)
	}
}

// TestAsMiniProgramMsg 小程序消息：素材字段与 SendAppMsg 参数一一对应。
func TestAsMiniProgramMsg(t *testing.T) {
	msg := &CallBackMsg{Json: `{
		"title": "小程序标题",
		"desc": "小程序描述",
		"appName": "某某应用",
		"weappIconUrl": "http://wx.example/icon.png",
		"pagepath": "pages/index?x=1",
		"username": "gh_xxx@app",
		"appid": "wx1234567890",
		"thumbFileId": "thumb-1",
		"thumbMD5": "md5-1",
		"thumbAESKey": "aes-thumb",
		"size": 65536
	}`}
	p := msg.AsMiniProgramMsg()
	if p.Title != "小程序标题" || p.Desc != "小程序描述" || p.AppName != "某某应用" ||
		p.WeappIconURL != "http://wx.example/icon.png" || p.Pagepath != "pages/index?x=1" ||
		p.Username != "gh_xxx@app" || p.AppID != "wx1234567890" {
		t.Errorf("小程序基础字段解析错误: %+v", p)
	}
	if p.ThumbFileID != "thumb-1" || p.ThumbMD5 != "md5-1" ||
		p.ThumbAESKey != "aes-thumb" || p.FileSize != 65536 {
		t.Errorf("小程序封面字段解析错误: %+v", p)
	}
}

// TestAsGifMsg 表情消息：EmotionType 区分静态/动态。
func TestAsGifMsg(t *testing.T) {
	msg := &CallBackMsg{Json: `{"EmotionType":"EMOTION_DYNAMIC","static_url":"http://wx.example/s.png"}`}
	p := msg.AsGifMsg()
	if p.EmotionType != "EMOTION_DYNAMIC" || p.StaticURL != "http://wx.example/s.png" {
		t.Errorf("表情字段解析错误: %+v", p)
	}
}

// TestMsgAccessorsTolerateMissing 宽容语义总契约：
// 空 JSON 下所有消息域访问器均不 panic 且返回零值。
func TestMsgAccessorsTolerateMissing(t *testing.T) {
	msg := &CallBackMsg{Json: `{}`}
	if got := msg.AsTextMsg(); got.Content != "" {
		t.Errorf("AsTextMsg = %+v", got)
	}
	if got := msg.AsSystemMsg(); got.Content != "" || got.BanJumpURL != "" {
		t.Errorf("AsSystemMsg = %+v", got)
	}
	if got := msg.AsGifMsg(); got != (GifMsgPayload{}) {
		t.Errorf("AsGifMsg = %+v", got)
	}
	if got := msg.AsMiniProgramMsg(); got != (MiniProgramMsgPayload{}) {
		t.Errorf("AsMiniProgramMsg = %+v", got)
	}
	if got := msg.AsLinkMsg(); got != (LinkMsgPayload{}) {
		t.Errorf("AsLinkMsg = %+v", got)
	}
	if got := msg.AsChannelMsg(); got != (ChannelMsgPayload{}) {
		t.Errorf("AsChannelMsg = %+v", got)
	}
	if got := msg.AsLocationMsg(); got != (LocationMsgPayload{}) {
		t.Errorf("AsLocationMsg = %+v", got)
	}
	if got := msg.AsRedPacketMsg(); got != (RedPacketMsgPayload{}) {
		t.Errorf("AsRedPacketMsg = %+v", got)
	}
	if got := msg.AsCardMsg(); got != (CardMsgPayload{}) {
		t.Errorf("AsCardMsg = %+v", got)
	}
	if got := msg.AsMediaFile(); got != (MediaFilePayload{}) {
		t.Errorf("AsMediaFile = %+v", got)
	}
	if got := msg.AsVideoMsg(); got != (VideoMsgPayload{}) {
		t.Errorf("AsVideoMsg = %+v", got)
	}
	if got := msg.AsVoiceMsg(); got != (VoiceMsgPayload{}) {
		t.Errorf("AsVoiceMsg = %+v", got)
	}
	if got := msg.AsRevokeMsg(); got != (RevokePayload{}) {
		t.Errorf("AsRevokeMsg = %+v", got)
	}
}

// TestAsTextMsgContent 文本消息内容与非 JSON 容错（gjson 对非法 JSON 返回零值不报错）。
func TestAsTextMsgContent(t *testing.T) {
	if got := (&CallBackMsg{Json: `{"content":"你好"}`}).AsTextMsg(); got.Content != "你好" {
		t.Errorf("Content = %q, want 你好", got.Content)
	}
	// 非法 JSON：宽容语义返回零值（不 panic、不报错）
	if got := (&CallBackMsg{Json: `{bad`}).AsTextMsg(); got.Content != "" {
		t.Errorf("非法 JSON 应返回零值, got %q", got.Content)
	}
}
