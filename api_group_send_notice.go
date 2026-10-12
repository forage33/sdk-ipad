// api_group_send_notice.go SendNotice 发送群公告。

package sdk_ipad

import "errors"

// SendNotice 发送群公告。
// 端点: POST /wxwork/SendNotice
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - roomID: 群 ID
//   - msg: 公告内容
//
// 返回: *CommonGroupResp 操作结果
func (s *SdkIPad) SendNotice(host, uuid string, roomID int64, msg string) (*CommonGroupResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"roomid": roomID,
		"msg":    msg,
	}
	var resp CommonGroupResp
	if err := s.httpClient.Post(host+"/wxwork/SendNotice", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
