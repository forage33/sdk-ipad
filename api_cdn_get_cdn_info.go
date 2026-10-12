// api_cdn_get_cdn_info.go GetCdnInfo 获取账号的 CDN 初始化参数。

package sdk_ipad

import "errors"

// GetCdnInfoResp 获取 CDN 初始化参数的返回结果。
type GetCdnInfoResp struct {
	Data struct {
		CorpID  int64  `json:"corpid"`  // 企业 id
		Vid     int64  `json:"vid"`     // 当前账号 id
		CdnInfo string `json:"cdninfo"` // CDN 初始化参数
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetCdnInfo 获取账号的 CDN 初始化参数。
// 端点: POST /wxwork/GetCdnInfo
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *GetCdnInfoResp 含 corpid/vid/cdninfo
//
// 注意: 从账号服务获取，哪个账号操作 CDN 就取哪个账号的，不可混用；
// 获取结果可缓存（1-2 天更新一次）。
func (s *SdkIPad) GetCdnInfo(host, uuid string) (*GetCdnInfoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
	}
	var resp GetCdnInfoResp
	if err := s.httpClient.Post(host+"/wxwork/GetCdnInfo", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
