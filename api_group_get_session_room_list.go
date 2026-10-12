// api_group_get_session_room_list.go GetSessionRoomList 获取会话中的群列表。

package sdk_ipad

import "errors"

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
