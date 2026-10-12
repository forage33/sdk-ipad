// api_moments_delete_sns.go DeleteSns 删除朋友圈。

package sdk_ipad

import "errors"

// DeleteSns 删除朋友圈。
// 端点: POST /wxwork/DeleteSns
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sid: 朋友圈 SID
//
// 返回: *CommonSnsResp 操作结果
func (s *SdkIPad) DeleteSns(host, uuid string, sid int64) (*CommonSnsResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid": uuid,
		"sid":  sid,
	}
	var resp CommonSnsResp
	if err := s.httpClient.Post(host+"/wxwork/DeleteSns", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
