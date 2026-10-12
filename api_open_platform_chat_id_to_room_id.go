// api_open_platform_chat_id_to_room_id.go ChatIdToRoomId chatid 转群 id。

package sdk_ipad

import "errors"

// ChatIdToRoomIdResp chatid 转群 id 的返回结果。
type ChatIdToRoomIdResp struct {
	Data    RoomChatIdPair `json:"data"`
	Errcode int            `json:"errcode"`
	Errmsg  string         `json:"errmsg"`
}

// ChatIdToRoomId chatid 转群 id。
// 端点: POST /wxwork/ChatIdToRoomId
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - corpID: 企业 id，不传（0）默认当前账号登录企业
//   - chatID: 开放平台 chatid
//
// 返回: *ChatIdToRoomIdResp 含 room_id 与对应 chatid
func (s *SdkIPad) ChatIdToRoomId(host, uuid string, corpID int64, chatID string) (*ChatIdToRoomIdResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"chatid": chatID,
	}
	if corpID != 0 {
		reqParams["corpid"] = corpID
	}
	var resp ChatIdToRoomIdResp
	if err := s.httpClient.Post(host+"/wxwork/ChatIdToRoomId", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
