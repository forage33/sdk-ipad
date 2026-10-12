// api_contact_add_share_user.go AddShareUser 添加同事共享过来的用户好友。

package sdk_ipad

import "errors"

// AddShareUserResp 添加共享联系人的返回结果。
type AddShareUserResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AddShareUser 添加同事共享过来的用户好友。
// 端点: POST /wxwork/AddShareUser
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 同事共享过来的用户 VID
//   - content: 添加好友时的验证消息
//
// 返回: *AddShareUserResp 添加结果
func (s *SdkIPad) AddShareUser(host, uuid string, vid int64, content string) (*AddShareUserResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"vid":     vid,
		"content": content,
	}

	var resp AddShareUserResp
	if err := s.httpClient.Post(host+"/wxwork/AddShareUser", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
