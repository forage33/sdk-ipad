// api_cdn_download_we_chat_file.go DownloadWeChatFile 下载微信文件（图片、视频等）。

package sdk_ipad

import "errors"

// DownloadWechatFileResp 微信文件下载的返回结果。
// 注意：协议服务实际返回 errcode 字段（非 errorcode），字段名不匹配会静默吞掉错误。
type DownloadWechatFileResp struct {
	Data    string `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// DownloadWeChatFile 下载微信文件（图片、视频等）。
// 端点: POST /wxwork/DownloadWeChatFile
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - url: 文件地址
//   - aesKey: AES 解密密钥
//   - authKey: 认证密钥
//   - fileName: 保存文件名
//   - size: 文件大小（字节）
//
// 返回: *DownloadWechatFileResp 包含文件数据
func (s *SdkIPad) DownloadWeChatFile(host, uuid, url, aesKey, authKey, fileName string, size int64) (*DownloadWechatFileResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":      uuid,
		"url":       url,
		"aes_key":   aesKey,
		"auth_key":  authKey,
		"file_name": fileName,
		"size":      size,
	}
	var resp DownloadWechatFileResp
	if err := s.httpClient.Post(host+"/wxwork/DownloadWeChatFile", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
