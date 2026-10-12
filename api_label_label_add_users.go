// api_label_label_add_users.go LabelAddUsers 为一个标签批量关联/取消关联用户。

package sdk_ipad

import "errors"

// LabelAddUsersResp 标签关联用户的返回结果。
type LabelAddUsersResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
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
