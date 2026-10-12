// api_message_send_business_card_msg.go SendBusinessCardMsg 发送名片消息。

package sdk_ipad

import "errors"

// SendBusinessCardResp 发送名片消息的返回结果。
type SendBusinessCardResp struct {
	Data struct {
		Receiver       int64  `json:"receiver"`
		HeadImg        string `json:"headImg"`
		SenderName     string `json:"sender_name"`
		IsRoom         int    `json:"is_room"`
		ServerId       int    `json:"server_id"`
		Sender         int64  `json:"sender"`
		Nickname       string `json:"nickname"`
		Sendtime       int    `json:"sendtime"`
		AppInfo        string `json:"app_info"`
		MsgId          int    `json:"msg_id"`
		Msgtype        int    `json:"msgtype"`
		EnterpriseName string `json:"enterpriseName"`
		CorpId         int    `json:"corp_id"`
		Username       int64  `json:"username"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SendBusinessCardMsg 发送名片消息。
// 端点: POST /wxwork/SendBusinessCardMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - headImg: 名片头像地址
//   - username: 名片用户 ID
//   - nickname: 名片昵称
//   - enterpriseName: 企业名称
//
// 返回: *SendBusinessCardResp 发送结果
func (s *SdkIPad) SendBusinessCardMsg(host, uuid string, sendUserId int64, isRoom bool, headImg string, username int64, nickname, enterpriseName string) (*SendBusinessCardResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":           uuid,
		"send_userid":    sendUserId,
		"isRoom":         isRoom,
		"headImg":        headImg,
		"username":       username,
		"nickname":       nickname,
		"enterpriseName": enterpriseName,
	}
	var resp SendBusinessCardResp
	if err := s.httpClient.Post(host+"/wxwork/SendBusinessCardMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
