// group.go 群相关接口。
// 提供群列表获取、群成员管理、建群、群公告、群名修改、邀请进群、
// 群二维码、欢迎语管理、退群、解散群等功能。

package sdk_ipad

import (
	"encoding/json"
	"errors"
	"strings"
)

// ---------------------------------------------------------------------------
// 返回数据结构
// ---------------------------------------------------------------------------

// RoomSessionResp 会话中群列表的返回结果。
type RoomSessionResp struct {
	Data struct {
		Seq      int `json:"seq"`
		RoomList []struct {
			RoomId       string        `json:"room_id"`
			CreateUserId int64         `json:"create_user_id"`
			Total        int           `json:"total"`
			Flag         int           `json:"flag"`
			CreateTime   int64         `json:"create_time"`
			Nickname     string        `json:"nickname"`
			Managers     []interface{} `json:"managers"`
		} `json:"room_list"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ChatroomMembersResp 客户群列表的返回结果。
type ChatroomMembersResp struct {
	Data struct {
		Total     int `json:"total"`
		NextStart int `json:"next_start"`
		RoomList  []struct {
			RoomId       string `json:"room_id"`
			CreateUserId int64  `json:"create_user_id"`
			Infoticket   string `json:"infoticket"`
			UpdateTime   int64  `json:"update_time"`
			Total        int    `json:"total"`
			CreateTime   int64  `json:"create_time"`
			Nickname     string `json:"nickname"`
			Roomurl      string `json:"roomurl"`
		} `json:"roomList"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// RoomUserListResp 群成员列表的返回结果（data 为动态 JSON）。
type RoomUserListResp struct {
	Data    json.RawMessage `json:"data"`
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
}

// RoomMember 群成员详细信息。
type RoomMember struct {
	Unionid      string `json:"unionid"`
	CreateTime   int64  `json:"create_time"`
	Sex          int    `json:"sex"`
	Mobile       string `json:"mobile"`
	Acctid       string `json:"acctid"`
	JoinScene    int    `json:"join_scene"`
	Avatar       string `json:"avatar"`
	EnglishName  string `json:"english_name"`
	Realname     string `json:"realname"`
	RoomNotes    string `json:"room_notes"`
	Jointime     int64  `json:"jointime"`
	Nickname     string `json:"nickname"`
	RoomNickname string `json:"room_nickname"`
	Position     string `json:"position"`
	Uin          int64  `json:"uin"`
	InviteUserId int64  `json:"invite_user_id"`
	CorpId       int64  `json:"corp_id"`
}

// CreateRoomResp 创建群聊的返回结果。
type CreateRoomResp struct {
	Data struct {
		Roomname   string      `json:"roomname"`
		CreateTime int64       `json:"createTime"`
		Createid   int64       `json:"createid"`
		Roomid     interface{} `json:"roomid"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AgreeAddRoomResp 同意进群的返回结果。
type AgreeAddRoomResp struct {
	Data struct {
		Data struct {
			Vid    string `json:"vid"`
			Status string `json:"status"`
		} `json:"data"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AgreeCardAddRoomResp 二维码进群的返回结果。
type AgreeCardAddRoomResp struct {
	Data struct {
		Data struct {
			Vid    string `json:"vid"`
			Status string `json:"status"`
		} `json:"data"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// CommonGroupResp 群操作通用返回结果。
type CommonGroupResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// UpdateRoomNameResp 修改群名的返回结果。
type UpdateRoomNameResp struct {
	Data struct {
		Roomname string `json:"roomname"`
		Roomid   string `json:"roomid"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// WxRoomInviteResp 获取群二维码的返回结果。
type WxRoomInviteResp struct {
	Data struct {
		QrCodePath string      `json:"QrCodePath"`
		ImageUrl   string      `json:"image_url"`
		Roomid     json.Number `json:"roomid"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// NormalizeQrCodePath 规范化群二维码图片 URL。
// 协议服务可能返回反斜杠分隔的 Windows 风格路径（如 http://host/download\RoomQrCode\xx.jpg），
// 统一替换为正斜杠，保证前端 <img src> 可正常加载；同时兼容 URL 编码的反斜杠
// （RFC 3986 百分号编码十六进制大小写不敏感，%5C 与 %5c 均需处理）。
func NormalizeQrCodePath(path string) string {
	path = strings.ReplaceAll(path, "%5C", "/")
	path = strings.ReplaceAll(path, "%5c", "/")
	return strings.ReplaceAll(path, "\\", "/")
}

// WelcomeMsgItem 欢迎语列表项。
type WelcomeMsgItem struct {
	Updateid       int64  `json:"updateid"`
	Createtime     int64  `json:"createtime"`
	Createid       int64  `json:"createid"`
	FileId         string `json:"file_id"`
	AesKey         string `json:"aes_key"`
	Id             int    `json:"id"`
	Type           int    `json:"type"`
	ImgSize        int    `json:"img_size"`
	Updatetime     int64  `json:"updatetime"`
	Content        string `json:"content"`
	WelcineMsgLink string `json:"WelcineMsgLink"`
	Md5            string `json:"md5"`
}

// GetWelcomeMsgListResp 获取欢迎语列表的返回结果。
type GetWelcomeMsgListResp struct {
	Data struct {
		Offset          string           `json:"offset"`
		WelocomeMsgList []WelcomeMsgItem `json:"WelocomeMsgList"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AddWelcomeMsgResp 添加欢迎语的返回结果。
type AddWelcomeMsgResp struct {
	Data struct {
		Id int `json:"id"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// OutRoomReqResp 退出群聊的返回结果。
type OutRoomReqResp struct {
	Data struct {
		RoomId string `json:"room_id"`
		MsgId  int    `json:"msg_id"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// RoomDesUserItem 群成员简要信息（简洁版）。
type RoomDesUserItem struct {
	RoomNotes    string `json:"room_notes"`
	Jointime     int64  `json:"jointime"`
	Nickname     string `json:"nickname"`
	JoinScene    int    `json:"join_scene"`
	Uin          int64  `json:"uin"`
	InviteUserId int64  `json:"invite_user_id"`
}

// GetRoomDesUserListResp 获取群成员简洁列表的返回结果。
type GetRoomDesUserListResp struct {
	Data struct {
		RoomId          string            `json:"room_id"`
		CreateUserId    int64             `json:"create_user_id"`
		Total           int               `json:"total"`
		NoticeSendervid int64             `json:"notice_sendervid"`
		CreateTime      int64             `json:"create_time"`
		NoticeContent   string            `json:"notice_content"`
		Nickname        string            `json:"nickname"`
		Flag            int               `json:"flag"`
		NewFlag         int               `json:"new_flag"`
		MemberList      []RoomDesUserItem `json:"member_list"`
		NoticeTime      int64             `json:"notice_time"`
		Managers        []interface{}     `json:"managers"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetRoomInfoResp 获取群详细信息和头像的返回结果（data 为动态 JSON）。
type GetRoomInfoResp struct {
	Data    json.RawMessage `json:"data"`
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
}

// WxRoomHeaderResp 获取群头像的返回结果。
type WxRoomHeaderResp struct {
	Data struct {
		ImageUrl string      `json:"image_url"`
		Roomid   json.Number `json:"roomid"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 调用接口
// ---------------------------------------------------------------------------

// GetSessionRoomList 获取会话中的群列表。
// 端点: POST /wxwork/GetSessionRoomList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - limit: 每页数量
//   - starIndex: 起始索引
//
// 返回: *RoomSessionResp 群列表
func (s *SdkIPad) GetSessionRoomList(host, uuid string, limit, starIndex int) (*RoomSessionResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":       uuid,
		"limit":      limit,
		"star_index": starIndex,
	}
	var resp RoomSessionResp
	if err := s.httpClient.Post(host+"/wxwork/GetSessionRoomList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetChatroomMembers 获取客户群列表。
// 端点: POST /wxwork/GetChatroomMembers
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - limit: 每页数量
//   - startIndex: 起始索引
//
// 返回: *ChatroomMembersResp 客户群列表
func (s *SdkIPad) GetChatroomMembers(host, uuid string, limit, startIndex int) (*ChatroomMembersResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":       uuid,
		"limit":      limit,
		"star_index": startIndex,
	}
	var resp ChatroomMembersResp
	if err := s.httpClient.Post(host+"/wxwork/GetChatroomMembers", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRoomUserList 获取群成员详细列表。
// 端点: POST /wxwork/GetRoomUserList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *RoomUserListResp 群成员列表（Data 为 json.RawMessage，需自行解析）
func (s *SdkIPad) GetRoomUserList(host, uuid string, roomID int64) (*RoomUserListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp RoomUserListResp
	if err := s.httpClient.Post(host+"/wxwork/GetRoomUserList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRoomWx 创建外部群聊。
// 端点: POST /wxwork/CreateRoomWx
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomName: 群名称
//   - vids: 初始成员 VID 列表
//
// 返回: *CreateRoomResp 包含新群 ID 和创建信息
func (s *SdkIPad) CreateRoomWx(host, uuid, roomName string, vids []int64) (*CreateRoomResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":     uuid,
		"roomName": roomName,
		"vids":     vids,
	}
	var resp CreateRoomResp
	if err := s.httpClient.Post(host+"/wxwork/CreateRoomWx", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AgreeAddRoom 同意群邀请链接进群。
// 端点: POST /wxwork/AgreeAddRoom
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - url: 群邀请链接
//
// 返回: *AgreeAddRoomResp 进群结果
func (s *SdkIPad) AgreeAddRoom(host, uuid, url string) (*AgreeAddRoomResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid": uuid,
		"url":  url,
	}
	var resp AgreeAddRoomResp
	if err := s.httpClient.Post(host+"/wxwork/AgreeAddRoom", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AgreeCardAddRoom 通过二维码进群。
// 端点: POST /wxwork/AgreeCardAddRoom
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - code: 群二维码 code
//
// 返回: *AgreeCardAddRoomResp 进群结果
func (s *SdkIPad) AgreeCardAddRoom(host, uuid, code string) (*AgreeCardAddRoomResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid": uuid,
		"code": code,
	}
	var resp AgreeCardAddRoomResp
	if err := s.httpClient.Post(host+"/wxwork/AgreeCardAddRoom", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRoomNei 创建内部群聊。
// 端点: POST /wxwork/CreateRoomNei
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomName: 群名称
//   - vids: 初始成员 VID 列表
//
// 返回: *CreateRoomResp 包含新群 ID 和创建信息
func (s *SdkIPad) CreateRoomNei(host, uuid, roomName string, vids []int64) (*CreateRoomResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":     uuid,
		"vids":     vids,
		"roomName": roomName,
	}
	var resp CreateRoomResp
	if err := s.httpClient.Post(host+"/wxwork/CreateRoomNei", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendNotice 发送群公告。
// 端点: POST /wxwork/SendNotice
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//   - msg: 公告内容
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) SendNotice(host, uuid string, roomID int64, msg string) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
		"msg":    msg,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/SendNotice", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateRoomName 修改群名称。
// 端点: POST /wxwork/UpdateRoomName
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//   - roomName: 新群名称
//
// 返回: *UpdateRoomNameResp 修改结果
func (s *SdkIPad) UpdateRoomName(host, uuid string, roomID int64, roomName string) (*UpdateRoomNameResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":     uuid,
		"roomid":   roomID,
		"roomname": roomName,
	}
	var resp UpdateRoomNameResp
	if err := s.httpClient.Post(host+"/wxwork/UpdateRoomName", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// InvitationToRoomLink 通过链接邀请成员进群。
// 端点: POST /wxwork/InvitationToRoomLink
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//   - vids: 被邀请成员 VID 列表
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) InvitationToRoomLink(host, uuid string, roomID int64, vids []int64) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
		"vids":   vids,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/InvitationToRoomLink", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// InvitationToRoom 直接邀请成员进群。
// 端点: POST /wxwork/InvitationToRoom
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//   - vids: 被邀请成员 VID 列表
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) InvitationToRoom(host, uuid string, roomID int64, vids []int64) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
		"vids":   vids,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/InvitationToRoom", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DelRoomUsers 移除群成员。
// 端点: POST /wxwork/DelRoomUsers
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//   - vids: 要移除的成员 VID 列表
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) DelRoomUsers(host, uuid string, roomID int64, vids []int64) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
		"vids":   vids,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/DelRoomUsers", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetMyRoomName 设置自己在群内的昵称。
// 端点: POST /wxwork/SetMyRoomName
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//   - content: 群内昵称
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) SetMyRoomName(host, uuid string, roomID int64, content string) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"roomid":  roomID,
		"content": content,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/SetMyRoomName", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// TransferChatroomOwner 转让群主。
// 端点: POST /wxwork/TransferChatroomOwner
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//   - vid: 新群主 VID
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) TransferChatroomOwner(host, uuid string, roomID, vid int64) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
		"vid":    vid,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/TransferChatroomOwner", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DissolutionRoom 解散群聊。
// 端点: POST /wxwork/DissolutionRoom
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) DissolutionRoom(host, uuid string, roomID int64) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/DissolutionRoom", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// WxRoomInvite 获取群二维码。
// 端点: POST /wxwork/WxRoomInvite
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *WxRoomInviteResp 包含群二维码路径和图片地址
func (s *SdkIPad) WxRoomInvite(host, uuid string, roomID int64) (*WxRoomInviteResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp WxRoomInviteResp
	if err := s.httpClient.Post(host+"/wxwork/WxRoomInvite", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWelcomeMsgList 获取欢迎语列表。
// 端点: POST /wxwork/GetWelcomeMsgList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - offset: 分页偏移量
//   - limit: 每页数量
//
// 返回: *GetWelcomeMsgListResp 欢迎语列表
func (s *SdkIPad) GetWelcomeMsgList(host, uuid, offset string, limit int) (*GetWelcomeMsgListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"offset": offset,
		"limit":  limit,
	}
	var resp GetWelcomeMsgListResp
	if err := s.httpClient.Post(host+"/wxwork/GetWelcomeMsgList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddWelcomeMsg 添加欢迎语。
// 端点: POST /wxwork/AddWelcomeMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - list: 欢迎语文本列表
//   - url: 链接地址
//   - title: 链接标题
//   - content: 链接描述
//   - headImg: 链接缩略图
//
// 返回: *AddWelcomeMsgResp 包含新增欢迎语 ID
func (s *SdkIPad) AddWelcomeMsg(host, uuid string, list []string, url, title, content, headImg string) (*AddWelcomeMsgResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"list":    list,
		"url":     url,
		"title":   title,
		"content": content,
		"headImg": headImg,
	}
	var resp AddWelcomeMsgResp
	if err := s.httpClient.Post(host+"/wxwork/AddWelcomeMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetWelcomeMsg 为群设置欢迎语。
// 端点: POST /wxwork/SetWelcomeMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - indexId: 欢迎语索引 ID
//   - msgType: 消息类型
//   - roomID: 群 ID
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) SetWelcomeMsg(host, uuid string, indexId, msgType int, roomID int64) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"indexid": indexId,
		"type":    msgType,
		"roomid":  roomID,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/SetWelcomeMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DelWelcomeMsg 取消群欢迎语。
// 端点: POST /wxwork/DelWelcomeMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) DelWelcomeMsg(host, uuid string, roomID int64) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/DelWelcomeMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRoomFriends 添加群成员为好友。
// 端点: POST /wxwork/AddRoomFriends
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//   - corpId: 企业 ID
//   - vid: 目标成员 VID
//   - content: 好友申请验证消息
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) AddRoomFriends(host, uuid string, roomID int64, corpId int, vid int64, content string) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"roomid":  roomID,
		"corpid":  corpId,
		"vid":     vid,
		"content": content,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/AddRoomFriends", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// OutRoomReq 退出群聊。
// 端点: POST /wxwork/OutRoomReq
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *OutRoomReqResp 退出结果
func (s *SdkIPad) OutRoomReq(host, uuid string, roomID int64) (*OutRoomReqResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp OutRoomReqResp
	if err := s.httpClient.Post(host+"/wxwork/OutRoomReq", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRoomDesUserList 获取群成员简洁列表。
// 端点: POST /wxwork/GetRoomDesUserList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *GetRoomDesUserListResp 包含群信息和简洁成员列表
func (s *SdkIPad) GetRoomDesUserList(host, uuid string, roomID int64) (*GetRoomDesUserListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp GetRoomDesUserListResp
	if err := s.httpClient.Post(host+"/wxwork/GetRoomDesUserList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRoomInfo 获取群详细信息和头像。
// 端点: POST /wxwork/GetRoomInfo
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *GetRoomInfoResp 群详情（Data 为 json.RawMessage，需自行解析）
func (s *SdkIPad) GetRoomInfo(host, uuid string, roomID int64) (*GetRoomInfoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp GetRoomInfoResp
	if err := s.httpClient.Post(host+"/wxwork/GetRoomInfo", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// WxRoomHeader 获取群头像。
// 端点: POST /wxwork/WxRoomHeader
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *WxRoomHeaderResp 包含群头像图片地址
func (s *SdkIPad) WxRoomHeader(host, uuid string, roomID int64) (*WxRoomHeaderResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp WxRoomHeaderResp
	if err := s.httpClient.Post(host+"/wxwork/WxRoomHeader", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
