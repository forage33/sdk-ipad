// api_cdn_download_file.go DownloadFile 下载 CDN 文件（图片、视频、缩略图、文件、语音等）。

package sdk_ipad

import "errors"

// DownloadFileResp CDN 下载文件的返回结果。
type DownloadFileResp struct {
	Data    string `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// DownloadFile 下载 CDN 文件（图片、视频、缩略图、文件、语音等）。
// 端点: POST /wxwork/DownloadFile
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - fileId: 文件 ID
//   - aesKey: AES 解密密钥
//   - fileName: 文件名
//   - fileType: 文件类型（1=图片, 2=高清图, 3=缩略图, 4=视频, 5=文件/语音）
//   - size: 文件大小（字节）
//
// 返回: *DownloadFileResp 包含下载地址
func (s *SdkIPad) DownloadFile(host, uuid, fileId, aesKey, fileName string, fileType int, size int64) (*DownloadFileResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":      uuid,
		"fileid":    fileId,
		"aes_key":   aesKey,
		"filetype":  fileType,
		"file_name": fileName,
		"size":      size,
	}
	var resp DownloadFileResp
	if err := s.httpClient.Post(host+"/wxwork/DownloadFile", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
