// api_group_transfer_chatroom_owner.go TransferChatroomOwner 转让群主。

package sdk_ipad

import "errors"

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
