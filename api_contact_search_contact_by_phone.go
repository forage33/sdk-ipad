// api_contact_search_contact_by_phone.go SearchContactByPhone 根据手机号搜索联系人。

package sdk_ipad

import (
	"errors"

	"go.uber.org/zap"
)

// SearchPhoneResp 手机号搜索联系人的返回结果。
type SearchPhoneResp struct {
	Data    SearchPhoneData `json:"data"`
	Errcode int             `json:"errcode"`
	Errmsg  string          `json:"errmsg"`
}

// SearchPhoneData 搜索结果数据。
type SearchPhoneData struct {
	UserList []UserInfo `json:"UserList"`
}

// SearchContactByPhone 根据手机号搜索联系人。
// 端点: POST /wxwork/SearchContact
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - phoneNum: 手机号
//
// 返回: *SearchPhoneResp 搜索结果列表
func (s *SdkIPad) SearchContactByPhone(host, uuid, phoneNum string) (*SearchPhoneResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}

	// 请求参数
	reqParams := map[string]string{
		"uuid":        uuid,
		"phoneNumber": phoneNum,
	}

	s.logger.Info("SearchContactByPhone 请求", zap.String("host", host), zap.String("wx_uuid", uuid), zap.String("phone", phoneNum))

	var searchResp SearchPhoneResp
	if err := s.httpClient.Post(host+"/wxwork/SearchContact", reqParams, &searchResp); err != nil {
		s.logger.Error("SearchContactByPhone HTTP请求失败", zap.String("host", host), zap.String("wx_uuid", uuid), zap.String("phone", phoneNum), zap.Error(err))
		return nil, err
	}

	s.logger.Info("SearchContactByPhone 响应", zap.Int("errcode", searchResp.Errcode), zap.String("errmsg", searchResp.Errmsg), zap.Int("user_count", len(searchResp.Data.UserList)))
	// 使用HTTP客户端发送请求
	return &searchResp, nil
}
