// api_account_get_corp_info.go GetCorpInfo 根据企业 ID 获取企业详情。

package sdk_ipad

import "errors"

// CorpInfoItem 企业详情信息。
type CorpInfoItem struct {
	Staffnum       int    `json:"Staffnum"`
	AuthTime       int64  `json:"Auth_time"`
	CardUrl        string `json:"Card_url"`
	Corpid         int64  `json:"Corpid"`
	AdminVid       int64  `json:"Admin_vid"`
	VSuperadminVid int64  `json:"V_superadmin_vid"`
	CorpStat       string `json:"Corp_stat"`
	CorpFullName   string `json:"Corp_full_name"`
	IsAccepted     bool   `json:"IsAccepted"`
	CreateTime     int64  `json:"Create_time"`
	Vid            int64  `json:"Vid"`
	ScorpId        string `json:"Scorp_id"`
	Mail           string `json:"Mail"`
	Ownername      string `json:"Ownername"`
	CorpDesc       string `json:"Corp_desc"`
	CorpName       string `json:"Corp_name"`
	CorpLogo       string `json:"Corp_logo"`
}

// GetCorpInfoResp 获取企业详情的返回结果。
type GetCorpInfoResp struct {
	Data struct {
		List []CorpInfoItem `json:"list"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetCorpInfo 根据企业 ID 获取企业详情。
// 端点: POST /wxwork/GetCorpInfo
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - corpid: 企业 ID
//
// 返回: *GetCorpInfoResp 企业详细信息
func (s *SdkIPad) GetCorpInfo(host, uuid string, corpid int64) (*GetCorpInfoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"corpid": corpid,
	}

	var resp GetCorpInfoResp
	if err := s.httpClient.Post(host+"/wxwork/GetCorpInfo", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
