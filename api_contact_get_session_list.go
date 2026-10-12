// api_contact_get_session_list.go GetSessionList 获取会话列表。

package sdk_ipad

import "errors"

// SessionItem 会话列表中的会话项。
type SessionItem struct {
	Unreadcnt   int   `json:"unreadcnt"`
	Sessionid   int64 `json:"sessionid"`
	Beginmsgseq int64 `json:"beginmsgseq"`
	Msgtype     int   `json:"msgtype"`
}

// SessionSimpleItem 会话列表中的简单项（置顶/免打扰）。
type SessionSimpleItem struct {
	Messagetype int   `json:"messagetype"`
	Id          int64 `json:"id"`
}

// SessionListResp 获取会话列表的返回结果。
type SessionListResp struct {
	Data struct {
		TopList     []SessionSimpleItem `json:"top_list"`
		CollectList []interface{}       `json:"collect_list"`
		ShieldList  []SessionSimpleItem `json:"shield_list"`
		Seq         int                 `json:"seq"`
		RoomList    []SessionItem       `json:"room_list"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetSessionList 获取会话列表。
// 端点: POST /wxwork/GetSessionList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - limit: 每页数量
//   - starIndex: 起始索引
//
// 返回: *SessionListResp 会话列表
func (s *SdkIPad) GetSessionList(host, uuid string, limit, starIndex int) (*SessionListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":       uuid,
		"limit":      limit,
		"star_index": starIndex,
	}

	var resp SessionListResp
	if err := s.httpClient.Post(host+"/wxwork/GetSessionList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
