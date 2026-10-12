// api_moments_comment_sns.go CommentSns 发送朋友圈评论。

package sdk_ipad

import "errors"

// CommentSnsResp 发送朋友圈评论的返回结果。
type CommentSnsResp struct {
	Data struct {
		CommentID int64 `json:"commentid"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// CommentSns 发送朋友圈评论。
// 端点: POST /wxwork/CommentSnsReq
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sid: 朋友圈 SID
//   - comment: 评论内容
//   - refCommentID: 引用评论 ID，传 0 为直接评论，非 0 为回复指定评论
//
// 返回: *CommentSnsResp 包含新评论 ID
func (s *SdkIPad) CommentSns(host, uuid string, sid int64, comment string, refCommentID int64) (*CommentSnsResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":          uuid,
		"sid":           sid,
		"comment":       comment,
		"ref_commentid": refCommentID,
	}
	var resp CommentSnsResp
	if err := s.httpClient.Post(host+"/wxwork/CommentSnsReq", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
