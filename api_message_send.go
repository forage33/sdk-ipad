// api_message_send.go 消息发送相关接口。
// 提供文本、图片、文件、链接、语音、视频、表情、名片、位置、@消息、
// 小程序、视频号、群发等消息发送功能（撤回/已读/语音转文字/同步见 api_message_action.go）。

package sdk_ipad

import (
	"errors"
)

// ---------------------------------------------------------------------------
// 返回数据结构 - 基础消息
// ---------------------------------------------------------------------------

// SendTextResp 发送文本消息的返回结果。
type SendTextResp struct {
	Data struct {
		Receiver   int64         `json:"receiver"`
		Sender     int64         `json:"sender"`
		AtList     []interface{} `json:"at_list"`
		SenderName string        `json:"sender_name"`
		AppInfo    string        `json:"app_info"`
		IsRoom     int           `json:"is_room"`
		Sendtime   int           `json:"sendtime"`
		MsgId      int           `json:"msg_id"`
		ServerId   int           `json:"server_id"`
		Msgtype    int           `json:"msgtype"`
		Content    string        `json:"content"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SendCDNImageResp 发送 CDN 图片消息的返回结果。
type SendCDNImageResp struct {
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

// ---------------------------------------------------------------------------
// 返回数据结构 - CDN 媒体类消息
// ---------------------------------------------------------------------------

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

// SendCDNBigFileResp 发送大文件消息的返回结果。
type SendCDNBigFileResp struct {
	Data struct {
		Receiver   int64  `json:"receiver"`
		CdnKey     string `json:"cdn_key"`
		FileName   string `json:"file_name"`
		SenderName string `json:"sender_name"`
		IsRoom     int    `json:"is_room"`
		ServerId   int    `json:"server_id"`
		Size       int    `json:"size"`
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

// ---------------------------------------------------------------------------
// 返回数据结构 - 应用类消息
// ---------------------------------------------------------------------------

// SendAppMsgResp 发送小程序消息的返回结果。
type SendAppMsgResp struct {
	Data    map[string]interface{} `json:"data"`
	Errcode int                    `json:"errcode"`
	Errmsg  string                 `json:"errmsg"`
}

// SendVideoNumberResp 发送视频号消息的返回结果。
type SendVideoNumberResp struct {
	Data struct {
		ThumbUrl   string `json:"thumb_url"`
		CoverUrl   string `json:"cover_url"`
		Receiver   int64  `json:"receiver"`
		Extras     string `json:"extras"`
		SenderName string `json:"sender_name"`
		AppInfo    string `json:"app_info"`
		IsRoom     int    `json:"is_room"`
		Avatar     string `json:"avatar"`
		ServerId   int    `json:"server_id"`
		Url        string `json:"url"`
		Sender     int64  `json:"sender"`
		Nickname   string `json:"nickname"`
		Sendtime   int    `json:"sendtime"`
		MsgId      int    `json:"msg_id"`
		Msgtype    int    `json:"msgtype"`
		Desc       string `json:"desc"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SendVideoNumberZhiBoResp 发送视频号直播消息的返回结果。
type SendVideoNumberZhiBoResp struct {
	Data struct {
		ThumbUrl      string `json:"thumb_url"`
		CoverUrl      string `json:"cover_url"`
		ObjectNonceId string `json:"objectNonceId"`
		Receiver      int64  `json:"receiver"`
		Extras        string `json:"extras"`
		SenderName    string `json:"sender_name"`
		IsRoom        int    `json:"is_room"`
		Avatar        string `json:"avatar"`
		ServerId      int    `json:"server_id"`
		Url           string `json:"url"`
		CoverUrlField string `json:"coverUrl"`
		Sender        int64  `json:"sender"`
		Nickname      string `json:"nickname"`
		AppInfo       string `json:"app_info"`
		Sendtime      int    `json:"sendtime"`
		MsgId         int    `json:"msg_id"`
		ThumbUrlField string `json:"thumbUrl"`
		Msgtype       int    `json:"msgtype"`
		FeedType      int    `json:"feed_type"`
		ObjectId      string `json:"objectId"`
		Desc          string `json:"desc"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 返回数据结构 - 互动类消息
// ---------------------------------------------------------------------------

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

// SendLocationResp 发送位置消息的返回结果。
type SendLocationResp struct {
	Data struct {
		DetailedAddress string  `json:"detailed_address"`
		Address         string  `json:"address"`
		Receiver        int64   `json:"receiver"`
		Sender          int64   `json:"sender"`
		Latitude        float64 `json:"latitude"`
		SenderName      string  `json:"sender_name"`
		AppInfo         string  `json:"app_info"`
		IsRoom          int     `json:"is_room"`
		Sendtime        int     `json:"sendtime"`
		MsgId           int     `json:"msg_id"`
		ServerId        int     `json:"server_id"`
		Msgtype         int     `json:"msgtype"`
		Longitude       float64 `json:"longitude"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SendTextAtResp 发送群 @消息的返回结果。
type SendTextAtResp struct {
	Data struct {
		Receiver int64 `json:"receiver"`
		Sender   int64 `json:"sender"`
		AtList   []struct {
			UserId   int64  `json:"user_id"`
			NikeName string `json:"nikeName"`
		} `json:"at_list"`
		SenderName         string `json:"sender_name"`
		AppInfo            string `json:"app_info"`
		RoomConversationId string `json:"room_conversation_id"`
		IsRoom             int    `json:"is_room"`
		Sendtime           int    `json:"sendtime"`
		MsgId              int    `json:"msg_id"`
		ServerId           int    `json:"server_id"`
		Msgtype            int    `json:"msgtype"`
		Content            string `json:"content"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 返回数据结构 - 群发
// ---------------------------------------------------------------------------

// SendGroupsMsgResp 群发消息的返回结果。
type SendGroupsMsgResp struct {
	Data struct {
		MsgId int `json:"msg_id"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 调用接口 - 基础消息
// ---------------------------------------------------------------------------

// SendTextMsg 发送文本消息。
// 端点: POST /wxwork/SendTextMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserID: 接收者 ID（好友或群 ID）
//   - content: 文本内容
//   - isRoom: 是否为群消息
//
// 返回: *SendTextResp 包含消息 ID 等发送详情
func (s *SdkIPad) SendTextMsg(host, uuid string, sendUserID int64, content string, isRoom bool) (*SendTextResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"send_userid": sendUserID,
		"isRoom":      isRoom,
		"content":     content,
	}
	var resp SendTextResp
	if err := s.httpClient.Post(host+"/wxwork/SendTextMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendCDNImage 发送 CDN 图片消息。
// 端点: POST /wxwork/SendCDNImgMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - cdnKey: CDN 文件标识
//   - aesKey: AES 解密密钥
//   - md5: 文件 MD5 校验值
//   - fileSize: 文件大小（字节）
//
// 返回: *SendCDNImageResp 包含消息 ID 等发送详情
func (s *SdkIPad) SendCDNImage(host, uuid string, sendUserId int64, isRoom bool, cdnKey, aesKey, md5 string, fileSize int64) (*SendCDNImageResp, error) {
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
		"fileSize":    fileSize,
		"isRoom":      isRoom,
	}

	var resp SendCDNImageResp
	if err := s.httpClient.Post(host+"/wxwork/SendCDNImgMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// ---------------------------------------------------------------------------
// 调用接口 - CDN 媒体类消息
// ---------------------------------------------------------------------------

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

// SendCDNBigFileMsg 发送大文件消息（需走大文件上传流程）。
// 端点: POST /wxwork/SendCDNBigFileMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - cdnKey: CDN 文件标识
//   - md5: 文件 MD5 校验值
//   - fileName: 文件名
//   - fileSize: 文件大小（字节）
//
// 返回: *SendCDNBigFileResp 包含消息 ID 等发送详情
func (s *SdkIPad) SendCDNBigFileMsg(host, uuid string, sendUserId int64, isRoom bool, cdnKey, md5, fileName string, fileSize int64) (*SendCDNBigFileResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"send_userid": sendUserId,
		"cdnkey":      cdnKey,
		"md5":         md5,
		"file_name":   fileName,
		"fileSize":    fileSize,
		"isRoom":      isRoom,
	}
	var resp SendCDNBigFileResp
	if err := s.httpClient.Post(host+"/wxwork/SendCDNBigFileMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ---------------------------------------------------------------------------
// 调用接口 - 应用类消息
// ---------------------------------------------------------------------------

// SendAppMsg 发送小程序消息。
// 端点: POST /wxwork/SendAppMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - desc: 小程序描述
//   - appName: 小程序名称
//   - title: 消息标题
//   - weappIconUrl: 小程序图标地址
//   - pagepath: 小程序页面路径
//   - username: 小程序原始 ID
//   - appid: 小程序 AppID
//   - cdnKey: 封面 CDN 文件标识
//   - md5: 封面 MD5
//   - aesKey: 封面 AES 密钥
//   - fileSize: 封面文件大小（字节）
//
// 返回: *SendAppMsgResp 发送结果
func (s *SdkIPad) SendAppMsg(host, uuid string, sendUserId int64, isRoom bool, desc, appName, title, weappIconUrl, pagepath, username, appid, cdnKey, md5, aesKey string, fileSize int64) (*SendAppMsgResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":         uuid,
		"send_userid":  sendUserId,
		"desc":         desc,
		"appName":      appName,
		"title":        title,
		"weappIconUrl": weappIconUrl,
		"pagepath":     pagepath,
		"username":     username,
		"appid":        appid,
		"cdnkey":       cdnKey,
		"md5":          md5,
		"aeskey":       aesKey,
		"fileSize":     fileSize,
		"isRoom":       isRoom,
	}
	var resp SendAppMsgResp
	if err := s.httpClient.Post(host+"/wxwork/SendAppMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendVideoNumber 发送视频号消息。
// 端点: POST /wxwork/SendVideoNumber
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - coverUrl: 视频封面地址
//   - thumbUrl: 缩略图地址
//   - avatar: 发布者头像
//   - nickname: 发布者昵称
//   - desc: 视频描述
//   - url: 视频播放地址
//   - extras: 额外数据
//   - objectId: 视频号对象 ID
//   - objectNonceId: 视频号对象随机 ID
//
// 返回: *SendVideoNumberResp 发送结果
func (s *SdkIPad) SendVideoNumber(host, uuid string, sendUserId int64, isRoom bool, coverUrl, thumbUrl, avatar, nickname, desc, url, extras, objectId, objectNonceId string) (*SendVideoNumberResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":          uuid,
		"send_userid":   sendUserId,
		"isRoom":        isRoom,
		"coverUrl":      coverUrl,
		"thumbUrl":      thumbUrl,
		"avatar":        avatar,
		"nickname":      nickname,
		"desc":          desc,
		"url":           url,
		"extras":        extras,
		"objectId":      objectId,
		"objectNonceId": objectNonceId,
	}
	var resp SendVideoNumberResp
	if err := s.httpClient.Post(host+"/wxwork/SendVideoNumber", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendVideoNumberZhiBo 发送视频号直播消息。
// 端点: POST /wxwork/SendVideoNumberZhiBo
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - coverUrl: 直播封面地址
//   - thumbUrl: 缩略图地址
//   - avatar: 主播头像
//   - nickname: 主播昵称
//   - desc: 直播描述
//   - url: 直播地址
//   - extras: 额外数据
//   - objectId: 视频号对象 ID
//   - objectNonceId: 视频号对象随机 ID
//
// 返回: *SendVideoNumberZhiBoResp 发送结果
func (s *SdkIPad) SendVideoNumberZhiBo(host, uuid string, sendUserId int64, isRoom bool, coverUrl, thumbUrl, avatar, nickname, desc, url, extras, objectId, objectNonceId string) (*SendVideoNumberZhiBoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":          uuid,
		"send_userid":   sendUserId,
		"isRoom":        isRoom,
		"coverUrl":      coverUrl,
		"thumbUrl":      thumbUrl,
		"avatar":        avatar,
		"nickname":      nickname,
		"desc":          desc,
		"url":           url,
		"extras":        extras,
		"objectId":      objectId,
		"objectNonceId": objectNonceId,
	}
	var resp SendVideoNumberZhiBoResp
	if err := s.httpClient.Post(host+"/wxwork/SendVideoNumberZhiBo", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ---------------------------------------------------------------------------
// 调用接口 - 互动类消息
// ---------------------------------------------------------------------------

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

// SendLocationMsg 发送位置消息。
// 端点: POST /wxwork/SendLocationMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - longitude: 经度
//   - latitude: 纬度
//   - address: 地址名称
//   - detailedAddress: 详细地址
//
// 返回: *SendLocationResp 发送结果
func (s *SdkIPad) SendLocationMsg(host, uuid string, sendUserId int64, isRoom bool, longitude, latitude float64, address, detailedAddress string) (*SendLocationResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":             uuid,
		"send_userid":      sendUserId,
		"isRoom":           isRoom,
		"longitude":        longitude,
		"latitude":         latitude,
		"address":          address,
		"detailed_address": detailedAddress,
	}
	var resp SendLocationResp
	if err := s.httpClient.Post(host+"/wxwork/SendLocationMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendTextAtMsg 发送群 @消息。
// 端点: POST /wxwork/SendTextAtMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 群 ID
//   - atIds: 要 @的用户 ID 列表
//   - content: 消息文本内容
//
// 返回: *SendTextAtResp 发送结果
//
// 注意: 此接口 isRoom 固定为 true，仅用于群聊场景
func (s *SdkIPad) SendTextAtMsg(host, uuid string, sendUserId int64, atIds []int64, content string) (*SendTextAtResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":        uuid,
		"send_userid": sendUserId,
		"atids":       atIds,
		"content":     content,
		"isRoom":      true,
	}
	var resp SendTextAtResp
	if err := s.httpClient.Post(host+"/wxwork/SendTextAtMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ---------------------------------------------------------------------------
// 调用接口 - 群发
// ---------------------------------------------------------------------------

// SendGroupsMsg 群发消息（支持多人/多群同时发送）。
// 端点: POST /wxwork/SendGroupsMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vids: 接收者 VID 列表
//   - isRoom: 是否为群消息
//   - msgList: 消息内容列表，每个元素为一条消息的参数 map
//
// 返回: *SendGroupsMsgResp 包含群发消息 ID
func (s *SdkIPad) SendGroupsMsg(host, uuid string, vids []string, isRoom bool, msgList []map[string]interface{}) (*SendGroupsMsgResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":     uuid,
		"vids":     vids,
		"isroom":   isRoom,
		"msg_list": msgList,
	}
	var resp SendGroupsMsgResp
	if err := s.httpClient.Post(host+"/wxwork/SendGroupsMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
