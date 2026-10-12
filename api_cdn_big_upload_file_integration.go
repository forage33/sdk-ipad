// api_cdn_big_upload_file_integration.go BigUploadFileIntegration 通过 multipart/form-data 上传大文件。

package sdk_ipad

import (
	"errors"
	"io"
)

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

// BigUploadFileIntegration 通过 multipart/form-data 上传大文件。
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
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	fields := map[string]string{
		"uuid":    uuid,
		"authkey": authKey,
		"filekey": fileKey,
	}
	var resp BigUploadFileResp
	if err := s.httpClient.PostMultipart(host+"/wxwork/BigUploadFile", fields, "file", fileName, file, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
