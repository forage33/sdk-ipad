// api_message_send_cdn_video_msg.go SendCDNVideoMsg 发送 CDN 视频消息。

package sdk_ipad

import "errors"

// SendCDNVideoResp 发送 CDN 视频消息的返回结果。
type SendCDNVideoResp struct {
	Data struct {
		VideoDuration   int    `json:"video_duration"`
		VideoHeight     int    `json:"video_height"`
		VideoWidth      int    `json:"video_width"`
		Receiver        int64  `json:"receiver"`
		OpenimCdnLdSize int    `json:"openim_cdn_ld_size"`
		FileName        string `json:"file_name"`
		SenderName      string `json:"sender_name"`
		CdnKey          string `json:"cdn_Key"`
		IsRoom          int    `json:"is_room"`
		ServerId        int    `json:"server_id"`
		FileSize        int    `json:"file_size"`
		Size            int    `json:"size"`
		Sender          int64  `json:"sender"`
		FileId          string `json:"file_id"`
		AesKey          string `json:"aes_key"`
		AppInfo         string `json:"app_info"`
		Sendtime        int    `json:"sendtime"`
		MsgId           int    `json:"msg_id"`
		Msgtype         int    `json:"msgtype"`
		Md5             string `json:"md5"`
		PreviewImgUrl   string `json:"preview_img_url"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SendCDNVideoMsg 发送 CDN 视频消息。
// 端点: POST /wxwork/SendCDNVideoMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - cdnKey: CDN 文件标识
//   - aesKey: AES 解密密钥
//   - md5: 文件 MD5 校验值
//   - videoDuration: 视频时长（秒）
//   - fileSize: 文件大小（字节）
//   - videoImgSize: 视频封面大小（字节）
//   - videoWidth: 视频宽度（像素）
//   - videoHeight: 视频高度（像素）
//
// 返回: *SendCDNVideoResp 包含消息 ID 等发送详情
func (s *SdkIPad) SendCDNVideoMsg(host, uuid string, sendUserId int64, isRoom bool, cdnKey, aesKey, md5 string, videoDuration int, fileSize, videoImgSize int64, videoWidth, videoHeight int) (*SendCDNVideoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":           uuid,
		"send_userid":    sendUserId,
		"cdnkey":         cdnKey,
		"aeskey":         aesKey,
		"md5":            md5,
		"video_duration": videoDuration,
		"fileSize":       fileSize,
		"video_img_size": videoImgSize,
		"video_width":    videoWidth,
		"video_height":   videoHeight,
		"isRoom":         isRoom,
	}
	var resp SendCDNVideoResp
	if err := s.httpClient.Post(host+"/wxwork/SendCDNVideoMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
