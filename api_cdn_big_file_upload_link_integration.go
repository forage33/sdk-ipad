// api_cdn_big_file_upload_link_integration.go BigFileUploadLinkIntegration 通过网络 URL 上传大文件到服务器。

package sdk_ipad

import "errors"

// BigFileUploadLinkResp 大文件网络上传的返回结果。
type BigFileUploadLinkResp struct {
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

// BigFileUploadLinkIntegration 通过网络 URL 上传大文件到服务器。
// 端点: POST /wxwork/BigFileUploadLink
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - authKey: 认证密钥
//   - fileKey: 文件唯一标识
//   - fileUrl: 文件网络地址
//   - fileName: 文件名
//
// 返回: *BigFileUploadLinkResp 包含 file_key、file_id、file_size 和 MD5 等信息
func (s *SdkIPad) BigFileUploadLinkIntegration(host, uuid, authKey, fileKey, fileUrl, fileName string) (*BigFileUploadLinkResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":     uuid,
		"authkey":  authKey,
		"filekey":  fileKey,
		"fileurl":  fileUrl,
		"filename": fileName,
	}
	var resp BigFileUploadLinkResp
	if err := s.httpClient.Post(host+"/wxwork/BigFileUploadLink", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
