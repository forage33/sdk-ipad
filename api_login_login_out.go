// api_login_login_out.go LoginOut 退出登录。

package sdk_ipad

import "errors"

// LoginOutResp 退出登录的返回结果。
type LoginOutResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// LoginOut 退出登录。
// 端点: POST /wxwork/LoginOut
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *LoginOutResp 退出结果
func (s *SdkIPad) LoginOut(host, uuid string) (*LoginOutResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
	}
	var loginOutResp LoginOutResp
	// 使用HTTP客户端发送请求
	if err := s.httpClient.Post(host+"/wxwork/LoginOut", reqParams, &loginOutResp); err != nil {
		return nil, err
	}
	return &loginOutResp, nil
}
