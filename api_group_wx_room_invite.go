// api_group_wx_room_invite.go WxRoomInvite 获取群二维码。

package sdk_ipad

import (
	"encoding/json"
	"errors"
	"strings"
)

// WxRoomInviteResp 获取群二维码的返回结果。
type WxRoomInviteResp struct {
	Data struct {
		QrCodePath string      `json:"QrCodePath"`
		ImageUrl   string      `json:"image_url"`
		Roomid     json.Number `json:"roomid"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// WxRoomInvite 获取群二维码。
// 端点: POST /wxwork/WxRoomInvite
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *WxRoomInviteResp 包含群二维码路径和图片地址
func (s *SdkIPad) WxRoomInvite(host, uuid string, roomID int64) (*WxRoomInviteResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp WxRoomInviteResp
	if err := s.httpClient.Post(host+"/wxwork/WxRoomInvite", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NormalizeQrCodePath 规范化群二维码图片 URL。
// 协议服务可能返回反斜杠分隔的 Windows 风格路径（如 http://host/download\RoomQrCode\xx.jpg），
// 统一替换为正斜杠，保证前端 <img src> 可正常加载；同时兼容 URL 编码的反斜杠
// （RFC 3986 百分号编码十六进制大小写不敏感，%5C 与 %5c 均需处理）。
func NormalizeQrCodePath(path string) string {
	path = strings.ReplaceAll(path, "%5C", "/")
	path = strings.ReplaceAll(path, "%5c", "/")
	return strings.ReplaceAll(path, "\\", "/")
}
