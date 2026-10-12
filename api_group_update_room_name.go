// api_group_update_room_name.go UpdateRoomName 修改群名称。

package sdk_ipad

import "errors"

// UpdateRoomNameResp 修改群名的返回结果。
type UpdateRoomNameResp struct {
	Data struct {
		Roomname string `json:"roomname"`
		Roomid   string `json:"roomid"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// UpdateRoomName 修改群名称。
// 端点: POST /wxwork/UpdateRoomName
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//   - roomName: 新群名称
//
// 返回: *UpdateRoomNameResp 修改结果
func (s *SdkIPad) UpdateRoomName(host, uuid string, roomID int64, roomName string) (*UpdateRoomNameResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":     uuid,
		"roomid":   roomID,
		"roomname": roomName,
	}
	var resp UpdateRoomNameResp
	if err := s.httpClient.Post(host+"/wxwork/UpdateRoomName", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
