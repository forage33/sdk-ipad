// api_wx_init.go 初始化相关接口。
// 提供企业微信 iPad 协议的实例初始化、回调设置、运行状态查询和连接关闭等功能。

package sdk_ipad

import (
	"errors"
)

// ---------------------------------------------------------------------------
// 返回数据结构
// ---------------------------------------------------------------------------

// InitResp 初始化接口返回结果。
type InitResp struct {
	ErrCode int      `json:"errcode"`
	ErrMsg  string   `json:"errmsg"`
	Data    InitData `json:"data"`
}

// InitData 初始化返回的核心数据，包含实例 UUID 和登录状态。
type InitData struct {
	UUID    string `json:"uuid"`
	IsLogin string `json:"is_login"`
}

// 回调类型常量
const (
	// CallbackTypeHTTP HTTP 回调
	CallbackTypeHTTP = "HTTP"
	// CallbackTypeRabbitMQ RABBITMQ 回调
	CallbackTypeRabbitMQ = "RABBITMQ"
)

// CallbackConfig 设置回调地址的请求参数。
type CallbackConfig struct {
	// CallbackType 消息回调类型: HTTP 或 RABBITMQ
	CallbackType string
	// URL HTTP 回调地址
	URL string
	// MqExchange RABBITMQ 交换机（direct 类型）
	MqExchange string
	// MqRoutingKey RABBITMQ 路由键
	MqRoutingKey string
	// ExtraContent 自定义回调额外内容，可为空
	ExtraContent string
}

// CallbackData 设置回调地址返回的数据部分。
type CallbackData struct {
	CallbackType string `json:"callbackType"`
	URL          string `json:"url"`
	MqExchange   string `json:"mqExchange"`
	MqRoutingKey string `json:"mqRoutingKey"`
	ExtraContent string `json:"extraContent"`
	Msg          string `json:"msg"`
	StatusCode   int    `json:"statusCode"`
}

// CallbackResp 设置回调地址的返回结果。
type CallbackResp struct {
	Data    CallbackData `json:"data"`
	ErrCode int          `json:"errcode"`
	ErrMsg  string       `json:"errmsg"`
}

// ClientResponse 获取运行中实例列表的返回结果。
type ClientResponse struct {
	Data    []map[string]UserData `json:"data"`
	ErrCode int                   `json:"errcode"`
	ErrMsg  string                `json:"errmsg"`
}

// UserData 运行中实例的用户数据。
type UserData struct {
	IsLogin    bool   `json:"isLogin"`
	ClientID   string `json:"clientId"`
	UpdateTime int64  `json:"updateTime"`
	UserID     int64  `json:"userid"`
	UUID       string `json:"uuid"`
	ReqURL     string `json:"requrl"`
	Object     User   `json:"object"`
}

// User 企业微信用户基本信息。
type User struct {
	UnionID      string `json:"unionid"`
	CreateTime   int64  `json:"create_time"`
	Sex          int    `json:"sex"`
	Mobile       string `json:"mobile"`
	CorpFullName string `json:"Corp_full_name"`
	AcctID       string `json:"acctid"`
	SCorpID      string `json:"scorp_id"`
	Avatar       string `json:"avatar"`
	CorpName     string `json:"corp_name"`
	EnglishName  string `json:"english_name"`
	RealName     string `json:"realname"`
	UserID       int64  `json:"user_id"`
	Nickname     string `json:"nickname"`
	Position     string `json:"position"`
	CorpDesc     string `json:"Corp_desc"`
	CorpID       int64  `json:"corp_id"`
}

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

// CloseConnResp 关闭连接的返回结果。
type CloseConnResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 调用接口
// ---------------------------------------------------------------------------

