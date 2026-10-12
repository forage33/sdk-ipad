// api_moments_get_sns_info.go GetSnsInfo 获取朋友圈详情。

package sdk_ipad

import "errors"

// SnsInfoResp 获取朋友圈详情的返回结果。
type SnsInfoResp struct {
	Data    SnsItem `json:"data"`
	Errcode int     `json:"errcode"`
	Errmsg  string  `json:"errmsg"`
}

// GetSnsInfo 获取朋友圈详情。
// 端点: POST /wxwork/GetSnsInfo
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sid: 朋友圈 SID
//
// 返回: *SnsInfoResp 朋友圈详情
func (s *SdkIPad) GetSnsInfo(host, uuid string, sid int64) (*SnsInfoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid": uuid,
		"sid":  sid,
	}
	var resp SnsInfoResp
	if err := s.httpClient.Post(host+"/wxwork/GetSnsInfo", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
