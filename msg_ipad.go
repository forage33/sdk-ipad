package sdk_ipad

// msg_ipad.go 定义回调消息类型常量及回调/发送消息的基础结构体。
// 所有通过 /callback 接口推送的消息均使用此处定义的类型码进行分类。

// 回调消息类型常量。
//
// 按业务场景分组：登录、消息、群、联系人、标签、聊天标签、断开连接。
// 数值范围约定：
//   - 100xxx: 登录 / 连接 / 断开类
//   - 102xxx: 聊天消息 / 已读 / 删除 / 拉黑 / 聊天标签
//   - 115xxx: 群变动类
//   - 116xxx: 联系人变动类
//   - 130xxx: 标签变动类
const (
	// ---- 登录 ----

	ScanReturn         = 100001 // 扫码返回数据
	CancelReturn       = 100003 // 取消扫码返回
	ConfirmReturn      = 100002 // 确认扫码返回数据
	LoginSuccessReturn = 104001 // 登录成功返回数据
	CorpChange         = 100006 // 企业切换回调
	InputCode          = 100004 // 需要输入验证码消息回调

	// ---- 消息 ----

	MsgCallback = 102000 // 聊天消息回调

	// ---- 群 ----

	NewGroup              = 115001 // 群新增通知
	RemoveGroupMember     = 115002 // 移除群成员通知
	InviteInGroup         = 115003 // 邀请进群通知
	QuitGroup             = 115004 // 自己主动退出群聊通知（包括手机端或其他端退出）
	GroupNameChange       = 115005 // 群名变更通知
	GroupSuperAdminChange = 115006 // 群主转让通知及其他设置通知
	GroupAdminChange      = 115007 // 群管理员变动通知
	GroupDisBand          = 115008 // 群解散通知
	GroupMemberChange     = 115009 // 群变动通知

	// ---- 联系人 ----

	ApplyFriend      = 116001 // 联系人添加申请通知
	FriendChange     = 116002 // 联系人变动通知
	InnerLabelChange = 116003 // 内部联系人标签描述变动通知
	MsgRead          = 102006 // 当前联系人已读消息通知回调
	HasDeleted       = 102001 // 联系人发消息提示删除通知
	HasBlacked       = 102002 // 联系人发消息提示拉黑通知
	PinnedMsg        = 116005 // 置顶免打扰标记操作通知

	// ---- 标签变动 ----

	MyselfLabelChange = 130001 // 个人标签新增/删除回调
	CorpLabelChange   = 130002 // 企业标签新增/删除回调

	// ---- 聊天标签 ----

	ChatLabelChange       = 102004 // 聊天标签变动通知
	ChatLabelMemberChange = 102005 // 聊天标签中的联系人变动通知

	// ---- 断开 / 异常 ----

	SecondAuth         = 100012 // 二次验证消息通知
	EndLoginByPhone    = 100005 // 手机端结束登录
	LoginInOtherDevice = 100008 // 当前账号在其他设备登录
	OverTime           = 100009 // 初始化接口二维码扫码超时断开
	Exception          = 100007 // 异常断开
)

// 登录状态字符串常量。
//
// 用于在轮询登录流程中标识当前阶段。
const (
	WaitCode     = "waitCode"     // 等待获取二维码
	UNScan       = "unScan"       // 未扫码
	Scanning     = "scanning"     // 扫码返回
	CancelScan   = "cancelScan"   // 取消扫码
	Scanned      = "scanned"      // 确认扫码
	LoginSuccess = "loginSuccess" // 登录成功
	VerifyCode   = "verifyCode"   // 需要输入验证码消息回调
	AutoLogging  = "autoLogging"  // 自动登录中（等待回调确认）
)

// CallBackMsg 回调消息结构体。
//
// 协议服务通过 /callback 接口推送的统一消息格式。
//
// 字段:
//   - Type: 消息类型，对应上方常量（如 ScanReturn、MsgCallback 等）
//   - UUID: 运行实例唯一标识
//   - Json: 消息体 JSON 字符串，具体内容根据 Type 不同而不同
type CallBackMsg struct {
	Type int    `json:"type"` // 消息类型
	UUID string `json:"uuid"` // 运行实例
	Json string `json:"json"` // json字符串
}

// SendMsg 发送消息结构体。
//
// 用于内部传递待发送消息的上下文信息。
//
// 字段:
//   - SyncKey: 同步 key
//   - Host: 协议服务地址
//   - Identity: 标识
//   - VID: 企业微信 VID
//   - UUID: 运行实例唯一标识
//   - MsgType: 消息类型
//   - Content: 消息内容
//   - To: 接收方 ID
//   - IsRoom: 是否为群聊消息
type SendMsg struct {
	SyncKey        string  `json:"syncKey"`  // 同步key
	Host           string  `json:"host"`     // 调用地址
	Identity       string  `json:"identity"` // 标识
	VID            string  `json:"vid"`
	AccountID      int64   `json:"accountID"`
	UUID           string  `json:"uuid"`
	MsgType        int     `json:"msgType"`
	Content        string  `json:"content"`
	To             string  `json:"to"`
	IsRoom         bool    `json:"isRoom"`
	TenantID       int64   `json:"tenantID"`       // 租户ID，用于失败时更新消息状态
	ConversationID string  `json:"conversationID"` // 会话ID，用于失败时更新消息状态
	CustomerName   string  `json:"customerName"`   // 客服名称，用于失败时推送 WS 通知
	AtUserIds      []int64 `json:"atUserIds"`      // 被@的群成员 uin 列表（群聊 at 消息，仅 IsRoom=true 时生效）
}
