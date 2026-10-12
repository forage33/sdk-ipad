// api_message_mark_as_read.go MarkAsRead 标记消息为已读。

package sdk_ipad

import "errors"

// MarkAsReadResp 消息已读的返回结果。
type MarkAsReadResp struct {
	Data struct {
		Receiver           int    `json:"receiver"`
		Sender             int64  `json:"sender"`
		SenderName         string `json:"sender_name"`
		AppInfo            string `json:"app_info"`
		RoomConversationId string `json:"room_conversation_id"`
		IsRoom             int    `json:"is_room"`
		Sendtime           int    `json:"sendtime"`
		MsgId              int    `json:"msg_id"`
		ServerId           int    `json:"server_id"`
		Msgtype            int    `json:"msgtype"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// MarkAsRead 标记消息为已读。
// 端点: POST /wxwork/MarkAsRead
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 好友或群 ID
//   - isRoom: 是否为群消息
//
// 返回: *MarkAsReadResp 操作结果
func (s *SdkIPad) MarkAsRead(host, uuid string, sendUserId int64, isRoom bool) (*MarkAsReadResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"send_userid": sendUserId, // 好友或者群id
		"isRoom":      isRoom,
	}
	var resp MarkAsReadResp
	if err := s.httpClient.Post(host+"/wxwork/MarkAsRead", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