// InItWXWork 初始化企业微信实例（首次启动）。
// 端点: POST /wxwork/init
//
// 参数:
//   - host: 协议服务地址
//   - vid: 企业微信 VID
//   - proxyType: 代理类型，空字符串表示不使用代理，"http" 表示 HTTP 代理，"socks5" 表示 SOCKS5 代理
//   - ip: 代理 IP 地址（不使用代理时可为空）
//   - port: 代理端口（不使用代理时可为 0）
//   - userName: 代理用户名（无则为空）
//   - password: 代理密码（无则为空）
//   - deviceType: 设备类型（当前固定为 "ipad"）
//   - proxySituation: 代理场景标识
//
// 返回: *InitResp 包含实例 UUID 和登录状态
func (s *SdkIPad) InItWXWork(host, vid, proxyType string, ip string, port int, userName, password, deviceType string, proxySituation int) (*InitResp, error) {
	var reqParams map[string]interface{}
	// 请求参数
	switch proxyType {
	case "":
		// fmt.Println("==> not set proxy type")
		reqParams = map[string]interface{}{
			"vid":            vid, // vid
			"ip":             "",  // 代理ip
			"port":           0,   // 代理端口
			"proxyType":      "",  // http代理类型
			"userName":       "",  // 代理账号 没有不传
			"passward":       "",  // 代理密码 没有不传
			"deverType":      "ipad",
			"proxySituation": 0,
		}
	case "http":
		// fmt.Println("==> set http proxy type")
		reqParams = map[string]interface{}{
			"vid":            vid,       // vid
			"ip":             ip,        // 代理ip
			"port":           port,      // 代理端口
			"proxyType":      proxyType, // http代理类型
			"userName":       userName,  // 代理账号 没有不传
			"passward":       password,  // 代理密码 没有不传
			"deverType":      "ipad",
			"proxySituation": proxySituation,
		}
	case "socks5":
		reqParams = map[string]interface{}{
			"vid":            vid,
			"ip":             ip,
			"port":           port,
			"proxyType":      proxyType,
			"userName":       userName,
			"passward":       password,
			"deverType":      "ipad",
			"proxySituation": proxySituation,
		}
	}
	// fmt.Println("json msg:", reqParams)

	var initResp InitResp
	if err := s.httpClient.Post(host+"/wxwork/init", reqParams, &initResp); err != nil {
		return nil, err
	}

	return &initResp, nil
}

// SetCallBackURL 设置协议服务的回调推送地址。
// 端点: POST /wxwork/SetCallbackUrl
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - config: 回调配置，CallbackType 为空时默认 HTTP 回调
//
// 返回: *CallbackResp 设置结果
func (s *SdkIPad) SetCallBackURL(host, uuid string, config CallbackConfig) (*CallbackResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid": uuid,
	}
	switch config.CallbackType {
	case CallbackTypeRabbitMQ:
		reqParams["callbackType"] = CallbackTypeRabbitMQ
		reqParams["mqExchange"] = config.MqExchange
		reqParams["mqRoutingKey"] = config.MqRoutingKey
		if config.ExtraContent != "" {
			reqParams["extraContent"] = config.ExtraContent
		}
	default: // HTTP
		reqParams["callbackType"] = CallbackTypeHTTP
		reqParams["url"] = config.URL
	}

	var callbackResp CallbackResp

	if err := s.httpClient.Post(host+"/wxwork/SetCallbackUrl", reqParams, &callbackResp); err != nil {
		return nil, err
	}

	return &callbackResp, nil
}

// GetRunClient 获取所有运行中的实例列表。
// 端点: POST /wxwork/GetRunClient
//
// 参数:
//   - host: 协议服务地址
//
// 返回: *ClientResponse 运行中实例列表
func (s *SdkIPad) GetRunClient(host string) (*ClientResponse, error) {
	var clientResp ClientResponse
	if err := s.httpClient.Post(host+"/wxwork/GetRunClient", nil, &clientResp); err != nil {
		return nil, err
	}
	return &clientResp, nil
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

// CloseConnect 关闭指定实例的连接。
// 端点: POST /wxwork/CloseConnent
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *CloseConnResp 关闭结果
func (s *SdkIPad) CloseConnect(host, uuid string) (*CloseConnResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
	}

	var closeResp CloseConnResp
	if err := s.httpClient.Post(host+"/wxwork/CloseConnent", reqParams, &closeResp); err != nil {
		return nil, err
	}
	return &closeResp, nil
}

// ---------------------------------------------------------------------------
// 辅助方法
// ---------------------------------------------------------------------------

// IsRunningByVid 判断指定 VID 的账号是否正在运行中。
//
// 参数:
//   - host: 协议服务地址
//   - vid: 企业微信 VID
//
// 返回: bool 是否在运行中
func (s *SdkIPad) IsRunningByVid(host, vid string) bool {
	resp, err := s.GetRunClient(host)
	if err != nil {
		return false
	}
	for _, item := range resp.Data {
		for userID, userData := range item {
			if vid == userID && userData.IsLogin {
				return true
			}
		}
	}
	return false
}

// GetRunningByVid 根据 VID 查找运行中的实例，返回用户数据和是否找到。
//
// 参数:
//   - host: 协议服务地址
//   - vid: 企业微信 VID
//
// 返回:
//   - *UserData: 用户数据，未找到时为 nil
//   - bool: 是否找到
func (s *SdkIPad) GetRunningByVid(host, vid string) (*UserData, bool) {
	resp, err := s.GetRunClient(host)
	if err != nil {
		return nil, false
	}
	for _, item := range resp.Data {
		for userID, userData := range item {
			if vid == userID && userData.IsLogin {
				ud := userData
				return &ud, true
			}
		}
	}
	return nil, false
}
