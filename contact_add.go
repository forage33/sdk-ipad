// contact_add.go 联系人查找与添加相关接口（加好友链路）。
// 提供手机号搜索、搜索添加、名片添加、好友申请与同意、企微搜索添加、
// 共享好友添加、好友共享（转给同事）等功能。

package sdk_ipad

import (
	"errors"

	"go.uber.org/zap"
)

// ---------------------------------------------------------------------------
// 返回数据结构
// ---------------------------------------------------------------------------

// SearchPhoneResp 手机号搜索联系人的返回结果。
type SearchPhoneResp struct {
	Data    SearchPhoneData `json:"data"`
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
}

// SearchPhoneData 搜索结果数据。
type SearchPhoneData struct {
	UserList []UserInfo `json:"UserList"`
}

// UserInfo 搜索结果中的用户信息。
type UserInfo struct {
	HeadImg string `json:"headImg"`
	Ticket  string `json:"ticket"`
	UserId  int64  `json:"user_id"`
	Sex     int    `json:"sex"`
	Name    string `json:"name"`
	State   string `json:"state"`
	CorpId  int64  `json:"corp_id,omitempty"`
	Openid  string `json:"openid,omitempty"`
}

// AddSearchResp 添加搜索结果好友的返回结果。
type AddSearchResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AddCardResp 通过名片添加好友的返回结果。
type AddCardResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// OperCustomerResp 好友操作（共享/转发）的返回结果。
type OperCustomerResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

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

// AddShareUserResp 添加共享联系人的返回结果。
type AddShareUserResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 调用接口
// ---------------------------------------------------------------------------

// SearchContactByPhone 根据手机号搜索联系人。
// 端点: POST /wxwork/SearchContact
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - phoneNum: 手机号
//
// 返回: *SearchPhoneResp 搜索结果列表
func (s *SdkIPad) SearchContactByPhone(host, uuid, phoneNum string) (*SearchPhoneResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	// 请求参数
	reqParams := map[string]string{
		"uuid":        uuid,
		"phoneNumber": phoneNum,
	}

	s.logger.Info("SearchContactByPhone 请求", zap.String("host", host), zap.String("wx_uuid", uuid), zap.String("phone", phoneNum))

	var searchResp SearchPhoneResp
	if err := s.httpClient.Post(host+"/wxwork/SearchContact", reqParams, &searchResp); err != nil {
		s.logger.Error("SearchContactByPhone HTTP请求失败", zap.String("host", host), zap.String("wx_uuid", uuid), zap.String("phone", phoneNum), zap.Error(err))
		return nil, err
	}

	s.logger.Info("SearchContactByPhone 响应", zap.Int("errcode", searchResp.Errcode), zap.String("errmsg", searchResp.Errmsg), zap.Int("user_count", len(searchResp.Data.UserList)))
	// 使用HTTP客户端发送请求
	return &searchResp, nil
}

// AddSearch 通过搜索结果添加好友。
// 端点: POST /wxwork/AddSearch
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 目标用户 VID
//   - optionID: 操作 ID
//   - phone: 手机号
//   - content: 好友申请验证消息
//   - ticket: 搜索结果 ticket
//
// 返回: *AddSearchResp 添加结果
func (s *SdkIPad) AddSearch(host, uuid string, vid int64, optionID, phone, content, ticket string) (*AddSearchResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":     uuid,
		"vid":      vid,
		"optionid": optionID,
		"phone":    phone,
		"content":  content,
		"ticket":   ticket,
	}

	s.logger.Info("AddSearch 请求", zap.String("host", host), zap.String("wx_uuid", uuid), zap.Int64("vid", vid), zap.String("option_id", optionID), zap.String("phone", phone), zap.String("ticket", ticket))

	var addSearchResp AddSearchResp
	if err := s.httpClient.Post(host+"/wxwork/AddSearch", reqParams, &addSearchResp); err != nil {
		s.logger.Error("AddSearch HTTP请求失败", zap.String("host", host), zap.String("wx_uuid", uuid), zap.Error(err))
		return nil, err
	}

	s.logger.Info("AddSearch 响应", zap.Int("errcode", addSearchResp.Errcode), zap.String("errmsg", addSearchResp.Errmsg))
	return &addSearchResp, nil
}

