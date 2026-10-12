// api_contact_get_managed_conversations.go GetManagedConversations 获取置顶/免打扰/折叠/标记列表。

package sdk_ipad

import "errors"

// ConversationItem 会话管理列表项（置顶/免打扰/标记）。
type ConversationItem struct {
	Messagetype int   `json:"messagetype"`
	Id          int64 `json:"id"`
	FwId        int64 `json:"fw_id"`
	State       int   `json:"state"`
}

// ManagedConversationsResp 获取置顶/免打扰/折叠/标记列表的返回结果。
type ManagedConversationsResp struct {
	Data struct {
		TopList         []ConversationItem `json:"top_list"`
		ShieldList      []ConversationItem `json:"shield_list"`
		MarkSessionList []ConversationItem `json:"mark_session_list"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetManagedConversations 获取置顶/免打扰/折叠/标记列表。
// 端点: POST /wxwork/getManagedConversations
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *ManagedConversationsResp 包含置顶、免打扰和标记列表
func (s *SdkIPad) GetManagedConversations(host, uuid string) (*ManagedConversationsResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]string{
		"uuid": uuid,
	}

	var resp ManagedConversationsResp
	if err := s.httpClient.Post(host+"/wxwork/getManagedConversations", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
