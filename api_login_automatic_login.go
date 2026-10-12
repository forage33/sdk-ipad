// api_login_automatic_login.go AutomaticLogin 使用已有登录态自动登录（免扫码）。

package sdk_ipad

import "errors"

// AutoLoginResp 自动登录的返回结果。
type AutoLoginResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AutomaticLogin 使用已有登录态自动登录（免扫码）。
// 端点: POST /wxwork/automaticLogin
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *AutoLoginResp 登录结果
func (s *SdkIPad) AutomaticLogin(host, uuid string) (*AutoLoginResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
	}

	var autoLoginResp AutoLoginResp
	if err := s.httpClient.Post(host+"/wxwork/automaticLogin", reqParams, &autoLoginResp); err != nil {
		return nil, err
	}

	return &autoLoginResp, nil
}
