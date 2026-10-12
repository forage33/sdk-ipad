// api_open_platform_vid2_open_user_tripartite_id.go Vid2OpenUserTripartiteId 企业用户 id 转开放平台代开发应用版本 id。

package sdk_ipad

import "errors"

// Vid2OpenUserTripartiteIdResp 企业用户 id 转开放平台代开发应用版本 id 的返回结果。
type Vid2OpenUserTripartiteIdResp struct {
	Data    []VidOpenIdPair `json:"data"`
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
}

// Vid2OpenUserTripartiteId 企业用户 id 转开放平台代开发应用版本 id。
// 端点: POST /wxwork/vid2OpenUserTripartiteId
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - scorpID: 企业 id（scorpid/短标识）
//   - corpID: 企业 id（长整型版本）
//   - appID: 应用 id
//   - userIDs: 要转换的企业用户 id（VID）列表
//
// 返回: *Vid2OpenUserTripartiteIdResp 每项含 user_id 与对应 openid
func (s *SdkIPad) Vid2OpenUserTripartiteId(host, uuid, scorpID string, corpID, appID int64, userIDs []int64) (*Vid2OpenUserTripartiteIdResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"user_id": userIDs,
	}
	if scorpID != "" {
		reqParams["scorpid"] = scorpID
	}
	if corpID != 0 {
		reqParams["corpid"] = corpID
	}
	if appID != 0 {
		reqParams["appId"] = appID
	}
	var resp Vid2OpenUserTripartiteIdResp
	if err := s.httpClient.Post(host+"/wxwork/vid2OpenUserTripartiteId", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
