// callback_payload_account.go 登录/群/联系人/断连回调的数据契约层。
//
// 与 callback_payload_msg.go 同为协议回调契约：字段名与 docs/callback/ 对应
// 分册一致，协议字段知识收敛于此。
//
// 解析语义约定（与历史业务实现严格等价）：
//   - 登录域（扫码/登录成功等）历史上使用 encoding/json 反序列化并以错误
//     判定丢弃畸形数据，故访问器保持 json.Unmarshal 严格语义（返回 error）；
//   - 群/联系人/断连域历史上使用 gjson 逐字段宽容解析，访问器保持 gjson
//     宽容语义（缺失/类型不匹配返回零值，不报错）。

package sdk_ipad

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

// ---------------------------------------------------------------------------
// 登录域（严格解析语义）
// ---------------------------------------------------------------------------

// ScanReturnPayload 扫码返回（ScanReturn 100001）。
type ScanReturnPayload struct {
	Vid       int64  `json:"Vid"`
	QrcodeKey string `json:"Qrcode_key"`
	CorpID    int64  `json:"Corpid"`
	NickName  string `json:"Nick_name"`
	IconUrl   string `json:"Icon_url"`
	Logo      string `json:"Logo"`
}

// AsScanReturn 解析扫码返回 payload；畸形 JSON 时返回错误（由调用方决定跳过）。
func (m *CallBackMsg) AsScanReturn() (*ScanReturnPayload, error) {
	var p ScanReturnPayload
	err := json.Unmarshal([]byte(m.Json), &p)
	return &p, err
}

// CancelReturnPayload 取消扫码（CancelReturn 100003）。
type CancelReturnPayload struct {
	Vid       int64  `json:"Vid"`
	QrcodeKey string `json:"Qrcode_key"`
	CorpID    int64  `json:"Corpid"`
	NickName  string `json:"Nick_name"`
	IconUrl   string `json:"Icon_url"`
	Logo      string `json:"Logo"`
}

// AsCancelReturn 解析取消扫码 payload。
func (m *CallBackMsg) AsCancelReturn() (*CancelReturnPayload, error) {
	var p CancelReturnPayload
	err := json.Unmarshal([]byte(m.Json), &p)
	return &p, err
}

// ConfirmReturnPayload 确认扫码（ConfirmReturn 100002）。
type ConfirmReturnPayload struct {
	Vid       int64  `json:"Vid"`
	QrcodeKey string `json:"Qrcode_key"`
	CorpID    int64  `json:"Corpid"`
	NickName  string `json:"Nick_name"`
	IconUrl   string `json:"Icon_url"`
	Logo      string `json:"Logo"`
}

// AsConfirmReturn 解析确认扫码 payload。
func (m *CallBackMsg) AsConfirmReturn() (*ConfirmReturnPayload, error) {
	var p ConfirmReturnPayload
	err := json.Unmarshal([]byte(m.Json), &p)
	return &p, err
}

// LoginSuccessPayload 登录成功（LoginSuccessReturn 104001）。
//
// 字段说明：协议同时下发 corp_name（小写，企业名称，有值）与 Corp_name（大写开头，
// 协议示例恒为空串）；历史实现分别映射 CorpName/CorpName1 双字段，其中 Corp_name
// 映射字段业务从未使用，契约层收敛时剔除，CorpName 保留有值映射。
type LoginSuccessPayload struct {
	CorpID       int    `json:"Corpid"`
	Unionid      string `json:"unionid"`
	CreateTime   int64  `json:"create_time"`
	AdminVid     int64  `json:"Admin_vid"`
	Sex          int    `json:"sex"`
	Mobile       string `json:"mobile"`
	CorpFullName string `json:"Corp_full_name"`
	AcctID       string `json:"acctid"`
	Avatar       string `json:"avatar"`
	CorpName     string `json:"corp_name"`
	CreateTime1  int    `json:"Create_time"`
	EnglishName  string `json:"english_name"`
	RealName     string `json:"realname"`
	Vid          int64  `json:"Vid"`
	Mail         string `json:"Mail"`
	OwnerName    string `json:"Ownername"`
	UserID       int64  `json:"user_id"`
	Nickname     string `json:"nickname"`
	Position     string `json:"position"`
	CorpDesc     string `json:"Corp_desc"`
	CorpID2      int64  `json:"corp_id"`
	CorpLogo     string `json:"Corp_logo"`
}

