// api_contact_get_external_contacts.go GetExternalContacts 获取外部联系人（客户）列表。

package sdk_ipad

import "errors"

// ExternalResp 外部联系人列表的返回结果。
type ExternalResp struct {
	Data struct {
		List []struct {
			Unionid         string        `json:"unionid"`
			CreateTime      int64         `json:"create_time"`
			AddCustomerTime int64         `json:"add_customer_time"`
			Sex             int           `json:"sex"`
			Mobile          string        `json:"mobile"`
			CompanyRemark   string        `json:"company_remark"`
			Acctid          string        `json:"acctid"`
			Avatar          string        `json:"avatar"`
			Source          int           `json:"source"`
			EnglishName     string        `json:"english_name"`
			RemarkPhone     []interface{} `json:"remark_phone"`
			Realname        string        `json:"realname"`
			RealRemarks     string        `json:"real_remarks"`
			Labelid         []string      `json:"labelid"`
			UserId          int64         `json:"user_id"`
			Nickname        string        `json:"nickname"`
			Position        string        `json:"position"`
			CorpId          int64         `json:"corp_id"`
			SourceInfo      struct {
				Vid          int    `json:"vid"`
				WxFriendName string `json:"wx_friend_name"`
				ApplyMode    int    `json:"apply_mode"`
				Mobile       string `json:"mobile"`
				SourceRoomid int64  `json:"source_roomid"`
				SourceType   int    `json:"source_type"`
			} `json:"source_info"`
			Remarks string `json:"remarks"`
			Seq     int    `json:"seq"`
			Status  int    `json:"status"`
		} `json:"list"`
		Seq int `json:"seq"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetExternalContacts 获取外部联系人（客户）列表。
// 端点: POST /wxwork/GetExternalContacts
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - limit: 每页数量
//   - seq: 分页序列号（首次传 0）
//
// 返回: *ExternalResp 外部联系人列表
func (s *SdkIPad) GetExternalContacts(host, uuid string, limit, seq int) (*ExternalResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"limit": limit,
		"seq":   seq,
	}
	var resp ExternalResp
	if err := s.httpClient.Post(host+"/wxwork/GetExternalContacts", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
