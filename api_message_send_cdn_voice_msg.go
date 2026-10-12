// api_message_send_cdn_voice_msg.go SendCDNVoiceMsg 发送 CDN 语音消息。

package sdk_ipad

import "errors"

// SendCDNVoiceResp 发送 CDN 语音消息的返回结果。
type SendCDNVoiceResp struct {
	Data struct {
		Receiver   int64  `json:"receiver"`
		CdnKey     string `json:"cdn_key"`
		FileName   string `json:"file_name"`
		SenderName string `json:"sender_name"`
		IsRoom     int    `json:"is_room"`
		ServerId   int    `json:"server_id"`
		Size       int    `json:"size"`
		AppInfo    string `json:"app_info"`
		Sender     int64  `json:"sender"`
		AesKey     string `json:"aes_key"`
		Sendtime   int    `json:"sendtime"`
		MsgId      int    `json:"msg_id"`
		Msgtype    int    `json:"msgtype"`
		Md5        string `json:"md5"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SendCDNVoiceMsg 发送 CDN 语音消息。
// 端点: POST /wxwork/SendCDNVoiceMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - cdnKey: CDN 文件标识
//   - aesKey: AES 解密密钥
//   - md5: 文件 MD5 校验值
//   - voiceTime: 语音时长（秒）
//   - fileSize: 文件大小（字节）
//
// 返回: *SendCDNVoiceResp 包含消息 ID 等发送详情
func (s *SdkIPad) SendCDNVoiceMsg(host, uuid string, sendUserId int64, isRoom bool, cdnKey, aesKey, md5 string, voiceTime int, fileSize int64) (*SendCDNVoiceResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"send_userid": sendUserId,
		"cdnkey":      cdnKey,
		"aeskey":      aesKey,
		"md5":         md5,
		"voice_time":  voiceTime,
		"fileSize":    fileSize,
		"isRoom":      isRoom,
	}
	var resp SendCDNVoiceResp
	if err := s.httpClient.Post(host+"/wxwork/SendCDNVoiceMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
