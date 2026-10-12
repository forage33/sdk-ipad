// api_message_send_groups_msg.go SendGroupsMsg 群发消息（支持多人/多群同时发送）。

package sdk_ipad

import "errors"

// SendGroupsMsgResp 群发消息的返回结果。
type SendGroupsMsgResp struct {
	Data struct {
		MsgId int `json:"msg_id"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SendGroupsMsg 群发消息（支持多人/多群同时发送）。
// 端点: POST /wxwork/SendGroupsMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vids: 接收者 VID 列表
//   - isRoom: 是否为群消息
//   - msgList: 消息内容列表，每个元素为一条消息的参数 map
//
// 返回: *SendGroupsMsgResp 包含群发消息 ID
func (s *SdkIPad) SendGroupsMsg(host, uuid string, vids []string, isRoom bool, msgList []map[string]interface{}) (*SendGroupsMsgResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":     uuid,
		"vids":     vids,
		"isroom":   isRoom,
		"msg_list": msgList,
	}
	var resp SendGroupsMsgResp
	if err := s.httpClient.Post(host+"/wxwork/SendGroupsMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
