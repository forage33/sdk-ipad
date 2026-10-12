// api_group_wx_room_header.go WxRoomHeader 获取群头像。

package sdk_ipad

import (
	"encoding/json"
	"errors"
)

// WxRoomHeaderResp 获取群头像的返回结果。
type WxRoomHeaderResp struct {
	Data struct {
		ImageUrl string      `json:"image_url"`
		Roomid   json.Number `json:"roomid"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// WxRoomHeader 获取群头像。
// 端点: POST /wxwork/WxRoomHeader
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *WxRoomHeaderResp 包含群头像图片地址
func (s *SdkIPad) WxRoomHeader(host, uuid string, roomID int64) (*WxRoomHeaderResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp WxRoomHeaderResp
	if err := s.httpClient.Post(host+"/wxwork/WxRoomHeader", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
