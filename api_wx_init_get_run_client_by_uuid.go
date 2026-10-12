// api_wx_init_get_run_client_by_uuid.go GetRunClientByUuid 根据 UUID 获取单个运行中的实例详情。

package sdk_ipad

import "errors"

// ClientRespByUUID 根据 UUID 查询单个实例的返回结果。
type ClientRespByUUID struct {
	Data    UUIDInfo `json:"data"`
	Errcode int      `json:"errcode"`
	Errmsg  string   `json:"errmsg"`
}

// UUIDInfo UUID 对应实例的详细信息，包含登录时间、登录方式和用户详情。
type UUIDInfo struct {
	Logintime int `json:"logintime"`
	LoginType int `json:"loginType"`
	UserInfo  struct {
		Object struct {
			Unionid      string      `json:"unionid"`
			CreateTime   int         `json:"create_time"`
			Sex          int         `json:"sex"`
			Mobile       string      `json:"mobile"`
			Acctid       string      `json:"acctid"`
			ScorpId      string      `json:"scorp_id"`
			Avatar       string      `json:"avatar"`
			CorpName     string      `json:"corp_name"`
			EnglishName  string      `json:"english_name"`
			Realname     string      `json:"realname"`
			UserId       interface{} `json:"user_id"`
			Nickname     string      `json:"nickname"`
			Position     string      `json:"position"`
			CorpId       interface{} `json:"corp_id"`
			CorpFullName string      `json:"corp_full_name"`
			CorpDesc     string      `json:"corp_desc"`
		} `json:"object"`
		Userid     interface{} `json:"userid"`
		ClientId   string      `json:"clientId"`
		UpdateTime int         `json:"updateTime"`
		Uuid       string      `json:"uuid"`
		Requrl     string      `json:"requrl"`
		IsLogin    bool        `json:"isLogin"`
	} `json:"user_info"`
	Uuid string `json:"uuid"`
}

// GetRunClientByUuid 根据 UUID 获取单个运行中的实例详情。
// 端点: POST /wxwork/GetRunClientByUuid
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *ClientRespByUUID 实例详细信息
func (s *SdkIPad) GetRunClientByUuid(host, uuid string) (*ClientRespByUUID, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
	}

	var clientResp ClientRespByUUID
	if err := s.httpClient.Post(host+"/wxwork/GetRunClientByUuid", reqParams, &clientResp); err != nil {
		return nil, err

	}

	return &clientResp, nil
}
