// api_group_set_my_room_name.go SetMyRoomName 设置自己在群内的昵称。

package sdk_ipad

import "errors"

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
