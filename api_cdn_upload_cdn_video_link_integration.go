// api_cdn_upload_cdn_video_link_integration.go UploadCdnVideoLinkIntegration 通过网络 URL 上传视频到 CDN。

package sdk_ipad

import "errors"

// UploadVideoResp CDN 上传网络视频的返回结果。
type UploadVideoResp struct {
	Data struct {
		VideoHeight   int    `json:"video_height"`
		Size          int64  `json:"size"`
		VideoDuration int    `json:"VideoDuration"` // NOTE: API 返回 PascalCase，与其他字段 snake_case 风格不一致
		CdnKey        string `json:"cdn_key"`
		AesKey        string `json:"aes_key"`
		VideoWidth    int    `json:"video_width"`
		VideoImgSize  int64  `json:"video_img_size"`
		Fileid        string `json:"fileid"`
		Md5           string `json:"md5"`
		Retcode       string `json:"retcode"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// UploadCdnVideoLinkIntegration 通过网络 URL 上传视频到 CDN。
// 端点: POST /wxwork/UploadCdnVideoLink
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - url: 视频网络地址（小于 25MB，超过需走大文件上传）
//   - fileName: 文件名
//   - videoImgUrl: 视频封面图片网络地址
//
// 返回: *UploadVideoResp 包含 CDN Key、AES Key、视频尺寸、时长等信息
func (s *SdkIPad) UploadCdnVideoLinkIntegration(host, uuid, url, fileName, videoImgUrl string) (*UploadVideoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":          uuid,
		"url":           url,
		"filename":      fileName,
		"video_img_url": videoImgUrl,
	}
	var resp UploadVideoResp
	if err := s.httpClient.Post(host+"/wxwork/UploadCdnVideoLink", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
