// api_contact_get_collea_remark_list.go GetColleaRemarkList 获取内部联系人备注列表。

package sdk_ipad

import "errors"

// ColleaRemarkListResp 内部联系人备注列表的返回结果。
type ColleaRemarkListResp struct {
	Data struct {
		List []struct {
			Vid         int64  `json:"vid"`
			Description string `json:"description"`
			Remark      string `json:"remark"`
			Timestamp   int64  `json:"timestamp"`
		} `json:"list"`
		Seq int `json:"seq"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetColleaRemarkList 获取内部联系人备注列表。
// 端点: POST /wxwork/GetColleaRemarkList
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - limit: 每页数量
//   - seq: 分页序列号
//
// 返回: *ColleaRemarkListResp 备注列表
func (s *SdkIPad) GetColleaRemarkList(host, uuid string, limit, seq int) (*ColleaRemarkListResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"limit": limit,
		"seq":   seq,
	}

	var resp ColleaRemarkListResp
	if err := s.httpClient.Post(host+"/wxwork/GetColleaRemarkList", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
