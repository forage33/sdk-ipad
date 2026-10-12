// api_account_get_ban_detail.go GetBanDetail 获取账号封禁详情。

package sdk_ipad

import "errors"

// BanResp 获取封禁详情的返回结果。
type BanResp struct {
	Data    BanData `json:"data"`
	Errcode int     `json:"errcode"`
	Errmsg  string  `json:"errmsg"`
}

// BanData 封禁数据。
type BanData struct {
	Vid          string   `json:"vid"`
	Gid          string   `json:"gid"`
	MainData     MainData `json:"mainData"`
	IsWechat     bool     `json:"isWechat"`
	IsWxwork     bool     `json:"isWxwork"`
	IsSuperadmin int      `json:"is_superadmin"`
	OrderId      string   `json:"order_id"`
	CorpId       string   `json:"corp_id"`
}

// MainData 封禁详细信息。
type MainData struct {
	BanStartTime  FlexInt64 `json:"ban_start_time"`
	AppealId      string    `json:"appeal_id"`
	BanSource     int       `json:"ban_source"`
	IsShowTitle   bool      `json:"is_show_title"`
	ErrCode       int       `json:"err_code"`
	BanTitle      string    `json:"ban_title"`
	Type          string    `json:"type"`
	IsShowReason  bool      `json:"is_show_reason"`
	IsShowFunc    bool      `json:"is_show_func"`
	BanFinishTime FlexInt64 `json:"ban_finish_time"`
	BanCategory   int       `json:"ban_category"`
	AppealStatus  int       `json:"appeal_status"`
	CanAppeal     bool      `json:"can_appeal"`
	BanIdType     int       `json:"ban_id_type"`
	NeedLogin     string    `json:"need_login"`
	BanReason     string    `json:"ban_reason"`
	TemplateType  int       `json:"template_type"`
	ExposedId     string    `json:"exposed_id"`
	BanFunc       string    `json:"ban_func"`
	IsShowTime    bool      `json:"is_show_time"`
	CanUrge       bool      `json:"can_urge"`
	Key           string    `json:"key"`
}

// GetBanDetail 获取账号封禁详情。
// 端点: POST /wxwork/getBanDetile
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - url: 封禁详情页地址
//
// 返回: *BanResp 封禁详细信息
func (s *SdkIPad) GetBanDetail(host, uuid, url string) (*BanResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
		"url":  url,
	}

	var banResp BanResp
	if err := s.httpClient.Post(host+"/wxwork/getBanDetile", reqParams, &banResp); err != nil {
		return nil, err
	}

	return &banResp, nil
}
