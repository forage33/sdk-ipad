// api_login_get_qr_code.go GetQrCode 获取登录二维码。

package sdk_ipad

import "errors"

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
