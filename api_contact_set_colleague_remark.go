// api_contact_set_colleague_remark.go SetColleagueRemark 修改内部联系人备注。

package sdk_ipad

import "errors"

// SetColleagueRemark 修改内部联系人备注。
// 端点: POST /wxwork/SetColleagueRemark
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人 VID
//   - remark: 备注名
//   - des: 描述
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) SetColleagueRemark(host, uuid string, vid int64, remark, des string) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"vid":    vid,
		"remark": remark,
		"des":    des,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/SetColleagueRemark", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
