// api_group_agree_add_room.go AgreeAddRoom 同意群邀请链接进群。

package sdk_ipad

import "errors"

// AgreeAddRoomResp 同意进群的返回结果。
type AgreeAddRoomResp struct {
	Data struct {
		Data struct {
			Vid    string `json:"vid"`
			Status string `json:"status"`
		} `json:"data"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AgreeAddRoom 同意群邀请链接进群。
// 端点: POST /wxwork/AgreeAddRoom
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - url: 群邀请链接
//
// 返回: *AgreeAddRoomResp 进群结果
func (s *SdkIPad) AgreeAddRoom(host, uuid, url string) (*AgreeAddRoomResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid": uuid,
		"url":  url,
	}
	var resp AgreeAddRoomResp
	if err := s.httpClient.Post(host+"/wxwork/AgreeAddRoom", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
