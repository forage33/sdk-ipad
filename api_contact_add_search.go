// api_contact_add_search.go AddSearch 通过搜索结果添加好友。

package sdk_ipad

import (
	"errors"

	"go.uber.org/zap"
)

// AddSearchResp 添加搜索结果好友的返回结果。
type AddSearchResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// AddSearch 通过搜索结果添加好友。
// 端点: POST /wxwork/AddSearch
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 目标用户 VID
//   - optionID: 操作 ID
//   - phone: 手机号
//   - content: 好友申请验证消息
//   - ticket: 搜索结果 ticket
//
// 返回: *AddSearchResp 添加结果
func (s *SdkIPad) AddSearch(host, uuid string, vid int64, optionID, phone, content, ticket string) (*AddSearchResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":     uuid,
		"vid":      vid,
		"optionid": optionID,
		"phone":    phone,
		"content":  content,
		"ticket":   ticket,
	}

	s.logger.Info("AddSearch 请求", zap.String("host", host), zap.String("wx_uuid", uuid), zap.Int64("vid", vid), zap.String("option_id", optionID), zap.String("phone", phone), zap.String("ticket", ticket))

	var addSearchResp AddSearchResp
	if err := s.httpClient.Post(host+"/wxwork/AddSearch", reqParams, &addSearchResp); err != nil {
		s.logger.Error("AddSearch HTTP请求失败", zap.String("host", host), zap.String("wx_uuid", uuid), zap.Error(err))
		return nil, err
	}

	s.logger.Info("AddSearch 响应", zap.Int("errcode", addSearchResp.Errcode), zap.String("errmsg", addSearchResp.Errmsg))
	return &addSearchResp, nil
}
