// api_group_add_room_friends.go AddRoomFriends 添加群成员为好友。

package sdk_ipad

import "errors"

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
