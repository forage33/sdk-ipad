// api_wx_init_in_it_wx_work.go InItWXWork 初始化企业微信实例（首次启动）。

package sdk_ipad

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
