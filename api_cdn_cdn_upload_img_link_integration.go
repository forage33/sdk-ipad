// api_cdn_cdn_upload_img_link_integration.go CdnUploadImgLinkIntegration 通过网络 URL 上传图片到 CDN。

package sdk_ipad

import "errors"

// UploadImgResp CDN 上传网络图片的返回结果。
type UploadImgResp struct {
	Data struct {
		CdnKey           string `json:"cdn_key"`
		Fileid           string `json:"fileid"`
		AesKey           string `json:"aes_key"`
		Md5              string `json:"md5"`
		Width            int    `json:"width"`
		Height           int    `json:"height"`
		Size             int    `json:"size"`
		MidImageSize     int    `json:"mid_image_size"`
		ThumbImageWidth  int    `json:"thumb_image_width"`
		ThumbImageHeight int    `json:"thumb_image_height"`
		ThumbFileSize    int    `json:"thumb_file_size"`
		ThumbFileMd5     string `json:"thumb_file_md5"`
		Retcode          string `json:"retcode"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// CdnUploadImgLinkIntegration 通过网络 URL 上传图片到 CDN。
// 端点: POST /wxwork/CdnUploadImgLink
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - url: 图片网络地址
//
// 返回: *UploadImgResp 包含 CDN Key、AES Key、MD5 和尺寸等信息
func (s *SdkIPad) CdnUploadImgLinkIntegration(host, uuid, url string) (*UploadImgResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid": uuid,
		"url":  url,
	}

	var resp UploadImgResp
	if err := s.httpClient.Post(host+"/wxwork/CdnUploadImgLink", reqParams, &resp); err != nil {
		return nil, err
	}

	// 使用HTTP客户端发送请求
	return &resp, nil
}
