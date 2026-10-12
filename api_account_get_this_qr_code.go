// api_account_get_this_qr_code.go GetThisQRCode 获取当前账号的二维码。

package sdk_ipad

import "errors"

// GetQrCodeResp 获取账号二维码的返回结果。
type GetQrCodeResp struct {
	Data struct {
		QrCodeUrl string `json:"QrCodeUrl"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetThisQRCode 获取当前账号的二维码。
// 端点: POST /wxwork/GetThisQrCode
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *GetQrCodeResp 包含二维码图片地址
func (s *SdkIPad) GetThisQRCode(host, uuid string) (*GetQrCodeResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
	}

	var qrCodeResp GetQrCodeResp
	if err := s.httpClient.Post(host+"/wxwork/GetThisQrCode", reqParams, &qrCodeResp); err != nil {
		return nil, err
	}

	// 使用HTTP客户端发送请求
	return &qrCodeResp, nil
}
