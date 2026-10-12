// api_moments_get_sns_list.go GetSnsList 获取朋友圈列表。

package sdk_ipad

import "errors"

// SnsListResp 获取朋友圈列表的返回结果。
type SnsListResp struct {
	Data    []SnsItem `json:"data"`
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
}

// GetSnsList 获取朋友圈列表。
// 端点: POST /wxwork/GetSnsList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - seq: 分页序列号，首次传 -1，后续传返回的最后一条 sid 进行翻页
//   - limit: 每页数量
//
// 返回: *SnsListResp 朋友圈列表
func (s *SdkIPad) GetSnsList(host, uuid string, seq int64, limit int) (*SnsListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"seq":   seq,
		"limit": limit,
	}
	var resp SnsListResp
	if err := s.httpClient.Post(host+"/wxwork/GetSnsList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
