// api_group_create_room_wx.go CreateRoomWx 创建外部群聊。

package sdk_ipad

import "errors"

// CreateRoomWx 创建外部群聊。
// 端点: POST /wxwork/CreateRoomWx
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomName: 群名称
//   - vids: 初始成员 VID 列表
//
// 返回: *CreateRoomResp 包含新群 ID 和创建信息
func (s *SdkIPad) CreateRoomWx(host, uuid, roomName string, vids []int64) (*CreateRoomResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":     uuid,
		"roomName": roomName,
		"vids":     vids,
	}
	var resp CreateRoomResp
	if err := s.httpClient.Post(host+"/wxwork/CreateRoomWx", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
