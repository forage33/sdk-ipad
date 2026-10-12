// api_account_get_profile.go GetProfile 获取当前登录账号的详细资料。

package sdk_ipad

import "errors"

// ProfileResp 获取当前账号详情的返回结果。
type ProfileResp struct {
	Data struct {
		UnionID      string `json:"unionid"`
		CreateTime   int64  `json:"create_time"`
		Sex          int    `json:"sex"`
		Mobile       string `json:"mobile"`
		CorpFullName string `json:"corp_full_name"`
		AcctID       string `json:"acctid"`
		SCorpID      string `json:"scorp_id"`
		Avatar       string `json:"avatar"`
		CorpName     string `json:"corp_name"`
		EnglishName  string `json:"english_name"`
		RealName     string `json:"realname"`
		UserID       int64  `json:"user_id"`
		Nickname     string `json:"nickname"`
		Position     string `json:"position"`
		CorpDesc     string `json:"corp_desc"`
		CorpID       int64  `json:"corp_id"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetProfile 获取当前登录账号的详细资料。
// 端点: POST /wxwork/GetProfile
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *ProfileResp 账号详细信息
func (s *SdkIPad) GetProfile(host, uuid string) (*ProfileResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]string{
		"uuid": uuid,
	}
	var resp ProfileResp
	if err := s.httpClient.Post(host+"/wxwork/GetProfile", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
