// api_contact_get_user_info_by_vids.go GetUserInfoByVids 根据用户 VID 批量获取详细信息。

package sdk_ipad

import "errors"

// UserInfoResp 批量获取用户详细信息的返回结果。
type UserInfoResp struct {
	Data []struct {
		Unionid     string `json:"unionid"`
		CreateTime  int64  `json:"create_time"`
		UserId      int64  `json:"user_id"`
		Sex         int    `json:"sex"`
		Mobile      string `json:"mobile"`
		Nickname    string `json:"nickname"`
		Acctid      string `json:"acctid"`
		Avatar      string `json:"avatar"`
		Position    string `json:"position"`
		CorpId      int64  `json:"corp_id"`
		EnglishName string `json:"english_name"`
		Realname    string `json:"realname"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetUserInfoByVids 根据用户 VID 批量获取详细信息。
// 端点: POST /wxwork/GetUserInfoByVids
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vIDs: 用户 VID 列表
//
// 返回: *UserInfoResp 用户详细信息列表
func (s *SdkIPad) GetUserInfoByVids(host, uuid string, vIDs []int64) (*UserInfoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	// 请求参数
	reqParams := map[string]interface{}{
		"uuid": uuid,
		"vids": vIDs,
	}
	var resp UserInfoResp
	if err := s.httpClient.Post(host+"/wxwork/GetUserInfoByVids", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
