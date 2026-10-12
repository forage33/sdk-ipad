// api_message_send_cdn_big_video_msg.go SendCDNBigVideoMsg 发送大视频消息（需走大文件上传流程）。

package sdk_ipad

import "errors"

// SendCDNBigVideoResp 发送大视频消息的返回结果。
type SendCDNBigVideoResp struct {
	Data struct {
		Receiver   int64  `json:"receiver"`
		CdnKey     string `json:"cdn_Key"`
		SenderName string `json:"sender_name"`
		FileName   string `json:"file_name"`
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

// SendCDNBigVideoMsg 发送大视频消息（需走大文件上传流程）。
// 端点: POST /wxwork/SendCDNBigVideoMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - cdnKey: CDN 文件标识
//   - md5: 文件 MD5 校验值
//   - fileName: 文件名
//   - imgurl: 视频封面图地址
//   - videoDuration: 视频时长（秒）
//   - fileSize: 文件大小（字节）
//
// 返回: *SendCDNBigVideoResp 包含消息 ID 等发送详情
func (s *SdkIPad) SendCDNBigVideoMsg(host, uuid string, sendUserId int64, isRoom bool, cdnKey, md5, fileName, imgurl string, videoDuration int, fileSize int64) (*SendCDNBigVideoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":           uuid,
		"send_userid":    sendUserId,
		"cdnkey":         cdnKey,
		"md5":            md5,
		"file_name":      fileName,
		"imgurl":         imgurl,
		"video_duration": videoDuration,
		"fileSize":       fileSize,
		"isRoom":         isRoom,
	}
	var resp SendCDNBigVideoResp
	if err := s.httpClient.Post(host+"/wxwork/SendCDNBigVideoMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
