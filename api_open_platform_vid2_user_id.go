// api_open_platform_vid2_user_id.go Vid2UserId 企业用户 id 转开放平台 id。

package sdk_ipad

import "errors"

// Vid2UserIdResp 企业用户 id 转开放平台 id 的返回结果。
type Vid2UserIdResp struct {
	Data    []VidOpenIdPair `json:"data"`
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
}

// Vid2UserId 企业用户 id 转开放平台 id（external_userid）。
// 端点: POST /wxwork/Vid2UserId
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - scorpID: 第三方应用企业提供商的企业 id（scorpid），不传默认转当前账号企业的开放平台 id
//   - corpID: 第三方应用企业提供商的内部企业 id，不传（0）默认转当前账号企业的开放平台 id
//   - userIDs: 要转换的企业用户 id（VID）列表
//
// 返回: *Vid2UserIdResp 每项含 user_id 与对应 openid
func (s *SdkIPad) Vid2UserId(host, uuid, scorpID string, corpID int64, userIDs []int64) (*Vid2UserIdResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"user_id": userIDs,
	}
	if scorpID != "" {
		reqParams["scorpid"] = scorpID
	}
	if corpID != 0 {
		reqParams["corpid"] = corpID
	}
	var resp Vid2UserIdResp
	if err := s.httpClient.Post(host+"/wxwork/Vid2UserId", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
