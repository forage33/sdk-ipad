// api_contact_add_search_wxwork.go AddSearchWxwork 搜索并添加企业微信用户为好友。

package sdk_ipad

import (
	"errors"

	"go.uber.org/zap"
)

// AddSearchWxwork 搜索并添加企业微信用户为好友。
// 端点: POST /wxwork/AddSearchWxwork
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - vid: 目标用户 VID
//   - corpid: 目标用户企业 ID
//   - ticket: 搜索 ticket
//   - content: 好友申请验证消息
//
// 返回: *CommonContactResp 操作结果
func (s *SdkIPad) AddSearchWxwork(host, uuid string, vid, corpid int64, ticket, content string) (*CommonContactResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	reqParams := map[string]interface{}{
		"uuid":    uuid,
		"vid":     vid,
		"corpid":  corpid,
		"ticket":  ticket,
		"content": content,
	}

	s.logger.Info("AddSearchWxwork 请求", zap.String("host", host), zap.String("wx_uuid", uuid), zap.Int64("vid", vid), zap.Int64("corpid", corpid), zap.String("ticket", ticket))

	var resp CommonContactResp
	if err := s.httpClient.Post(host+"/wxwork/AddSearchWxwork", reqParams, &resp); err != nil {
		s.logger.Error("AddSearchWxwork HTTP请求失败", zap.String("host", host), zap.String("wx_uuid", uuid), zap.Error(err))
		return nil, err
	}

	s.logger.Info("AddSearchWxwork 响应", zap.Int("errcode", resp.Errcode), zap.String("errmsg", resp.Errmsg), zap.Any("data", resp.Data))
	return &resp, nil
}
