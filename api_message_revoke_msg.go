// api_message_revoke_msg.go RevokeMsg 撤回消息。

package sdk_ipad

import "errors"

// RevokeResp 撤回消息的返回结果。
type RevokeResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// RevokeMsg 撤回消息。
// 端点: POST /wxwork/RevokeMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - msgID: 消息 ID
//   - roomID: 会话 ID（好友或群 ID）
//
// 返回: *RevokeResp 撤回结果
//
// 注意: 与微信用户或群内含微信号的会话只能撤回 2 分钟内的消息
func (s *SdkIPad) RevokeMsg(host, uuid string, msgID, roomID int64) (*RevokeResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"msgid":  msgID,
		"roomid": roomID,
	}
	var resp RevokeResp
	if err := s.httpClient.Post(host+"/wxwork/RevokeMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