// AddCard 通过名片添加好友。
// 端点: POST /wxwork/AddBusinessCard
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 目标用户 VID
//   - recommendId: 推荐人 ID
//   - roomID: 来源群 ID（非群来源传 0）
//   - content: 好友申请验证消息
//
// 返回: *AddCardResp 添加结果
func (s *SdkIPad) AddCard(host, uuid string, vid int64, recommendId int64, roomID int64, content string) (*AddCardResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"vid":         vid,
		"recommendId": recommendId,
		"roomid":      roomID,
		"content":     content,
	}

	var addCardResp AddCardResp
	if err := s.httpClient.Post(host+"/wxwork/AddBusinessCard", reqParams, &addCardResp); err != nil {
		return nil, err
	}

	return &addCardResp, nil
}

// AgreeUser 同意好友申请。
// 端点: POST /wxwork/AgreeUser
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 申请人 VID
//   - corpid: 申请人企业 ID
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) AgreeUser(host, uuid string, vid int64, corpid int64) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"vid":    vid,
		"corpid": corpid,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/AgreeUser", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// AddSearchWxwork 搜索并添加企业微信用户为好友。
// 端点: POST /wxwork/AddSearchWxwork
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 目标用户 VID
//   - corpid: 目标用户企业 ID
//   - ticket: 搜索 ticket
//   - content: 好友申请验证消息
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) AddSearchWxwork(host, uuid string, vid, corpid int64, ticket, content string) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"vid":     vid,
		"corpid":  corpid,
		"ticket":  ticket,
		"content": content,
	}

	s.logger.Info("AddSearchWxwork 请求", zap.String("host", host), zap.String("wx_uuid", uuid), zap.Int64("vid", vid), zap.Int64("corpid", corpid), zap.String("ticket", ticket))

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/AddSearchWxwork", reqParams, &resp); err != nil {
		s.logger.Error("AddSearchWxwork HTTP请求失败", zap.String("host", host), zap.String("wx_uuid", uuid), zap.Error(err))
		return nil, err
	}

	s.logger.Info("AddSearchWxwork 响应", zap.Int("errcode", resp.Errcode), zap.String("errmsg", resp.Errmsg), zap.Any("data", resp.Data))
	return &resp, nil
}

// AddShareUser 添加同事共享过来的用户好友。
// 端点: POST /wxwork/AddShareUser
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 同事共享过来的用户 VID
//   - content: 添加好友时的验证消息
//
// 返回: *AddShareUserResp 添加结果
func (s *SdkIPad) AddShareUser(host, uuid string, vid int64, content string) (*AddShareUserResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"vid":     vid,
		"content": content,
	}

	var resp AddShareUserResp
	if err := s.httpClient.Post(host+"/wxwork/AddShareUser", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddWxUser 直接添加好友（仅对曾被删除的好友有效）。
// 端点: POST /wxwork/addWxUser
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 目标用户 VID
//   - content: 好友申请验证消息
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) AddWxUser(host, uuid string, vid int64, content string) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"vid":     vid,
		"content": content,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/addWxUser", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// OperCustomer 好友操作（共享给其他企微号）。
// 端点: POST /wxwork/OperCustomer
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - thisVID: 当前账号 VID
//   - recommendVID: 要共享的好友 user_id
//   - sendVID: 要共享给哪个同事的 VID
//
// 返回: *OperCustomerResp 操作结果
func (s *SdkIPad) OperCustomer(host, uuid, thisVID, recommendVID, sendVID string) (*OperCustomerResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":         uuid,
		"thisvid":      thisVID,
		"recommendvid": recommendVID,
		"sendvid":      sendVID,
	}

	var resp OperCustomerResp
	if err := s.httpClient.Post(host+"/wxwork/OperCustomer", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
