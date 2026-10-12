// api_message_send_cdn_file.go SendCDNFile 发送 CDN 文件消息。

package sdk_ipad

import "errors"

// SendCDNFileResp 发送 CDN 文件消息的返回结果。
type SendCDNFileResp struct {
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

// SendCDNFile 发送 CDN 文件消息。
// 端点: POST /wxwork/SendCDNFileMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - cdnKey: CDN 文件标识
//   - aesKey: AES 解密密钥
//   - md5: 文件 MD5 校验值
//   - fileName: 文件名
//   - fileSize: 文件大小（字节）
//
// 返回: *SendCDNFileResp 包含消息 ID 等发送详情
func (s *SdkIPad) SendCDNFile(host, uuid string, sendUserId int64, isRoom bool, cdnKey, aesKey, md5, fileName string, fileSize int64) (*SendCDNFileResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"send_userid": sendUserId,
		"cdnkey":      cdnKey,
		"aeskey":      aesKey,
		"md5":         md5,
		"file_name":   fileName,
		"fileSize":    fileSize,
		"isRoom":      isRoom,
	}
	var resp SendCDNFileResp
	if err := s.httpClient.Post(host+"/wxwork/SendCDNFileMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
