// api_account_update_user_info.go UpdateUserInfo 修改当前用户的个人信息。

package sdk_ipad

import "errors"

// UpdateUserInfo 修改当前用户的个人信息。
// 端点: POST /wxwork/UpdateUserInfo
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - emile: 邮箱
//   - name: 姓名
//   - englishName: 英文名
//   - phone: 电话
//   - position: 职位
//   - mobile: 手机号
//   - gender: 性别（1=男, 2=女）
//
// 返回: *CommonAccountResp 操作结果
func (s *SdkIPad) UpdateUserInfo(host, uuid, emile, name, englishName, phone, position, mobile string, gender int) (*CommonAccountResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":         uuid,
		"emile":        emile,
		"name":         name,
		"english_name": englishName,
		"phone":        phone,
		"position":     position,
		"mobile":       mobile,
		"gender":       gender,
	}

	var resp CommonAccountResp
	if err := s.httpClient.Post(host+"/wxwork/UpdateUserInfo", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
