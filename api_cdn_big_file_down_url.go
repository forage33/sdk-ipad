// api_cdn_big_file_down_url.go BigFileDownUrl 下载大文件（由服务端下载并返回本地地址）。

package sdk_ipad

import "errors"

// BigFileDownUrlResp 大文件下载的返回结果。
type BigFileDownUrlResp struct {
	Data struct {
		Url string `json:"url"`
		Md5 string `json:"md5"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// BigFileDownUrl 下载大文件（由服务端下载并返回本地地址）。
// 端点: POST /wxwork/bigFileDownUrl
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - fileId: 文件 ID
//   - fileName: 文件名
//   - size: 文件大小（字节）
//
// 返回: *BigFileDownUrlResp 包含下载 URL 和 MD5
func (s *SdkIPad) BigFileDownUrl(host, uuid, fileId, fileName string, size int64) (*BigFileDownUrlResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":      uuid,
		"fileid":    fileId,
		"file_name": fileName,
		"size":      size,
	}
	var resp BigFileDownUrlResp
	if err := s.httpClient.Post(host+"/wxwork/bigFileDownUrl", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
