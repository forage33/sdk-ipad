// contact.go 联系人操作相关接口。
// 提供联系人搜索、添加好友、获取联系人信息、内外部联系人列表、
// 删除联系人、好友申请、会话管理、备注修改、置顶/免打扰/标记等功能。

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

// CommonContactResp 联系人操作通用返回结果（data 为 null 或简单数据）。
type CommonContactResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
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

// ConversationItem 会话管理列表项（置顶/免打扰/标记）。
type ConversationItem struct {
	Messagetype int   `json:"messagetype"`
	Id          int64 `json:"id"`
	FwId        int64 `json:"fw_id"`
	State       int   `json:"state"`
}

// ManagedConversationsResp 获取置顶/免打扰/折叠/标记列表的返回结果。
type ManagedConversationsResp struct {
	Data struct {
		TopList         []ConversationItem `json:"top_list"`
		ShieldList      []ConversationItem `json:"shield_list"`
		MarkSessionList []ConversationItem `json:"mark_session_list"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SessionItem 会话列表中的会话项。
type SessionItem struct {
	Unreadcnt   int   `json:"unreadcnt"`
	Sessionid   int64 `json:"sessionid"`
	Beginmsgseq int64 `json:"beginmsgseq"`
	Msgtype     int   `json:"msgtype"`
}

// SessionSimpleItem 会话列表中的简单项（置顶/免打扰）。
type SessionSimpleItem struct {
	Messagetype int   `json:"messagetype"`
	Id          int64 `json:"id"`
}

// SessionListResp 获取会话列表的返回结果。
type SessionListResp struct {
	Data struct {
		TopList     []SessionSimpleItem `json:"top_list"`
		CollectList []interface{}       `json:"collect_list"`
		ShieldList  []SessionSimpleItem `json:"shield_list"`
		Seq         int                 `json:"seq"`
		RoomList    []SessionItem       `json:"room_list"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ColleaRemarkListResp 内部联系人备注列表的返回结果。
type ColleaRemarkListResp struct {
	Data struct {
		List []struct {
			Vid         int64  `json:"vid"`
			Description string `json:"description"`
			Remark      string `json:"remark"`
			Timestamp   int64  `json:"timestamp"`
		} `json:"list"`
		Seq int `json:"seq"`
	} `json:"data"`
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

// DeleteContact 删除联系人。
// 端点: POST /wxwork/DelContact
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人 VID
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) DeleteContact(host, uuid string, vid int64) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid": uuid,
		"vid":  vid,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/DelContact", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// GetManagedConversations 获取置顶/免打扰/折叠/标记列表。
// 端点: POST /wxwork/getManagedConversations
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *ManagedConversationsResp 包含置顶、免打扰和标记列表
func (s *SdkIPad) GetManagedConversations(host, uuid string) (*ManagedConversationsResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]string{
		"uuid": uuid,
	}

	var resp ManagedConversationsResp
	if err := s.httpClient.Post(host+"/wxwork/getManagedConversations", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSessionList 获取会话列表。
// 端点: POST /wxwork/GetSessionList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - limit: 每页数量
//   - starIndex: 起始索引
//
// 返回: *SessionListResp 会话列表
func (s *SdkIPad) GetSessionList(host, uuid string, limit, starIndex int) (*SessionListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":       uuid,
		"limit":      limit,
		"star_index": starIndex,
	}

	var resp SessionListResp
	if err := s.httpClient.Post(host+"/wxwork/GetSessionList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

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

// SetColleagueRemark 修改内部联系人备注。
// 端点: POST /wxwork/SetColleagueRemark
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人 VID
//   - remark: 备注名
//   - des: 描述
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) SetColleagueRemark(host, uuid string, vid int64, remark, des string) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"vid":    vid,
		"remark": remark,
		"des":    des,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/SetColleagueRemark", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetColleaRemarkList 获取内部联系人备注列表。
// 端点: POST /wxwork/GetColleaRemarkList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - limit: 每页数量
//   - seq: 分页序列号
//
// 返回: *ColleaRemarkListResp 备注列表
func (s *SdkIPad) GetColleaRemarkList(host, uuid string, limit, seq int) (*ColleaRemarkListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"limit": limit,
		"seq":   seq,
	}

	var resp ColleaRemarkListResp
	if err := s.httpClient.Post(host+"/wxwork/GetColleaRemarkList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetTop 设置或取消会话置顶。
// 端点: POST /wxwork/SetTop
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人或群 ID
//   - isRoom: 是否为群会话
//   - state: 1=置顶, 0=取消置顶
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) SetTop(host, uuid string, vid int64, isRoom bool, state int) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"vid":    vid,
		"isRoom": isRoom,
		"state":  state,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/SetTop", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetShield 设置或取消免打扰。
// 端点: POST /wxwork/SetShield
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人或群 ID
//   - isRoom: 是否为群会话
//   - state: 1=免打扰, 0=取消免打扰
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) SetShield(host, uuid string, vid int64, isRoom bool, state int) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"vid":    vid,
		"isRoom": isRoom,
		"state":  state,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/SetShield", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetTold 设置或取消免打扰并折叠。
// 端点: POST /wxwork/SetTold
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人或群 ID
//   - isRoom: 是否为群会话
//   - state: 1=折叠, 0=取消折叠
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) SetTold(host, uuid string, vid int64, isRoom bool, state int) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"vid":    vid,
		"isRoom": isRoom,
		"state":  state,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/SetTold", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetMark 设置或取消星标标记。
// 端点: POST /wxwork/SetMark
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人或群 ID
//   - isRoom: 是否为群会话
//   - state: 1=标记, 0=取消标记
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) SetMark(host, uuid string, vid int64, isRoom bool, state int) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"vid":    vid,
		"isRoom": isRoom,
		"state":  state,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/SetMark", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSubContactBlack 拉黑或移除黑名单联系人。
// 端点: POST /wxwork/AddSubContactBlack
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人 VID
//   - blackType: 操作类型（1=拉黑, 0=移除黑名单）
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) AddSubContactBlack(host, uuid string, vid int64, blackType int) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid": uuid,
		"vid":  vid,
		"type": blackType,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/AddSubContactBlack", reqParams, &resp); err != nil {
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

// AddShareUserResp 添加共享联系人的返回结果。
type AddShareUserResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
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
