// broadcast.go 控制台群发操作相关接口。
// 提供群发消息确认等功能。

package sdk_ipad

import "errors"

// ---------------------------------------------------------------------------
// 返回数据结构
// ---------------------------------------------------------------------------

// ConfirmGroupSendResp 确认群发消息的返回结果。
type ConfirmGroupSendResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 调用接口
// ---------------------------------------------------------------------------

// ConfirmGroupSend 确认群发消息（控制台下发后需客户端确认执行）。
// 端点: POST /wxwork/oksendgroup
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vids: 接收者 VID 列表
//   - msgID: 群发消息 ID
//   - isRoom: 是否群发到群（0=个人, 1=群）
//
// 返回: *ConfirmGroupSendResp 操作结果
func (s *SdkIPad) ConfirmGroupSend(host, uuid string, vids []int64, msgID int64, isRoom int) (*ConfirmGroupSendResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"vids":   vids,
		"msgid":  msgID,
		"isroom": isRoom,
	}
	var resp ConfirmGroupSendResp
	if err := s.httpClient.Post(host+"/wxwork/oksendgroup", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ---------------------------------------------------------------------------
// 待补全接口（参数待从文档确认后完善）
// ---------------------------------------------------------------------------

// TODO: GetPendingGroupSendList 获取未发送群发列表
// 预估路径: /wxwork/GetPendingGroupSendList
// 预估参数: uuid

// TODO: GetGroupSendCustomerList 获取群发客户群列表
// 预估路径: /wxwork/GetGroupSendCustomerList
// 预估参数: uuid, msgid(int64)

// TODO: GetGroupSendTaskDetail 获取群发任务详细信息
// 预估路径: /wxwork/GetGroupSendTaskDetail
// 预估参数: uuid, msgid(int64)

// TODO: ExecSnsConsoleTask 朋友圈控制台任务执行
// 预估路径: /wxwork/ExecSnsConsoleTask
// 预估参数: uuid, task_id(int64)
