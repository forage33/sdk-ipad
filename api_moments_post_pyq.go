// api_moments_post_pyq.go PostPyq 发送朋友圈。

package sdk_ipad

import "errors"

// PostPyqResp 发送朋友圈的返回结果。
type PostPyqResp struct {
	Data    SnsItem `json:"data"`
	Errcode int     `json:"errcode"`
	Errmsg  string  `json:"errmsg"`
}

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
