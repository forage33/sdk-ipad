// api_message_send_emotion_message.go SendEmotionMessage 发送 GIF 表情消息。

package sdk_ipad

import "errors"

// SendEmotionResp 发送表情消息的返回结果。
type SendEmotionResp struct {
	Data struct {
		Receiver    int64  `json:"receiver"`
		Sender      int64  `json:"sender"`
		EmotionType string `json:"EmotionType"`
		SenderName  string `json:"sender_name"`
		IsRoom      int    `json:"is_room"`
		Sendtime    int    `json:"sendtime"`
		MsgId       int    `json:"msg_id"`
		AppInfo     string `json:"app_info"`
		ServerId    int    `json:"server_id"`
		Msgtype     int    `json:"msgtype"`
		Url         string `json:"url"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SendEmotionMessage 发送 GIF 表情消息。
// 端点: POST /wxwork/SendEmotionMessage
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - imgurl: 表情图片地址
//   - width: 表情宽度（像素）
//   - height: 表情高度（像素）
//   - emotionType: 表情类型
//
// 返回: *SendEmotionResp 发送结果
func (s *SdkIPad) SendEmotionMessage(host, uuid string, sendUserId int64, isRoom bool, imgurl string, width, height, emotionType int) (*SendEmotionResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"send_userid": sendUserId,
		"isRoom":      isRoom,
		"imgurl":      imgurl,
		"width":       width,
		"height":      height,
		"emotionType": emotionType,
	}
	var resp SendEmotionResp
	if err := s.httpClient.Post(host+"/wxwork/SendEmotionMessage", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
