// api_label_edit_label.go EditLabel 修改标签名称。

package sdk_ipad

import "errors"

// EditLabelResp 修改标签的返回结果。
type EditLabelResp struct {
	Data    LabelItem `json:"data"`
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
}

// EditLabel 修改标签名称。
// 端点: POST /wxwork/EditLabelReq
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - labelName: 新标签名称
//   - labelID: 标签 ID
//   - groupID: 标签分组 ID
//
// 返回: *EditLabelResp 修改后的标签详情
func (s *SdkIPad) EditLabel(host, uuid, labelName string, labelID, groupID int64) (*EditLabelResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":      uuid,
		"labelName": labelName,
		"labelid":   labelID,
		"groupid":   groupID,
	}
	var resp EditLabelResp
	if err := s.httpClient.Post(host+"/wxwork/EditLabelReq", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
