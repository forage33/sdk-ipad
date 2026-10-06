// cdn.go CDN 相关接口。
// 提供网络图片上传、网络视频上传、网络文件上传、大文件上传下载和微信文件下载等功能。

package sdk_ipad

import (
	"errors"
	"io"
)

// ---------------------------------------------------------------------------
// 返回数据结构
// ---------------------------------------------------------------------------

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

// DownloadWechatFileResp 微信文件下载的返回结果。
// 注意：协议服务实际返回 errcode 字段（非 errorcode），字段名不匹配会静默吞掉错误。
type DownloadWechatFileResp struct {
	Data    string `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

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

// BigFileDownUrlResp 大文件下载的返回结果。
type BigFileDownUrlResp struct {
	Data struct {
		Url string `json:"url"`
		Md5 string `json:"md5"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetBigFileDownDataResp 大文件下载参数获取的返回结果。
type GetBigFileDownDataResp struct {
	Data struct {
		IsBigC2c     bool   `json:"isBigC2c"`
		AuthCookie   string `json:"authCookie"`
		ErrorMessage string `json:"errorMessage"`
		ErrorCode    int    `json:"errorCode"`
		Url          string `json:"url"`
		Md5          string `json:"md5"`
		Size         int64  `json:"size"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

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

// DownloadFileResp CDN 下载文件的返回结果。
type DownloadFileResp struct {
	Data    string `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 调用接口
// ---------------------------------------------------------------------------

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

// GetBigFileDownData 获取大文件下载所需参数（URL、Cookie 等），用于自行下载。
// 端点: POST /wxwork/getBigFileDownData
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - fileId: 文件 ID
//   - fileName: 文件名
//   - size: 文件大小（字节）
//
// 返回: *GetBigFileDownDataResp 包含下载 URL、authCookie、MD5 等信息
func (s *SdkIPad) GetBigFileDownData(host, uuid, fileId, fileName string, size int64) (*GetBigFileDownDataResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":      uuid,
		"fileid":    fileId,
		"file_name": fileName,
		"size":      size,
	}
	var resp GetBigFileDownDataResp
	if err := s.httpClient.Post(host+"/wxwork/getBigFileDownData", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
