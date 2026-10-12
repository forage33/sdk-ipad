// api_message_send_video_number_zhi_bo.go SendVideoNumberZhiBo 发送视频号直播消息。

package sdk_ipad

import "errors"

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
