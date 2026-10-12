// api_contact_agree_user.go AgreeUser 同意好友申请。

package sdk_ipad

import "errors"

// AgreeUser 同意好友申请。
// 端点: POST /wxwork/AgreeUser
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 申请人 VID
//   - corpid: 申请人企业 ID
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) AgreeUser(host, uuid string, vid int64, corpid int64) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"vid":    vid,
		"corpid": corpid,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/AgreeUser", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
