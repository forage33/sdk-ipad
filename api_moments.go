// api_moments.go 朋友圈操作相关接口。
// 提供朋友圈列表获取、详情查看、发送朋友圈、删除朋友圈、
// 评论、删除评论和点赞/取消点赞等功能。

package sdk_ipad

import "errors"

// ---------------------------------------------------------------------------
// 数据结构
// ---------------------------------------------------------------------------

// SnsMedia 朋友圈媒体文件信息。
type SnsMedia struct {
	Size      int    `json:"size"`
	Width     int    `json:"width"`
	VideoLen  int    `json:"video_len"`
	StoreType int    `json:"store_type"`
	Type      int    `json:"type"`   // 1=图片, 2=视频, 3=视频封面
	FileID    string `json:"fileid"` // CDN 文件 ID
	Md5       string `json:"md5"`
	Height    int    `json:"height"`
}

// SnsComment 朋友圈评论信息。
type SnsComment struct {
	FromID       int64  `json:"from_id"`       // 评论人 VID
	Flag         int    `json:"flag"`          // 标记
	CommentID    int64  `json:"commentid"`     // 评论 ID
	Comment      string `json:"comment"`       // 评论内容
	Time         int64  `json:"time"`          // 评论时间（Unix 时间戳）
	PostID       string `json:"postid"`        // 朋友圈 Post ID
	RefCommentID int64  `json:"ref_commentid"` // 引用评论 ID（0 表示直接评论）
	SID          int64  `json:"sid"`           // 朋友圈 SID
}

// SnsLinkInfo 朋友圈链接信息。
type SnsLinkInfo struct {
	Title      string `json:"title"`
	ContentUrl string `json:"content_url"`
}

// SnsLikeItem 朋友圈点赞信息。
type SnsLikeItem struct {
	Xid    int64  `json:"xid"`
	Time   int64  `json:"time"`
	PostID string `json:"postid"`
	SID    int64  `json:"sid"`
}

// SnsPoiInfo 朋友圈位置信息。
type SnsPoiInfo struct {
	City               string `json:"city"`
	Latitude           string `json:"latitude"`
	Longitude          string `json:"longitude"`
	PoiAddress         string `json:"poiAddress"`
	PoiClassifyId      string `json:"poiClassifyId"`
	PoiClassifyType    string `json:"poiClassifyType"`
	PoiClickableStatus string `json:"poiClickableStatus"`
	PoiName            string `json:"poiName"`
	PoiScale           string `json:"poiScale"`
}

// SnsItem 朋友圈条目详情。
type SnsItem struct {
	Flag              int           `json:"flag"`
	TaskSID           int64         `json:"task_sid"`
	PostID            string        `json:"postid"`
	Media             []SnsMedia    `json:"media"`
	Type              int           `json:"type"`
	Content           string        `json:"content"`
	SID               int64         `json:"sid"`
	UpdateTime        int64         `json:"update_time"`
	Xid               []int64       `json:"xid"`
	AuthorVid         int64         `json:"author_vid"`
	DeleteCommentList []SnsComment  `json:"delete_comment_list"`
	CommentList       []SnsComment  `json:"comment_list"`
	LinkInfo          SnsLinkInfo   `json:"link_info"`
	VisibleType       int           `json:"visible_type"`
	Time              int64         `json:"time"`
	Seq               int64         `json:"seq"`
	SnsLikeInfo       []SnsLikeItem `json:"SnsLikeInfo"`
	PoiInfo           *SnsPoiInfo   `json:"poiInfo,omitempty"`
}

// ---------------------------------------------------------------------------
// 响应结构
// ---------------------------------------------------------------------------

// SnsListResp 获取朋友圈列表的返回结果。
type SnsListResp struct {
	Data    []SnsItem `json:"data"`
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
}

// SnsInfoResp 获取朋友圈详情的返回结果。
type SnsInfoResp struct {
	Data    SnsItem `json:"data"`
	Errcode int     `json:"errcode"`
	Errmsg  string  `json:"errmsg"`
}

// PostPyqResp 发送朋友圈的返回结果。
type PostPyqResp struct {
	Data    SnsItem `json:"data"`
	Errcode int     `json:"errcode"`
	Errmsg  string  `json:"errmsg"`
}

// CommentSnsResp 发送朋友圈评论的返回结果。
type CommentSnsResp struct {
	Data struct {
		CommentID int64 `json:"commentid"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// CommonSnsResp 朋友圈操作通用返回结果。
type CommonSnsResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 调用接口 - 读取
// ---------------------------------------------------------------------------

// GetSnsList 获取朋友圈列表。
// 端点: POST /wxwork/GetSnsList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - seq: 分页序列号，首次传 -1，后续传返回的最后一条 sid 进行翻页
//   - limit: 每页数量
//
// 返回: *SnsListResp 朋友圈列表
func (s *SdkIPad) GetSnsList(host, uuid string, seq int64, limit int) (*SnsListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"seq":   seq,
		"limit": limit,
	}
	var resp SnsListResp
	if err := s.httpClient.Post(host+"/wxwork/GetSnsList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSnsInfo 获取朋友圈详情。
// 端点: POST /wxwork/GetSnsInfo
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sid: 朋友圈 SID
//
// 返回: *SnsInfoResp 朋友圈详情
func (s *SdkIPad) GetSnsInfo(host, uuid string, sid int64) (*SnsInfoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid": uuid,
		"sid":  sid,
	}
	var resp SnsInfoResp
	if err := s.httpClient.Post(host+"/wxwork/GetSnsInfo", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ---------------------------------------------------------------------------
// 调用接口 - 写入
// ---------------------------------------------------------------------------

// PostPyq 发送朋友圈。
// 端点: POST /wxwork/PostPyq
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - content: 朋友圈文字内容
//   - fileContent: 媒体文件列表（图片 type=1, 视频 type=2 + 封面 type=3, 图片与视频不可混发）
//   - vids: 可见人 VID 列表，空数组或 nil 表示全部可见
//   - linkInfo: 链接信息，不发链接朋友圈时传 nil
//
// 返回: *PostPyqResp 包含发送后的朋友圈详情
func (s *SdkIPad) PostPyq(host, uuid, content string, fileContent []SnsMedia, vids []int64, linkInfo *SnsLinkInfo) (*PostPyqResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	// 将 SnsMedia 转换为请求所需的 map 格式，使用 videoLen 驼峰字段名
	var files []map[string]interface{}
	for _, m := range fileContent {
		item := map[string]interface{}{
			"type":   m.Type,
			"fileid": m.FileID,
			"md5":    m.Md5,
			"size":   m.Size,
			"width":  m.Width,
			"height": m.Height,
		}
		if m.VideoLen > 0 {
			item["videoLen"] = m.VideoLen
		}
		files = append(files, item)
	}
	if files == nil {
		files = []map[string]interface{}{}
	}

	if vids == nil {
		vids = []int64{}
	}

	reqParams := map[string]interface{}{
		"uuid":         uuid,
		"content":      content,
		"file_content": files,
		"vids":         vids,
	}

	if linkInfo != nil {
		reqParams["link_info"] = map[string]interface{}{
			"title":       linkInfo.Title,
			"content_url": linkInfo.ContentUrl,
		}
	}

	var resp PostPyqResp
	if err := s.httpClient.Post(host+"/wxwork/PostPyq", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

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
