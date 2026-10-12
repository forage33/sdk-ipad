// api_group_invitation_to_room_link.go InvitationToRoomLink 通过链接邀请成员进群。

package sdk_ipad

import "errors"

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
