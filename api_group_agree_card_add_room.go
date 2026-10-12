// api_group_agree_card_add_room.go AgreeCardAddRoom 通过二维码进群。

package sdk_ipad

import "errors"

// AgreeCardAddRoomResp 二维码进群的返回结果。
type AgreeCardAddRoomResp struct {
	Data struct {
		Data struct {
			Vid    string `json:"vid"`
			Status string `json:"status"`
		} `json:"data"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AgreeCardAddRoom 通过二维码进群。
// 端点: POST /wxwork/AgreeCardAddRoom
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - code: 群二维码 code
//
// 返回: *AgreeCardAddRoomResp 进群结果
func (s *SdkIPad) AgreeCardAddRoom(host, uuid, code string) (*AgreeCardAddRoomResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid": uuid,
		"code": code,
	}
	var resp AgreeCardAddRoomResp
	if err := s.httpClient.Post(host+"/wxwork/AgreeCardAddRoom", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
