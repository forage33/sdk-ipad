// api_label_del_label.go DelLabel 删除标签。

package sdk_ipad

import "errors"

// DelLabelResp 删除标签的返回结果。
type DelLabelResp struct {
	Data    LabelItem `json:"data"`
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
}

// DelLabel 删除标签。
// 端点: POST /wxwork/DelLabelReq
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - labelID: 标签 ID
//   - groupID: 标签分组 ID
//
// 返回: *DelLabelResp 删除的标签详情
func (s *SdkIPad) DelLabel(host, uuid string, labelID, groupID int64) (*DelLabelResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"labelid": labelID,
		"groupid": groupID,
	}
	var resp DelLabelResp
	if err := s.httpClient.Post(host+"/wxwork/DelLabelReq", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
