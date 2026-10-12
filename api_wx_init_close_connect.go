// api_wx_init_close_connect.go CloseConnect 关闭指定实例的连接。

package sdk_ipad

import "errors"

// CloseConnResp 关闭连接的返回结果。
type CloseConnResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// CloseConnect 关闭指定实例的连接。
// 端点: POST /wxwork/CloseConnent
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *CloseConnResp 关闭结果
func (s *SdkIPad) CloseConnect(host, uuid string) (*CloseConnResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
	}

	var closeResp CloseConnResp
	if err := s.httpClient.Post(host+"/wxwork/CloseConnent", reqParams, &closeResp); err != nil {
		return nil, err
	}
	return &closeResp, nil
}
