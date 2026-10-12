// api_message_send_text_at_msg.go SendTextAtMsg 发送群 @消息。

package sdk_ipad

import "errors"

// SendTextAtResp 发送群 @消息的返回结果。
type SendTextAtResp struct {
	Data struct {
		Receiver int64 `json:"receiver"`
		Sender   int64 `json:"sender"`
		AtList   []struct {
			UserId   int64  `json:"user_id"`
			NikeName string `json:"nikeName"`
		} `json:"at_list"`
		SenderName         string `json:"sender_name"`
		AppInfo            string `json:"app_info"`
		RoomConversationId string `json:"room_conversation_id"`
		IsRoom             int    `json:"is_room"`
		Sendtime           int    `json:"sendtime"`
		MsgId              int    `json:"msg_id"`
		ServerId           int    `json:"server_id"`
		Msgtype            int    `json:"msgtype"`
		Content            string `json:"content"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SendTextAtMsg 发送群 @消息。
// 端点: POST /wxwork/SendTextAtMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 群 ID
//   - atIds: 要 @的用户 ID 列表
//   - content: 消息文本内容
//
// 返回: *SendTextAtResp 发送结果
//
// 注意: 此接口 isRoom 固定为 true，仅用于群聊场景
func (s *SdkIPad) SendTextAtMsg(host, uuid string, sendUserId int64, atIds []int64, content string) (*SendTextAtResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"send_userid": sendUserId,
		"atids":       atIds,
		"content":     content,
		"isRoom":      true,
	}
	var resp SendTextAtResp
	if err := s.httpClient.Post(host+"/wxwork/SendTextAtMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
