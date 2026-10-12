// api_group_del_room_users.go DelRoomUsers 移除群成员。

package sdk_ipad

import "errors"

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
