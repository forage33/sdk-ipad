// api_contact_oper_customer.go OperCustomer 好友操作（共享给其他企微号）。

package sdk_ipad

import "errors"

// OperCustomerResp 好友操作（共享/转发）的返回结果。
type OperCustomerResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// OperCustomer 好友操作（共享给其他企微号）。
// 端点: POST /wxwork/OperCustomer
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - thisVID: 当前账号 VID
//   - recommendVID: 要共享的好友 user_id
//   - sendVID: 要共享给哪个同事的 VID
//
// 返回: *OperCustomerResp 操作结果
func (s *SdkIPad) OperCustomer(host, uuid, thisVID, recommendVID, sendVID string) (*OperCustomerResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":         uuid,
		"thisvid":      thisVID,
		"recommendvid": recommendVID,
		"sendvid":      sendVID,
	}

	var resp OperCustomerResp
	if err := s.httpClient.Post(host+"/wxwork/OperCustomer", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
