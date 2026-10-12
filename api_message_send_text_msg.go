// api_message_send_text_msg.go SendTextMsg 发送文本消息。

package sdk_ipad

import "errors"

// SendTextResp 发送文本消息的返回结果。
type SendTextResp struct {
	Data struct {
		Receiver   int64         `json:"receiver"`
		Sender     int64         `json:"sender"`
		AtList     []interface{} `json:"at_list"`
		SenderName string        `json:"sender_name"`
		AppInfo    string        `json:"app_info"`
		IsRoom     int           `json:"is_room"`
		Sendtime   int           `json:"sendtime"`
		MsgId      int           `json:"msg_id"`
		ServerId   int           `json:"server_id"`
		Msgtype    int           `json:"msgtype"`
		Content    string        `json:"content"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SendTextMsg 发送文本消息。
// 端点: POST /wxwork/SendTextMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserID: 接收者 ID（好友或群 ID）
//   - content: 文本内容
//   - isRoom: 是否为群消息
//
// 返回: *SendTextResp 包含消息 ID 等发送详情
func (s *SdkIPad) SendTextMsg(host, uuid string, sendUserID int64, content string, isRoom bool) (*SendTextResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"send_userid": sendUserID,
		"isRoom":      isRoom,
		"content":     content,
	}
	var resp SendTextResp
	if err := s.httpClient.Post(host+"/wxwork/SendTextMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
