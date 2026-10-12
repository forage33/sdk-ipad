// api_group_add_welcome_msg.go AddWelcomeMsg 添加欢迎语。

package sdk_ipad

import "errors"

// AddWelcomeMsgResp 添加欢迎语的返回结果。
type AddWelcomeMsgResp struct {
	Data struct {
		Id int `json:"id"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AddWelcomeMsg 添加欢迎语。
// 端点: POST /wxwork/AddWelcomeMsg
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - list: 欢迎语文本列表
//   - url: 链接地址
//   - title: 链接标题
//   - content: 链接描述
//   - headImg: 链接缩略图
//
// 返回: *AddWelcomeMsgResp 包含新增欢迎语 ID
func (s *SdkIPad) AddWelcomeMsg(host, uuid string, list []string, url, title, content, headImg string) (*AddWelcomeMsgResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"list":    list,
		"url":     url,
		"title":   title,
		"content": content,
		"headImg": headImg,
	}
	var resp AddWelcomeMsgResp
	if err := s.httpClient.Post(host+"/wxwork/AddWelcomeMsg", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
