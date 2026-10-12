// api_contact_get_apply_user_list.go GetApplyUserList 获取好友申请列表。

package sdk_ipad

import "errors"

// ApplyUserListResp 好友申请列表的返回结果。
type ApplyUserListResp struct {
	Data struct {
		List []struct {
			ItemFlag        int           `json:"item_flag"`
			Unionid         string        `json:"unionid"`
			CreateTime      int64         `json:"create_time"`
			Sex             int           `json:"sex"`
			Mobile          string        `json:"mobile"`
			CompanyRemark   string        `json:"company_remark"`
			Acctid          string        `json:"acctid"`
			Avatar          string        `json:"avatar"`
			EnglishName     string        `json:"english_name"`
			RemarkPhone     []interface{} `json:"remark_phone"`
			Realname        string        `json:"realname"`
			ApplyCreateTime int64         `json:"apply_create_time"`
			RealRemarks     string        `json:"real_remarks"`
			UserId          int64         `json:"user_id"`
			ApplyUpdateTime int64         `json:"apply_update_time"`
			Nickname        string        `json:"nickname"`
			Position        string        `json:"position"`
			State           int           `json:"state"`
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

// GetApplyUserList 获取好友申请列表。
// 端点: POST /wxwork/GetApplyUserList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - limit: 每页数量
//   - seq: 分页序列号（首次传 0）
//
// 返回: *ApplyUserListResp 好友申请列表
func (s *SdkIPad) GetApplyUserList(host, uuid string, limit, seq int) (*ApplyUserListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"limit": limit,
		"seq":   seq,
	}

	var resp ApplyUserListResp
	if err := s.httpClient.Post(host+"/wxwork/GetApplyUserList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