// AsLoginSuccess 解析登录成功 payload。
func (m *CallBackMsg) AsLoginSuccess() (*LoginSuccessPayload, error) {
	var p LoginSuccessPayload
	err := json.Unmarshal([]byte(m.Json), &p)
	return &p, err
}

// NeedInputCodePayload 需要输入验证码（InputCode 100004）。
type NeedInputCodePayload struct {
	Vid       int64  `json:"Vid"`
	QrcodeKey string `json:"Qrcode_key"`
	CorpID    int64  `json:"Corpid"`
	NickName  string `json:"Nick_name"`
	IconUrl   string `json:"Icon_url"`
	Logo      string `json:"Logo"`
}

// AsNeedInputCode 解析输入验证码 payload。
func (m *CallBackMsg) AsNeedInputCode() (*NeedInputCodePayload, error) {
	var p NeedInputCodePayload
	err := json.Unmarshal([]byte(m.Json), &p)
	return &p, err
}

// ---------------------------------------------------------------------------
// 群域（宽容解析语义）
// ---------------------------------------------------------------------------

// RoomID 提取群事件回调中的群 ID（room_id，缺失时为空串）。
func (m *CallBackMsg) RoomID() string {
	return gjson.Get(m.Json, "room_id").String()
}

// GroupRoomPayload 建群回调（GroupCreate）payload。
type GroupRoomPayload struct {
	RoomID          string               `json:"room_id"`
	CreateUserID    int64                `json:"create_user_id"`
	Total           int                  `json:"total"`
	NoticeSenderVid int64                `json:"notice_sendervid"`
	Flag            int                  `json:"flag"`
	CreateTime      int64                `json:"create_time"`
	NoticeContent   string               `json:"notice_content"`
	Nickname        string               `json:"nickname"`
	MemberList      []GroupMemberPayload `json:"member_list"`
	NewFlag         int                  `json:"new_flag"`
	NoticeTime      int64                `json:"notice_time"`
	Managers        []interface{}        `json:"managers"`
}

// AsNewRoom 解析建群回调 payload；畸形 JSON 时返回错误（由调用方决定跳过）。
func (m *CallBackMsg) AsNewRoom() (*GroupRoomPayload, error) {
	var p GroupRoomPayload
	err := json.Unmarshal([]byte(m.Json), &p)
	return &p, err
}

// GroupMemberPayload 群成员信息（member_list 元素）。
//
// Uin 为兼容字段：协议各版本中可能为数字或字符串，保持 interface{}，
// 业务侧比较时自行格式化（历史实现即按此处理）。
type GroupMemberPayload struct {
	Unionid      string      `json:"unionid"`
	CreateTime   int64       `json:"create_time"`
	Sex          int         `json:"sex"`
	Mobile       string      `json:"mobile"`
	Acctid       string      `json:"acctid"`
	JoinScene    int         `json:"join_scene"`
	Avatar       string      `json:"avatar"`
	EnglishName  string      `json:"english_name"`
	Realname     string      `json:"realname"`
	RoomNotes    string      `json:"room_notes"`
	Jointime     int         `json:"jointime"`
	Nickname     string      `json:"nickname"`
	RoomNickname string      `json:"room_nickname"`
	Position     string      `json:"position"`
	Uin          interface{} `json:"uin"`
	InviteUserID int64       `json:"invite_user_id"`
	CorpID       int64       `json:"corp_id"`
}

// RoomMemberRef 群成员引用（acctid/uin 双字段），用于成员变动事件的账号归属判定。
//
// 系统内统一约定 uin（数字）== VID；acctid 是企微账号字符串标识，部分协议
// 版本 member_list 可能只有 acctid——双字段兼容，避免成员事件识别失败。
type RoomMemberRef struct {
	Acctid string // acctid 企微账号字符串标识
	UIN    int64  // uin 数字 VID（gjson Int 宽容语义）
}

// GroupMemberList 解析回调中的 member_list 为成员引用列表
// （宽容语义：字段缺失返回空切片）。
func (m *CallBackMsg) GroupMemberList() []RoomMemberRef {
	arr := gjson.Get(m.Json, "member_list").Array()
	members := make([]RoomMemberRef, 0, len(arr))
	for _, item := range arr {
		members = append(members, RoomMemberRef{
			Acctid: item.Get("acctid").String(),
			UIN:    item.Get("uin").Int(),
		})
	}
	return members
}

