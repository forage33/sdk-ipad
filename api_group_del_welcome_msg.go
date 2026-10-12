// api_group_del_welcome_msg.go DelWelcomeMsg 取消群欢迎语。

package sdk_ipad

import "errors"

// DelWelcomeMsg 取消群欢迎语。
// 端点: POST /wxwork/DelWelcomeMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) DelWelcomeMsg(host, uuid string, roomID int64) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/DelWelcomeMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
