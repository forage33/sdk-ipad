// api_moments_like_or_un_like_sns.go LikeOrUnLikeSns 朋友圈点赞或取消点赞。

package sdk_ipad

import "errors"

// LikeOrUnLikeSns 朋友圈点赞或取消点赞。
// 端点: POST /wxwork/LikeOrUnLikeSns
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sid: 朋友圈 SID
//   - operation: 0=点赞, 1=取消点赞
//
// 返回: *CommonSnsResp 操作结果
func (s *SdkIPad) LikeOrUnLikeSns(host, uuid string, sid int64, operation int) (*CommonSnsResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":      uuid,
		"sid":       sid,
		"operation": operation,
	}
	var resp CommonSnsResp
	if err := s.httpClient.Post(host+"/wxwork/LikeOrUnLikeSns", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
