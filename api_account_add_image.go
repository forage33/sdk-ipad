// api_account_add_image.go AddImage 设置账号头像。

package sdk_ipad

import "errors"

// SetImageResp 设置头像的返回结果。
type SetImageResp struct {
	Data    ImageData `json:"data"`
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
}

// ImageData 头像数据。
type ImageData struct {
	ImageUrl string `json:"image_url"`
	ImageId  string `json:"image_id"`
}

// AddImage 设置账号头像。
// 端点: POST /wxwork/AddImage
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - url: 头像图片地址
//
// 返回: *SetImageResp 包含新头像 URL 和 ID
func (s *SdkIPad) AddImage(host, uuid, url string) (*SetImageResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"imgurl": url,
	}

	var setResp SetImageResp
	if err := s.httpClient.Post(host+"/wxwork/AddImage", reqParams, &setResp); err != nil {
		return nil, err
	}
	return &setResp, nil
}
