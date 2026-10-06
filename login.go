// login.go 登录相关接口。
// 提供二维码获取、验证码校验、自动登录、退出登录和二次验证等功能。

package sdk_ipad

import (
	"errors"
)

// ---------------------------------------------------------------------------
// 返回数据结构
// ---------------------------------------------------------------------------

// GetQRCodeResp 获取登录二维码的返回结果。
type GetQRCodeResp struct {
	Data    QRCodeData `json:"data"`
	Errcode int        `json:"errcode"`
	Errmsg  string     `json:"errmsg"`
}

// QRCodeData 二维码核心数据，包含二维码内容、有效期和校验 Key。
type QRCodeData struct {
	Qrcode     string `json:"qrcode"`
	QrcodeData string `json:"qrcode_data"`
	Ttl        int    `json:"Ttl"`
	Key        string `json:"Key"`
}

// InputCodeResp 输入验证码的返回结果。
type InputCodeResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AutoLoginResp 自动登录的返回结果。
type AutoLoginResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// LoginOutResp 退出登录的返回结果。
type LoginOutResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

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

// ---------------------------------------------------------------------------
// 调用接口
// ---------------------------------------------------------------------------

// GetQrCode 获取登录二维码。
// 端点: POST /wxwork/getQrCode
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *GetQRCodeResp 包含二维码内容和有效期
func (s *SdkIPad) GetQrCode(host, uuid string) (*GetQRCodeResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
	}

	var getQRCodeResp GetQRCodeResp
	if err := s.httpClient.Post(host+"/wxwork/getQrCode", reqParams, &getQRCodeResp); err != nil {
		return nil, err
	}
	return &getQRCodeResp, nil
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
