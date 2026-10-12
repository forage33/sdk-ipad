// api_group_get_room_des_user_list.go GetRoomDesUserList 获取群成员简洁列表。

package sdk_ipad

import "errors"

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
