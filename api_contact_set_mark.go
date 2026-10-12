// api_contact_set_mark.go SetMark 设置或取消星标标记。

package sdk_ipad

import "errors"

// SetMark 设置或取消星标标记。
// 端点: POST /wxwork/SetMark
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 联系人或群 ID
//   - isRoom: 是否为群会话
//   - state: 1=标记, 0=取消标记
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) SetMark(host, uuid string, vid int64, isRoom bool, state int) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"vid":    vid,
		"isRoom": isRoom,
		"state":  state,
	}

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/SetMark", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
