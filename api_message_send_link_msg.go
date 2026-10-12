// api_message_send_link_msg.go SendLinkMsg 发送链接卡片消息。

package sdk_ipad

import "errors"

// SendLinkMsgResp 发送链接卡片消息的返回结果。
type SendLinkMsgResp struct {
	Data struct {
		Des        string `json:"des"`
		Receiver   int64  `json:"receiver"`
		HeadImg    string `json:"headImg"`
		Sender     int64  `json:"sender"`
		SenderName string `json:"sender_name"`
		AppInfo    string `json:"app_info"`
		IsRoom     int    `json:"is_room"`
		Sendtime   int    `json:"sendtime"`
		MsgId      int    `json:"msg_id"`
		ServerId   int    `json:"server_id"`
		Title      string `json:"title"`
		Msgtype    int    `json:"msgtype"`
		Url        string `json:"url"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SendLinkMsg 发送链接卡片消息。
// 端点: POST /wxwork/SendLinkMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - url: 链接地址
//   - title: 卡片标题
//   - content: 卡片描述
//   - imgurl: 卡片缩略图地址
//
// 返回: *SendLinkMsgResp 包含消息 ID 等发送详情
func (s *SdkIPad) SendLinkMsg(host, uuid string, sendUserId int64, isRoom bool, url, title, content, imgurl string) (*SendLinkMsgResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"send_userid": sendUserId,
		"url":         url,
		"title":       title,
		"content":     content,
		"imgurl":      imgurl,
		"isRoom":      isRoom,
	}
	var resp SendLinkMsgResp
	if err := s.httpClient.Post(host+"/wxwork/SendLinkMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
