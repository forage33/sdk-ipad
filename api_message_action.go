// api_message_action.go 消息操作相关接口（发送之外的消息与会话操作）。
// 提供消息撤回、标为已读、语音转文字、消息记录同步等功能。

package sdk_ipad

import (
	"errors"
)

// ---------------------------------------------------------------------------
// 返回数据结构
// ---------------------------------------------------------------------------

// RevokeResp 撤回消息的返回结果。
type RevokeResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// TextEntityResp 语音转文字的返回结果。
type TextEntityResp struct {
	Data struct {
		Text string `json:"text"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// MarkAsReadResp 消息已读的返回结果。
type MarkAsReadResp struct {
	Data struct {
		Receiver           int    `json:"receiver"`
		Sender             int64  `json:"sender"`
		SenderName         string `json:"sender_name"`
		AppInfo            string `json:"app_info"`
		RoomConversationId string `json:"room_conversation_id"`
		IsRoom             int    `json:"is_room"`
		Sendtime           int    `json:"sendtime"`
		MsgId              int    `json:"msg_id"`
		ServerId           int    `json:"server_id"`
		Msgtype            int    `json:"msgtype"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SyncAllDataResp 同步消息记录的返回结果。
type SyncAllDataResp struct {
	Data struct {
		IsSelect int    `json:"is_select"`
		Seq      int    `json:"seq"`
		Content  string `json:"content"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 调用接口
// ---------------------------------------------------------------------------

// RevokeMsg 撤回消息。
// 端点: POST /wxwork/RevokeMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - msgID: 消息 ID
//   - roomID: 会话 ID（好友或群 ID）
//
// 返回: *RevokeResp 撤回结果
//
// 注意: 与微信用户或群内含微信号的会话只能撤回 2 分钟内的消息
func (s *SdkIPad) RevokeMsg(host, uuid string, msgID, roomID int64) (*RevokeResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"msgid":  msgID,
		"roomid": roomID,
	}
	var resp RevokeResp
	if err := s.httpClient.Post(host+"/wxwork/RevokeMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SpeechToTextEntity 语音转文字。
// 端点: POST /wxwork/SpeechToTextEntity
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - msgID: 语音消息 ID
//
// 返回: *TextEntityResp 包含转换后的文本
func (s *SdkIPad) SpeechToTextEntity(host, uuid string, msgID int) (*TextEntityResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"msgid": msgID,
	}
	var entity TextEntityResp
	if err := s.httpClient.Post(host+"/wxwork/SpeechToTextEntity", reqParams, &entity); err != nil {
		return nil, err
	}
	return &entity, nil
}

// MarkAsRead 标记消息为已读。
// 端点: POST /wxwork/MarkAsRead
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 好友或群 ID
//   - isRoom: 是否为群消息
//
// 返回: *MarkAsReadResp 操作结果
func (s *SdkIPad) MarkAsRead(host, uuid string, sendUserId int64, isRoom bool) (*MarkAsReadResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"send_userid": sendUserId, // 好友或者群id
		"isRoom":      isRoom,
	}
	var resp MarkAsReadResp
	if err := s.httpClient.Post(host+"/wxwork/MarkAsRead", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncAllData 同步消息记录。
// 端点: POST /wxwork/SyncAllData
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - limit: 单次拉取消息条数
//   - seq: 消息序列号（从上次同步的位置继续）
//
// 返回: *SyncAllDataResp 包含同步到的消息内容和新的 seq
func (s *SdkIPad) SyncAllData(host, uuid string, limit, seq int) (*SyncAllDataResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"limit": limit,
		"seq":   seq,
	}
	var resp SyncAllDataResp
	if err := s.httpClient.Post(host+"/wxwork/SyncAllData", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
