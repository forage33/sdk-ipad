// api_cdn_c2c_to_wwfile_id.go C2CToWwfileId 大文件上传 fileid 转可发送 id。

package sdk_ipad

import "errors"

// C2CToWwfileIdResp 大文件上传 fileid 转可发送 id 的返回结果。
type C2CToWwfileIdResp struct {
	Data struct {
		MD5      string `json:"md5"`       // 文件的 MD5 哈希值
		Filename string `json:"filename"`  // 文件名
		FileID   string `json:"file_id"`   // 文件唯一标识符
		FileSize int64  `json:"file_size"` // 文件大小（字节）
		FileKey  string `json:"file_key"`  // 文件的唯一密钥
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// C2CToWwfileId 大文件上传 fileid 转可发送 id（大文件上传完成后经此转换方可发送）。
// 端点: POST /wxwork/c2c_to_wwfile_id
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - fileID: 文件 ID（大文件上传返回）
//   - fileSize: 文件大小（字节）
//   - fileKey: 文件密钥（GetBigAuthkey 返回的 filekey）
//   - fileMD5: 文件的 MD5 哈希值
//   - fileName: 文件名称
//
// 返回: *C2CToWwfileIdResp 含转换后可发送的 file_id/file_key 等
func (s *SdkIPad) C2CToWwfileId(host, uuid, fileID string, fileSize int64, fileKey, fileMD5, fileName string) (*C2CToWwfileIdResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":      uuid,
		"file_id":   fileID,
		"file_size": fileSize,
		"file_key":  fileKey,
		"file_md5":  fileMD5,
		"file_name": fileName,
	}
	var resp C2CToWwfileIdResp
	if err := s.httpClient.Post(host+"/wxwork/c2c_to_wwfile_id", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
