// api_cdn_upload_cdn_link_file_integration.go UploadCdnLinkFileIntegration 通过网络 URL 上传文件到 CDN。

package sdk_ipad

import "errors"

// UploadFileResp CDN 上传网络文件的返回结果。
type UploadFileResp struct {
	Data struct {
		Size   int    `json:"size"`
		CdnKey string `json:"cdn_key"`
		AesKey string `json:"aes_key"`
		Fileid string `json:"fileid"`
		Md5    string `json:"md5"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// UploadCdnLinkFileIntegration 通过网络 URL 上传文件到 CDN。
// 端点: POST /wxwork/UploadCdnLinkFile
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - url: 文件网络地址
//   - fileName: 文件名
//
// 返回: *UploadFileResp 包含 CDN Key、AES Key 和 MD5 等信息
func (s *SdkIPad) UploadCdnLinkFileIntegration(host, uuid, url, fileName string) (*UploadFileResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":     uuid,
		"url":      url,
		"filename": fileName,
	}
	var resp UploadFileResp
	if err := s.httpClient.Post(host+"/wxwork/UploadCdnLinkFile", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
