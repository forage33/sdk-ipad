// api_group_get_welcome_msg_list.go GetWelcomeMsgList 获取欢迎语列表。

package sdk_ipad

import "errors"

// WelcomeMsgItem 欢迎语列表项。
type WelcomeMsgItem struct {
	Updateid       int64  `json:"updateid"`
	Createtime     int64  `json:"createtime"`
	Createid       int64  `json:"createid"`
	FileId         string `json:"file_id"`
	AesKey         string `json:"aes_key"`
	Id             int    `json:"id"`
	Type           int    `json:"type"`
	ImgSize        int    `json:"img_size"`
	Updatetime     int64  `json:"updatetime"`
	Content        string `json:"content"`
	WelcineMsgLink string `json:"WelcineMsgLink"`
	Md5            string `json:"md5"`
}

// GetWelcomeMsgListResp 获取欢迎语列表的返回结果。
type GetWelcomeMsgListResp struct {
	Data struct {
		Offset          string           `json:"offset"`
		WelocomeMsgList []WelcomeMsgItem `json:"WelocomeMsgList"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetWelcomeMsgList 获取欢迎语列表。
// 端点: POST /wxwork/GetWelcomeMsgList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - offset: 分页偏移量
//   - limit: 每页数量
//
// 返回: *GetWelcomeMsgListResp 欢迎语列表
func (s *SdkIPad) GetWelcomeMsgList(host, uuid, offset string, limit int) (*GetWelcomeMsgListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"offset": offset,
		"limit":  limit,
	}
	var resp GetWelcomeMsgListResp
	if err := s.httpClient.Post(host+"/wxwork/GetWelcomeMsgList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
