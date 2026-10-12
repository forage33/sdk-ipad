// api_label_user_add_labels.go UserAddLabels 为一个用户批量关联标签。

package sdk_ipad

import "errors"

// UserAddLabelsResp 用户关联标签的返回结果。
type UserAddLabelsResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// UserAddLabels 为一个用户批量关联标签。
// 端点: POST /wxwork/UserAddLabelsReq
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - userID: 用户 ID
//   - labelIDList: 要关联的标签 ID 列表
//
// 返回: *UserAddLabelsResp 操作结果
func (s *SdkIPad) UserAddLabels(host, uuid string, userID int64, labelIDList []int64) (*UserAddLabelsResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":         uuid,
		"userid":       userID,
		"labelid_list": labelIDList,
	}
	var resp UserAddLabelsResp
	if err := s.httpClient.Post(host+"/wxwork/UserAddLabelsReq", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
