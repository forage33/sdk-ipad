// api_message_send_location_msg.go SendLocationMsg 发送位置消息。

package sdk_ipad

import "errors"

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
