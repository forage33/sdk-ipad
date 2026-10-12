// api_message_send_quote_msg.go SendQuoteMsg 发送引用消息。

package sdk_ipad

import "errors"

// QuoteContentItem 引用消息发送内容项（content 数组元素）。
//
// 固定两项：第一项为引用前缀行（如 `"发送人：\n原内容"\n------\n` 格式的展示文本），
// 第二项为本次回复的正文；两项的 msgtype 均为 0（文本）。
type QuoteContentItem struct {
	MsgType int    `json:"msgtype"` // 消息类型，0=文本
	Content string `json:"content"` // 文本内容
}

// QuoteMsg 引用的消息对象（quoteMsg）。
//
// 被引用的消息内容必须完整携带（否则客户端展示不全）：
// appinfo 与 content_type 取自原消息回调/查询结果，content 为原消息的
// 类型化内容对象（msgtype + 各消息类型字段，形态见协议文档 13.20）。
type QuoteMsg struct {
	AppInfo     string                 `json:"appinfo"`                // 引用的那条消息的 appinfo
	ContentType int                    `json:"content_type"`           // 引用的那条消息的类型
	Sender      int64                  `json:"sender"`                 // 引用的那条消息发送人 id
	QuoteRoomID int64                  `json:"quote_roomid,omitempty"` // 群聊引用时填群 id；单聊忽略
	Content     map[string]interface{} `json:"content"`                // 引用的消息内容（msgtype=0 文本/1 图片/2 文件/3 视频/4 语音/5 链接/6 名片/7 小程序）
}

// SendQuoteMsgResp 发送引用消息的返回结果。
type SendQuoteMsgResp struct {
	Data    SendQuoteMsgData `json:"data"`
	Errcode int              `json:"errcode"`
	Errmsg  string           `json:"errmsg"`
}

// SendQuoteMsgData 发送引用消息返回的数据部分。
type SendQuoteMsgData struct {
	InnerkfVid    int64             `json:"innerkf_vid"`
	Flag          int64             `json:"flag"`
	Receiver      int64             `json:"receiver"`
	DevInfo       int64             `json:"devInfo"`
	SenderName    string            `json:"sender_name"`
	IsRoom        int               `json:"is_room"`
	QuoteContent  SendQuoteMsgQuote `json:"quote_content"`
	ServerId      int64             `json:"server_id"`
	KfID          int64             `json:"kf_id"`
	Content       string            `json:"content"`
	Issync        bool              `json:"issync"`
	SendTime      int64             `json:"send_time"`
	Sender        int64             `json:"sender"`
	ReferID       int64             `json:"referid"`
	AppInfo       string            `json:"app_info"`
	ReadUinsCount int64             `json:"readuinscount"`
	MsgId         int64             `json:"msg_id"`
	Sendtime      int64             `json:"sendtime"`
	MsgType       int               `json:"msgtype"`
	AtList        []interface{}     `json:"atList"`
}

// SendQuoteMsgQuote 发送引用消息返回的引用内容详情（quote_content）。
type SendQuoteMsgQuote struct {
	Uin             int64         `json:"uin"`
	Time            int64         `json:"time"`
	IsReceipt       bool          `json:"is_receipt"`
	ContentType     int           `json:"contenttype"`
	Name            string        `json:"name"`
	AvatarURL       string        `json:"avatar_url"`
	ConvType        int           `json:"conv_type"`
	ConversationID  int64         `json:"conversation_id"`
	SenderName      string        `json:"sender_name"`
	SenderAvatarURL string        `json:"sender_avatar_url"`
	CorpID          int64         `json:"corp_id"`
	CorpName        string        `json:"corp_name"`
	RefAppInfo      string        `json:"ref_appinfo"`
	RefContent      string        `json:"ref_content"`
	QuoteMsgTips    string        `json:"quoteMsgTips"`
	RefAtList       []interface{} `json:"ref_at_list"`
}

// SendQuoteMsg 发送引用消息（回复指定消息并携带引用关系）。
// 端点: POST /wxwork/sendQuoteMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - kfID: 客服 id（用于回复客服消息，无则传 0）
//   - sendUserID: 接收者 id（好友或群 id）
//   - isRoom: 是否为群消息
//   - content: 发送内容数组（固定两项：引用前缀行 + 回复正文，见 QuoteContentItem）
//   - quoteMsg: 引用的消息对象（必须完整携带原消息内容，见 QuoteMsg）
//
// 返回: *SendQuoteMsgResp 含消息 id 等发送详情
func (s *SdkIPad) SendQuoteMsg(host, uuid string, kfID, sendUserID int64, isRoom bool, content []QuoteContentItem, quoteMsg *QuoteMsg) (*SendQuoteMsgResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"kf_id":       kfID,
		"send_userid": sendUserID,
		"isRoom":      isRoom,
		"content":     content,
		"quoteMsg":    quoteMsg,
	}
	var resp SendQuoteMsgResp
	if err := s.httpClient.Post(host+"/wxwork/sendQuoteMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
