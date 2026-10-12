// api_moments_delete_sns_comment.go DeleteSnsComment 删除朋友圈评论。

package sdk_ipad

import "errors"

// DeleteSnsComment 删除朋友圈评论。
// 端点: POST /wxwork/DeleteSnsComment
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sid: 朋友圈 SID
//   - commentID: 要删除的评论 ID
//
// 返回: *CommonSnsResp 操作结果
func (s *SdkIPad) DeleteSnsComment(host, uuid string, sid int64, commentID int64) (*CommonSnsResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":      uuid,
		"sid":       sid,
		"commentid": commentID,
	}
	var resp CommonSnsResp
	if err := s.httpClient.Post(host+"/wxwork/DeleteSnsComment", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
