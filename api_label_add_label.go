// api_label_add_label.go AddLabel 添加标签。

package sdk_ipad

import "errors"

// AddLabelResp 添加标签的返回结果。
type AddLabelResp struct {
	Data    LabelItem `json:"data"`
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
}

// AddLabel 添加标签。
// 端点: POST /wxwork/AddLabelReq
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - labelName: 标签名称
//   - dataType: 数据类型
//   - labelType: 标签类型
//   - groupID: 标签分组 ID
//
// 返回: *AddLabelResp 包含新标签详情
func (s *SdkIPad) AddLabel(host, uuid, labelName string, dataType, labelType int, groupID int64) (*AddLabelResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":       uuid,
		"labelName":  labelName,
		"data_type":  dataType,
		"label_type": labelType,
		"groupid":    groupID,
	}
	var resp AddLabelResp
	if err := s.httpClient.Post(host+"/wxwork/AddLabelReq", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
