// api_wx_init_get_run_client.go GetRunClient 获取所有运行中的实例列表。

package sdk_ipad

// ClientResponse 获取运行中实例列表的返回结果。
type ClientResponse struct {
	Data    []map[string]UserData `json:"data"`
	ErrCode int                   `json:"errcode"`
	ErrMsg  string                `json:"errmsg"`
}

// GetRunClient 获取所有运行中的实例列表。
// 端点: POST /wxwork/GetRunClient
//
// 参数:
//   - host: 协议服务地址
//
// 返回: *ClientResponse 运行中实例列表
func (s *SdkIPad) GetRunClient(host string) (*ClientResponse, error) {
	var clientResp ClientResponse
	if err := s.httpClient.Post(host+"/wxwork/GetRunClient", nil, &clientResp); err != nil {
		return nil, err
	}
	return &clientResp, nil
}
