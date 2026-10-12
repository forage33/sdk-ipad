// api_account_external_business_card.go ExternalBusinessCard 获取当前账号的对外名片二维码。

package sdk_ipad

import "errors"

// ExternalBusinessCardResp 获取对外名片二维码的返回结果。
type ExternalBusinessCardResp struct {
	Data struct {
		QrImage string `json:"qr_image"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ExternalBusinessCard 获取当前账号的对外名片二维码。
// 端点: POST /wxwork/ExternalBusinessCard
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *ExternalBusinessCardResp 包含名片二维码图片
func (s *SdkIPad) ExternalBusinessCard(host, uuid string) (*ExternalBusinessCardResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]string{
		"uuid": uuid,
	}

	var resp ExternalBusinessCardResp
	if err := s.httpClient.Post(host+"/wxwork/ExternalBusinessCard", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
