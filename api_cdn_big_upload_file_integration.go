// api_cdn_big_upload_file_integration.go BigUploadFileIntegration 通过 multipart/form-data 上传大文件。

package sdk_ipad

import (
	"errors"
	"io"
	"net/http"
	"time"
)

// DefaultBigUploadTimeout 大文件上传默认客户端超时。
const DefaultBigUploadTimeout = 30 * time.Second

// BigUploadFileResp 大文件上传（multipart/form-data）的返回结果。
type BigUploadFileResp struct {
	Data struct {
		FileKey  string `json:"file_key"`
		FileName string `json:"filename"`
		FileId   string `json:"file_id"`
		FileSize int64  `json:"file_size"`
		Md5      string `json:"md5"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// BigUploadFileIntegration 通过 multipart/form-data 上传大文件（默认 30s 超时）。
// 端点: POST /wxwork/BigUploadFile
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - authKey: 认证密钥
//   - fileKey: 文件唯一标识
//   - file: 文件内容（io.Reader）
//   - fileName: 文件名
//
// 返回: *BigUploadFileResp 包含 file_key、file_id、file_size 和 MD5 等信息
func (s *SdkIPad) BigUploadFileIntegration(host, uuid, authKey, fileKey string, file io.Reader, fileName string) (*BigUploadFileResp, error) {
	return s.BigUploadFileIntegrationWithTimeout(host, uuid, authKey, fileKey, file, fileName, 0)
}

// BigUploadFileIntegrationWithTimeout 同 BigUploadFileIntegration，但可为单次上传指定超时。
// 大文件（如朋友圈 10MB 图片）上传耗时可能远超全局 30s 客户端超时，
// 仅对该请求放宽，不影响其他 SDK 调用（timeout <= 0 时使用默认 30s）。
func (s *SdkIPad) BigUploadFileIntegrationWithTimeout(host, uuid, authKey, fileKey string, file io.Reader, fileName string, timeout time.Duration) (*BigUploadFileResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	if timeout <= 0 {
		timeout = DefaultBigUploadTimeout
	}
	// 复用全局连接池，仅覆盖本次请求的超时时间
	client := NewDefaultHTTPClient(&http.Client{Timeout: timeout, Transport: optimizedTransport})
	fields := map[string]string{
		"uuid":    uuid,
		"authkey": authKey,
		"filekey": fileKey,
	}
	var resp BigUploadFileResp
	if err := client.PostMultipart(host+"/wxwork/BigUploadFile", fields, "file", fileName, file, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
