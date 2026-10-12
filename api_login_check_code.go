// api_login_check_code.go CheckCode 提交扫码后的验证码进行校验。

package sdk_ipad

import "errors"

// InputCodeResp 输入验证码的返回结果。
type InputCodeResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// CheckCode 提交扫码后的验证码进行校验。
// 端点: POST /wxwork/CheckCode
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - qrCodeKey: 二维码校验 Key（由 GetQrCode 返回）
//   - code: 用户输入的验证码
//
// 返回: *InputCodeResp 校验结果
func (s *SdkIPad) CheckCode(host, uuid, qrCodeKey, code string) (*InputCodeResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid":      uuid,
		"qrcodeKey": qrCodeKey,
		"code":      code,
	}
	var inputCodeResp InputCodeResp
	if err := s.httpClient.Post(host+"/wxwork/CheckCode", reqParams, &inputCodeResp); err != nil {
		return nil, err
	}
	return &inputCodeResp, nil
}
