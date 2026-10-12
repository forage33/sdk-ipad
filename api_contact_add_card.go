// api_contact_add_card.go AddCard 通过名片添加好友。

package sdk_ipad

import "errors"

// AddCardResp 通过名片添加好友的返回结果。
type AddCardResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AddCard 通过名片添加好友。
// 端点: POST /wxwork/AddBusinessCard
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 目标用户 VID
//   - recommendId: 推荐人 ID
//   - roomID: 来源群 ID（非群来源传 0）
//   - content: 好友申请验证消息
//
// 返回: *AddCardResp 添加结果
func (s *SdkIPad) AddCard(host, uuid string, vid int64, recommendId int64, roomID int64, content string) (*AddCardResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"vid":         vid,
		"recommendId": recommendId,
		"roomid":      roomID,
		"content":     content,
	}

	var addCardResp AddCardResp
	if err := s.httpClient.Post(host+"/wxwork/AddBusinessCard", reqParams, &addCardResp); err != nil {
		return nil, err
	}

	return &addCardResp, nil
}
