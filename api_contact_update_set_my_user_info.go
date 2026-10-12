// api_contact_update_set_my_user_info.go UpdateSetMyUserInfo 修改联系人备注、手机号、简介和企业等信息。

package sdk_ipad

import "errors"

// UpdateSetMyUserInfo 修改联系人备注、手机号、简介和企业等信息。
// 端点: POST /wxwork/UpdateSetMyUserInfo
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人 VID
//   - remark: 备注名
//   - compny: 企业名称
//   - labelidList: 标签 ID 列表
//   - phone: 电话号码列表
//   - des: 描述/简介
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) UpdateSetMyUserInfo(host, uuid string, vid int64, remark, compny string, labelidList []string, phone []string, des string) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":         uuid,
		"vid":          vid,
		"remark":       remark,
		"compny":       compny,
		"labelid_list": labelidList,
		"phone":        phone,
		"des":          des,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/UpdateSetMyUserInfo", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
