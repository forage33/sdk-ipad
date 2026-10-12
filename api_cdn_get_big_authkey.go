// api_cdn_get_big_authkey.go GetBigAuthkey 获取大文件上传 AuthKey。

package sdk_ipad

import "errors"

// GetBigAuthkeyResp 获取大文件上传 AuthKey 的返回结果。
type GetBigAuthkeyResp struct {
	Data struct {
		FileKey string `json:"filekey"` // 文件键，用于唯一标识文件
		Vid     int64  `json:"vid"`     // 当前账号 id
		AuthKey string `json:"AuthKey"` // 认证密钥，用于授权访问相关资源
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetBigAuthkey 获取大文件上传 AuthKey。
// 端点: POST /wxwork/GetBigAuthkey
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *GetBigAuthkeyResp 含 filekey/vid/AuthKey
//
// 注意: 从账号登录服务获取，哪个账号操作大文件就取哪个账号的，不可混用。
func (s *SdkIPad) GetBigAuthkey(host, uuid string) (*GetBigAuthkeyResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
	}
	var resp GetBigAuthkeyResp
	if err := s.httpClient.Post(host+"/wxwork/GetBigAuthkey", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
