// api_group_out_room_req.go OutRoomReq 退出群聊。

package sdk_ipad

import "errors"

// OutRoomReqResp 退出群聊的返回结果。
type OutRoomReqResp struct {
	Data struct {
		RoomId string `json:"room_id"`
		MsgId  int    `json:"msg_id"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// OutRoomReq 退出群聊。
// 端点: POST /wxwork/OutRoomReq
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *OutRoomReqResp 退出结果
func (s *SdkIPad) OutRoomReq(host, uuid string, roomID int64) (*OutRoomReqResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp OutRoomReqResp
	if err := s.httpClient.Post(host+"/wxwork/OutRoomReq", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
