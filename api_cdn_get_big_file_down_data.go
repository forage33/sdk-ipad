// api_cdn_get_big_file_down_data.go GetBigFileDownData 获取大文件下载所需参数（URL、Cookie 等），用于自行下载。

package sdk_ipad

import "errors"

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
