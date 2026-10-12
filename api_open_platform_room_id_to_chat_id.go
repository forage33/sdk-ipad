// api_open_platform_room_id_to_chat_id.go RoomIdToChatId 群 id 转 chatid。

package sdk_ipad

import "errors"

// RoomIdToChatIdResp 群 id 转 chatid 的返回结果。
type RoomIdToChatIdResp struct {
	Data    RoomChatIdPair `json:"data"`
	Errcode int            `json:"errcode"`
	Errmsg  string         `json:"errmsg"`
}

// RoomIdToChatId 群 id 转 chatid。
// 端点: POST /wxwork/RoomIdToChatId
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - corpID: 企业 id，不传（0）默认当前账号登录企业
//   - roomID: 群 id
//
// 返回: *RoomIdToChatIdResp 含 room_id 与对应 chatid
func (s *SdkIPad) RoomIdToChatId(host, uuid string, corpID, roomID int64) (*RoomIdToChatIdResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"room_id": roomID,
	}
	if corpID != 0 {
		reqParams["corpid"] = corpID
	}
	var resp RoomIdToChatIdResp
	if err := s.httpClient.Post(host+"/wxwork/RoomIdToChatId", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
