// api_group_get_room_user_list.go GetRoomUserList 获取群成员详细列表。

package sdk_ipad

import (
	"encoding/json"
	"errors"
)

// RoomUserListResp 群成员列表的返回结果（data 为动态 JSON）。
type RoomUserListResp struct {
	Data    json.RawMessage `json:"data"`
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
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
