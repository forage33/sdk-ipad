// api_open_platform_user_id2_vid.go UserId2Vid 开放平台 id 转企业用户 id。

package sdk_ipad

import "errors"

// UserId2VidResp 开放平台 id 转企业用户 id 的返回结果。
type UserId2VidResp struct {
	Data    []VidOpenIdPair `json:"data"`
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
}

// UserId2Vid 开放平台 id（external_userid）转企业用户 id。
// 端点: POST /wxwork/UserId2Vid
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - scorpID: 企业 id（scorpId/短标识）
//   - corpID: 企业 id（长整型）
//   - appID: 应用 id（控制台分组下的应用列表获取）
//   - openIDs: 要转换的开放平台 id 列表
//
// 返回: *UserId2VidResp 每项含 user_id 与对应 openid
func (s *SdkIPad) UserId2Vid(host, uuid, scorpID string, corpID, appID int64, openIDs []string) (*UserId2VidResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"openid": openIDs,
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
	var resp UserId2VidResp
	if err := s.httpClient.Post(host+"/wxwork/UserId2Vid", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
