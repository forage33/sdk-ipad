// api_contact_delete_contact.go DeleteContact 删除联系人。

package sdk_ipad

import "errors"

// DeleteContact 删除联系人。
// 端点: POST /wxwork/DelContact
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人 VID
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) DeleteContact(host, uuid string, vid int64) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid": uuid,
		"vid":  vid,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/DelContact", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
