// api_label_get_label_list.go GetLabelList 获取标签列表。

package sdk_ipad

import "errors"

// LabelListResp 获取标签列表的返回结果。
type LabelListResp struct {
	Data struct {
		LabelList []LabelItem `json:"labellist"`
		Index     int64       `json:"index"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetLabelList 获取标签列表。
// 端点: POST /wxwork/GetLabelListReq
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - index: 分页索引
//   - syncType: 标签类型（1=企业标签, 2=个人标签）
//
// 返回: *LabelListResp 标签列表
func (s *SdkIPad) GetLabelList(host, uuid string, index, syncType int) (*LabelListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":      uuid,
		"index":     index,
		"sync_type": syncType,
	}
	var resp LabelListResp
	if err := s.httpClient.Post(host+"/wxwork/GetLabelListReq", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
