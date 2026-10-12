// api_group_get_room_info.go GetRoomInfo 获取群详细信息和头像。

package sdk_ipad

import (
	"encoding/json"
	"errors"
)

// GetRoomInfoResp 获取群详细信息和头像的返回结果（data 为动态 JSON）。
type GetRoomInfoResp struct {
	Data    json.RawMessage `json:"data"`
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
}

// GetRoomInfo 获取群详细信息和头像。
// 端点: POST /wxwork/GetRoomInfo
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//
// 返回: *GetRoomInfoResp 群详情（Data 为 json.RawMessage，需自行解析）
func (s *SdkIPad) GetRoomInfo(host, uuid string, roomID int64) (*GetRoomInfoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
	}
	var resp GetRoomInfoResp
	if err := s.httpClient.Post(host+"/wxwork/GetRoomInfo", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
