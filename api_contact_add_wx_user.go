// api_contact_add_wx_user.go AddWxUser 直接添加好友（仅对曾被删除的好友有效）。

package sdk_ipad

import "errors"

// AddWxUser 直接添加好友（仅对曾被删除的好友有效）。
// 端点: POST /wxwork/addWxUser
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 目标用户 VID
//   - content: 好友申请验证消息
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) AddWxUser(host, uuid string, vid int64, content string) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"vid":     vid,
		"content": content,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/addWxUser", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
