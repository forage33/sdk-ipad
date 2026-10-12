// api_group_set_welcome_msg.go SetWelcomeMsg 为群设置欢迎语。

package sdk_ipad

import "errors"

// SetWelcomeMsg 为群设置欢迎语。
// 端点: POST /wxwork/SetWelcomeMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - indexId: 欢迎语索引 ID
//   - msgType: 消息类型
//   - roomID: 群 ID
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) SetWelcomeMsg(host, uuid string, indexId, msgType int, roomID int64) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"indexid": indexId,
		"type":    msgType,
		"roomid":  roomID,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/SetWelcomeMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
