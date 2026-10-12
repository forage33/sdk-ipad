// api_contact_add_sub_contact_black.go AddSubContactBlack 拉黑或移除黑名单联系人。

package sdk_ipad

import "errors"

// AddSubContactBlack 拉黑或移除黑名单联系人。
// 端点: POST /wxwork/AddSubContactBlack
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人 VID
//   - blackType: 操作类型（1=拉黑, 0=移除黑名单）
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) AddSubContactBlack(host, uuid string, vid int64, blackType int) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid": uuid,
		"vid":  vid,
		"type": blackType,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/AddSubContactBlack", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
