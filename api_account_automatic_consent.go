// api_account_automatic_consent.go AutomaticConsent 设置是否自动同意好友申请。

package sdk_ipad

import "errors"

// AutomaticConsent 设置是否自动同意好友申请。
// 端点: POST /wxwork/AutomaticConsent
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - state: 0=关闭验证自动同意好友, 1=需要验证同意（手动点击同意）
//
// 返回: *CommonAccountResp 操作结果
func (s *SdkIPad) AutomaticConsent(host, uuid string, state int) (*CommonAccountResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"state": state,
	}

	var resp CommonAccountResp
	if err := s.httpClient.Post(host+"/wxwork/AutomaticConsent", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