// GroupNameChangePayload 群名变更回调字段（三级回退已收敛）。
type GroupNameChangePayload struct {
	RoomID    string // room_id
	GroupName string // 群名：room_name 优先，回退 nickname / name（协议各事件字段不稳定，历史约定的宽容回退）
}

// AsGroupNameChange 解析群名变更字段。
func (m *CallBackMsg) AsGroupNameChange() GroupNameChangePayload {
	j := gjson.Parse(m.Json)
	groupName := j.Get("room_name").String()
	if groupName == "" {
		groupName = j.Get("nickname").String()
	}
	if groupName == "" {
		groupName = j.Get("name").String()
	}
	return GroupNameChangePayload{
		RoomID:    j.Get("room_id").String(),
		GroupName: groupName,
	}
}

// ---------------------------------------------------------------------------
// 联系人域（宽容解析语义）
// ---------------------------------------------------------------------------

// FriendChangeItem 好友变更回调（116xxx）list 数组元素。
//
// UserID 为字符串视图（gjson String 语义），与 VID 的字符串比较口径一致；
// SourceType 取自 source_info.source_type，status=0（外部推送共享）时用于过滤。
type FriendChangeItem struct {
	Status          int64  // status 0=外部推送共享 9=新增好友（其余值业务侧按需处理）
	UserID          string // user_id（缺失时为空串）
	Nickname        string // nickname
	RealName        string // realname
	Avatar          string // avatar
	AddCustomerTime int64  // add_customer_time 添加客户时间（秒）
	SourceType      int64  // source_info.source_type
}

// AsFriendChangeList 解析好友变更列表（宽容语义：list 缺失返回空切片）。
func (m *CallBackMsg) AsFriendChangeList() []FriendChangeItem {
	arr := gjson.Get(m.Json, "list").Array()
	items := make([]FriendChangeItem, 0, len(arr))
	for _, user := range arr {
		items = append(items, FriendChangeItem{
			Status:          user.Get("status").Int(),
			UserID:          user.Get("user_id").String(),
			Nickname:        user.Get("nickname").String(),
			RealName:        user.Get("realname").String(),
			Avatar:          user.Get("avatar").String(),
			AddCustomerTime: user.Get("add_customer_time").Int(),
			SourceType:      user.Get("source_info.source_type").Int(),
		})
	}
	return items
}

// InnerLabelChangeItem 内部联系人标签描述变动回调 list 元素
// （回调数据格式：{"list":[{"vid":1688xxx,"description":"...","remark":"...","timestamp":xxx}],"seq":xxx}）。
type InnerLabelChangeItem struct {
	Vid         int64  // vid 内部联系人 VID
	Description string // description 标签描述
	Remark      string // remark 备注
	Timestamp   int64  // timestamp 变动时间
}

// AsInnerLabelChangeList 解析内部标签变动列表（宽容语义：list 缺失返回空切片）。
func (m *CallBackMsg) AsInnerLabelChangeList() []InnerLabelChangeItem {
	arr := gjson.Get(m.Json, "list").Array()
	items := make([]InnerLabelChangeItem, 0, len(arr))
	for _, item := range arr {
		items = append(items, InnerLabelChangeItem{
			Vid:         item.Get("vid").Int(),
			Description: item.Get("description").String(),
			Remark:      item.Get("remark").String(),
			Timestamp:   item.Get("timestamp").Int(),
		})
	}
	return items
}

// ContactEventPayload 联系人事件提示消息（被删除/被拉黑等通知）payload。
type ContactEventPayload struct {
	Content string // msg 提示内容
}

// AsContactEventMsg 解析联系人事件提示消息。
func (m *CallBackMsg) AsContactEventMsg() ContactEventPayload {
	return ContactEventPayload{Content: gjson.Get(m.Json, "msg").String()}
}

// ---------------------------------------------------------------------------
// 断连域（宽容解析语义）
// ---------------------------------------------------------------------------

// DisconnectVid 解析断连类回调中的 vid（兼容 string 和 number 类型）。
//
// 返回 (vid, found)；vid 字段不存在时 found=false（由调用方记录告警并跳过，
// 与历史 parseDisconnectVid 的缺失日志行为对应）。
func (m *CallBackMsg) DisconnectVid() (string, bool) {
	result := gjson.Get(m.Json, "vid")
	if !result.Exists() {
		return "", false
	}
	return result.String(), true
}
