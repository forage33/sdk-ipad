// api_broadcast_confirm_group_send.go ConfirmGroupSend 确认群发消息（控制台下发后需客户端确认执行）。

package sdk_ipad

import "errors"

// ConfirmGroupSendResp 确认群发消息的返回结果。
type ConfirmGroupSendResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// ConfirmGroupSend 确认群发消息（控制台下发后需客户端确认执行）。
// 端点: POST /wxwork/oksendgroup
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vids: 接收者 VID 列表
//   - msgID: 群发消息 ID
//   - isRoom: 是否群发到群（0=个人, 1=群）
//
// 返回: *ConfirmGroupSendResp 操作结果
func (s *SdkIPad) ConfirmGroupSend(host, uuid string, vids []int64, msgID int64, isRoom int) (*ConfirmGroupSendResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"vids":   vids,
		"msgid":  msgID,
		"isroom": isRoom,
	}
	var resp ConfirmGroupSendResp
	if err := s.httpClient.Post(host+"/wxwork/oksendgroup", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
