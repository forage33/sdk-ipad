// api_group_dissolution_room.go DissolutionRoom 解散群聊。

package sdk_ipad

import "errors"

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
