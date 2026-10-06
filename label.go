// label.go 标签操作相关接口。
// 提供标签列表获取、添加、修改、删除标签以及标签与用户的关联操作。

package sdk_ipad

import "errors"

// ---------------------------------------------------------------------------
// 返回数据结构
// ---------------------------------------------------------------------------

// LabelItem 标签详情。
type LabelItem struct {
	Op             int    `json:"op"`
	BDeleted       int    `json:"bDeleted"`
	CreateTime     int64  `json:"create_time"`
	LabelGroupID   int64  `json:"label_groupid"`
	LabelType      int    `json:"label_type"`
	SourceAppID    int    `json:"source_appid"`
	BusinessType   int    `json:"business_type"`
	Name           string `json:"name"`
	DataType       int    `json:"data_type"`
	ID             int64  `json:"id"`
	ServiceGroupID int64  `json:"service_groupid"`
	Seq            int64  `json:"seq,omitempty"`
	NewOrder       int    `json:"new_order"`
	Order          int    `json:"order"`
}

// LabelListResp 获取标签列表的返回结果。
type LabelListResp struct {
	Data struct {
		LabelList []LabelItem `json:"labellist"`
		Index     int64       `json:"index"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AddLabelResp 添加标签的返回结果。
type AddLabelResp struct {
	Data    LabelItem `json:"data"`
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
}

// EditLabelResp 修改标签的返回结果。
type EditLabelResp struct {
	Data    LabelItem `json:"data"`
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
}

// DelLabelResp 删除标签的返回结果。
type DelLabelResp struct {
	Data    LabelItem `json:"data"`
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
}

// LabelAddUsersResp 标签关联用户的返回结果。
type LabelAddUsersResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// UserAddLabelsResp 用户关联标签的返回结果。
type UserAddLabelsResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 调用接口
// ---------------------------------------------------------------------------

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

// LabelAddUsers 为一个标签批量关联/取消关联用户。
// 端点: POST /wxwork/LabelAddUserReq
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - labelID: 标签 ID
//   - userIDList: 要关联的用户 ID 列表
//   - delUserID: 要取消关联的用户 ID 列表
//
// 返回: *LabelAddUsersResp 操作结果
func (s *SdkIPad) LabelAddUsers(host, uuid string, labelID int64, userIDList, delUserID []int64) (*LabelAddUsersResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"labelid":     labelID,
		"userid_list": userIDList,
		"del_userid":  delUserID,
	}
	var resp LabelAddUsersResp
	if err := s.httpClient.Post(host+"/wxwork/LabelAddUserReq", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
