// api_message_send_app_msg.go SendAppMsg 发送小程序消息。

package sdk_ipad

import "errors"

// SendAppMsgResp 发送小程序消息的返回结果。
type SendAppMsgResp struct {
	Data    map[string]interface{} `json:"data"`
	Errcode int                    `json:"errcode"`
	Errmsg  string                 `json:"errmsg"`
}

// SendAppMsg 发送小程序消息。
// 端点: POST /wxwork/SendAppMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - sendUserId: 接收者 ID
//   - isRoom: 是否为群消息
//   - desc: 小程序描述
//   - appName: 小程序名称
//   - title: 消息标题
//   - weappIconUrl: 小程序图标地址
//   - pagepath: 小程序页面路径
//   - username: 小程序原始 ID
//   - appid: 小程序 AppID
//   - cdnKey: 封面 CDN 文件标识
//   - md5: 封面 MD5
//   - aesKey: 封面 AES 密钥
//   - fileSize: 封面文件大小（字节）
//
// 返回: *SendAppMsgResp 发送结果
func (s *SdkIPad) SendAppMsg(host, uuid string, sendUserId int64, isRoom bool, desc, appName, title, weappIconUrl, pagepath, username, appid, cdnKey, md5, aesKey string, fileSize int64) (*SendAppMsgResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":         uuid,
		"send_userid":  sendUserId,
		"desc":         desc,
		"appName":      appName,
		"title":        title,
		"weappIconUrl": weappIconUrl,
		"pagepath":     pagepath,
		"username":     username,
		"appid":        appid,
		"cdnkey":       cdnKey,
		"md5":          md5,
		"aeskey":       aesKey,
		"fileSize":     fileSize,
		"isRoom":       isRoom,
	}
	var resp SendAppMsgResp
	if err := s.httpClient.Post(host+"/wxwork/SendAppMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
