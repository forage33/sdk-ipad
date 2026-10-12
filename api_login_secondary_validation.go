// api_login_secondary_validation.go SecondaryValidation 获取二次验证二维码（新设备登录时需要）。

package sdk_ipad

import "errors"

// SecondaryValidation 二次验证的返回结果，包含新的二维码信息。
type SecondaryValidation struct {
	Data    SecondaryValidationData `json:"data"`
	Errcode int                     `json:"errcode"`
	Errmsg  string                  `json:"errmsg"`
}

// SecondaryValidationData 二次验证二维码数据。
type SecondaryValidationData struct {
	Qrcode string `json:"qrcode"`
	Ttl    int    `json:"Ttl"`
	Key    string `json:"Key"`
}

// SecondaryValidation 获取二次验证二维码（新设备登录时需要）。
// 端点: POST /wxwork/SecondaryValidation
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *SecondaryValidation 包含二次验证二维码
func (s *SdkIPad) SecondaryValidation(host, uuid string) (*SecondaryValidation, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
	}
	var second SecondaryValidation
	// 使用HTTP客户端发送请求
	if err := s.httpClient.Post(host+"/wxwork/SecondaryValidation", reqParams, &second); err != nil {
		return nil, err
	}
	return &second, nil
}
