// api_account.go 账号操作相关接口。
// 提供头像设置、用户信息修改、二维码获取、封禁详情、个人资料、
// 对外名片、自动同意好友和企业信息查询等功能。

package sdk_ipad

import (
	"errors"
)

// ---------------------------------------------------------------------------
// 返回数据结构
// ---------------------------------------------------------------------------

// SetImageResp 设置头像的返回结果。
type SetImageResp struct {
	Data    ImageData `json:"data"`
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
}

// ImageData 头像数据。
type ImageData struct {
	ImageUrl string `json:"image_url"`
	ImageId  string `json:"image_id"`
}

// CommonAccountResp 账号操作通用返回结果（data 为 null 或简单数据）。
type CommonAccountResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// GetQrCodeResp 获取账号二维码的返回结果。
type GetQrCodeResp struct {
	Data struct {
		QrCodeUrl string `json:"QrCodeUrl"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// BanResp 获取封禁详情的返回结果。
type BanResp struct {
	Data    BanData `json:"data"`
	Errcode int     `json:"errcode"`
	Errmsg  string  `json:"errmsg"`
}

// BanData 封禁数据。
type BanData struct {
	Vid          string   `json:"vid"`
	Gid          string   `json:"gid"`
	MainData     MainData `json:"mainData"`
	IsWechat     bool     `json:"isWechat"`
	IsWxwork     bool     `json:"isWxwork"`
	IsSuperadmin int      `json:"is_superadmin"`
	OrderId      string   `json:"order_id"`
	CorpId       string   `json:"corp_id"`
}

// MainData 封禁详细信息。
type MainData struct {
	BanStartTime  FlexInt64 `json:"ban_start_time"`
	AppealId      string    `json:"appeal_id"`
	BanSource     int       `json:"ban_source"`
	IsShowTitle   bool      `json:"is_show_title"`
	ErrCode       int       `json:"err_code"`
	BanTitle      string    `json:"ban_title"`
	Type          string    `json:"type"`
	IsShowReason  bool      `json:"is_show_reason"`
	IsShowFunc    bool      `json:"is_show_func"`
	BanFinishTime FlexInt64 `json:"ban_finish_time"`
	BanCategory   int       `json:"ban_category"`
	AppealStatus  int       `json:"appeal_status"`
	CanAppeal     bool      `json:"can_appeal"`
	BanIdType     int       `json:"ban_id_type"`
	NeedLogin     string    `json:"need_login"`
	BanReason     string    `json:"ban_reason"`
	TemplateType  int       `json:"template_type"`
	ExposedId     string    `json:"exposed_id"`
	BanFunc       string    `json:"ban_func"`
	IsShowTime    bool      `json:"is_show_time"`
	CanUrge       bool      `json:"can_urge"`
	Key           string    `json:"key"`
}

// ProfileResp 获取当前账号详情的返回结果。
type ProfileResp struct {
	Data struct {
		UnionID      string `json:"unionid"`
		CreateTime   int64  `json:"create_time"`
		Sex          int    `json:"sex"`
		Mobile       string `json:"mobile"`
		CorpFullName string `json:"corp_full_name"`
		AcctID       string `json:"acctid"`
		SCorpID      string `json:"scorp_id"`
		Avatar       string `json:"avatar"`
		CorpName     string `json:"corp_name"`
		EnglishName  string `json:"english_name"`
		RealName     string `json:"realname"`
		UserID       int64  `json:"user_id"`
		Nickname     string `json:"nickname"`
		Position     string `json:"position"`
		CorpDesc     string `json:"corp_desc"`
		CorpID       int64  `json:"corp_id"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ExternalBusinessCardResp 获取对外名片二维码的返回结果。
type ExternalBusinessCardResp struct {
	Data struct {
		QrImage string `json:"qr_image"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// CorpInfoItem 企业详情信息。
type CorpInfoItem struct {
	Staffnum       int    `json:"Staffnum"`
	AuthTime       int64  `json:"Auth_time"`
	CardUrl        string `json:"Card_url"`
	Corpid         int64  `json:"Corpid"`
	AdminVid       int64  `json:"Admin_vid"`
	VSuperadminVid int64  `json:"V_superadmin_vid"`
	CorpStat       string `json:"Corp_stat"`
	CorpFullName   string `json:"Corp_full_name"`
	IsAccepted     bool   `json:"IsAccepted"`
	CreateTime     int64  `json:"Create_time"`
	Vid            int64  `json:"Vid"`
	ScorpId        string `json:"Scorp_id"`
	Mail           string `json:"Mail"`
	Ownername      string `json:"Ownername"`
	CorpDesc       string `json:"Corp_desc"`
	CorpName       string `json:"Corp_name"`
	CorpLogo       string `json:"Corp_logo"`
}

// GetCorpInfoResp 获取企业详情的返回结果。
type GetCorpInfoResp struct {
	Data struct {
		List []CorpInfoItem `json:"list"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 调用接口
// ---------------------------------------------------------------------------

// AddImage 设置账号头像。
// 端点: POST /wxwork/AddImage
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - url: 头像图片地址
//
// 返回: *SetImageResp 包含新头像 URL 和 ID
func (s *SdkIPad) AddImage(host, uuid, url string) (*SetImageResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"imgurl": url,
	}

	var setResp SetImageResp
	if err := s.httpClient.Post(host+"/wxwork/AddImage", reqParams, &setResp); err != nil {
		return nil, err
	}
	return &setResp, nil
}

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

// GetThisQRCode 获取当前账号的二维码。
// 端点: POST /wxwork/GetThisQrCode
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *GetQrCodeResp 包含二维码图片地址
func (s *SdkIPad) GetThisQRCode(host, uuid string) (*GetQrCodeResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
	}

	var qrCodeResp GetQrCodeResp
	if err := s.httpClient.Post(host+"/wxwork/GetThisQrCode", reqParams, &qrCodeResp); err != nil {
		return nil, err
	}

	// 使用HTTP客户端发送请求
	return &qrCodeResp, nil
}

// GetBanDetail 获取账号封禁详情。
// 端点: POST /wxwork/getBanDetile
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - url: 封禁详情页地址
//
// 返回: *BanResp 封禁详细信息
func (s *SdkIPad) GetBanDetail(host, uuid, url string) (*BanResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]string{
		"uuid": uuid,
		"url":  url,
	}

	var banResp BanResp
	if err := s.httpClient.Post(host+"/wxwork/getBanDetile", reqParams, &banResp); err != nil {
		return nil, err
	}

	return &banResp, nil
}

// GetProfile 获取当前登录账号的详细资料。
// 端点: POST /wxwork/GetProfile
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *ProfileResp 账号详细信息
func (s *SdkIPad) GetProfile(host, uuid string) (*ProfileResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]string{
		"uuid": uuid,
	}
	var resp ProfileResp
	if err := s.httpClient.Post(host+"/wxwork/GetProfile", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExternalBusinessCard 获取当前账号的对外名片二维码。
// 端点: POST /wxwork/ExternalBusinessCard
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//
// 返回: *ExternalBusinessCardResp 包含名片二维码图片
func (s *SdkIPad) ExternalBusinessCard(host, uuid string) (*ExternalBusinessCardResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]string{
		"uuid": uuid,
	}

	var resp ExternalBusinessCardResp
	if err := s.httpClient.Post(host+"/wxwork/ExternalBusinessCard", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AutomaticConsent 设置是否自动同意好友申请。
// 端点: POST /wxwork/AutomaticConsent
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - state: 0=关闭验证自动同意好友, 1=需要验证同意（手动点击同意）
//
// 返回: *CommonAccountResp 操作结果
func (s *SdkIPad) AutomaticConsent(host, uuid string, state int) (*CommonAccountResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"state": state,
	}

	var resp CommonAccountResp
	if err := s.httpClient.Post(host+"/wxwork/AutomaticConsent", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCorpInfo 根据企业 ID 获取企业详情。
// 端点: POST /wxwork/GetCorpInfo
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - corpid: 企业 ID
//
// 返回: *GetCorpInfoResp 企业详细信息
func (s *SdkIPad) GetCorpInfo(host, uuid string, corpid int64) (*GetCorpInfoResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":   uuid,
		"corpid": corpid,
	}

	var resp GetCorpInfoResp
	if err := s.httpClient.Post(host+"/wxwork/GetCorpInfo", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
