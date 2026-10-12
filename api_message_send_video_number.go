// api_message_send_video_number.go SendVideoNumber 发送视频号消息。

package sdk_ipad

import "errors"

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
