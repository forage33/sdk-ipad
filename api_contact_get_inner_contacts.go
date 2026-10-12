// api_contact_get_inner_contacts.go GetInnerContacts 获取内部联系人列表（含部门信息）。

package sdk_ipad

import "errors"

// InnerResp 内部联系人列表的返回结果。
type InnerResp struct {
	Data struct {
		List []struct {
			DispOrder    int    `json:"disp_order,omitempty"`
			Corpid       int64  `json:"corpid,omitempty"`
			PartyName    string `json:"party_name,omitempty"`
			CreateTime   int    `json:"create_time"`
			Remark       string `json:"remark,omitempty"`
			Partyid      int64  `json:"partyid"`
			Type         int    `json:"type,omitempty"`
			IsDepartment int    `json:"is_Department"`
			Parentid     int    `json:"parentid,omitempty"`
			Status       int    `json:"status,omitempty"`
			Unionid      string `json:"unionid,omitempty"`
			Sex          int    `json:"sex,omitempty"`
			Mobile       string `json:"mobile,omitempty"`
			Acctid       string `json:"acctid,omitempty"`
			Avatar       string `json:"avatar,omitempty"`
			EnglishName  string `json:"english_name,omitempty"`
			Realname     string `json:"realname,omitempty"`
			SelfAttrInfo []struct {
				FiledName  string `json:"filed_name"`
				FiledValue string `json:"filed_value"`
				FieldId    string `json:"field_id"`
				FieldType  int    `json:"field_type"`
			} `json:"SelfAttrInfo,omitempty"`
			UserId   int64  `json:"user_id,omitempty"`
			Nickname string `json:"nickname,omitempty"`
			Position string `json:"position,omitempty"`
			CorpId   int64  `json:"corp_id,omitempty"`
		} `json:"list"`
		Seq string `json:"seq"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetInnerContacts 获取内部联系人列表（含部门信息）。
// 端点: POST /wxwork/GetInnerContacts
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - limit: 每页数量
//   - strSeq: 分页序列号（首次传空字符串）
//
// 返回: *InnerResp 内部联系人列表
func (s *SdkIPad) GetInnerContacts(host, uuid string, limit int, strSeq string) (*InnerResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"limit":  limit,
		"strSeq": strSeq,
	}
	var resp InnerResp
	if err := s.httpClient.Post(host+"/wxwork/GetInnerContacts", reqParams, &resp); err != nil {
		return nil, err
	}

	// 使用HTTP客户端发送请求
	return &resp, nil
}
