// api_group_get_chatroom_members.go GetChatroomMembers 获取客户群列表。

package sdk_ipad

import "errors"

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
