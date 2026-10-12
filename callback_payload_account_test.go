package sdk_ipad

import "testing"

// callback_payload_account_test.go 登录/群/联系人/断连域契约层
// （callback_payload_account.go）的语义契约测试。
//
// 覆盖的语义铁律（修改契约层前须保持测试通过）：
//   - 登录域严格语义：json.Unmarshal 判错丢弃畸形数据（返回 error）
//   - 群/联系人/断连域宽容语义：gjson 逐字段取值，缺失返回零值
//   - 群成员 uin 数字/字符串双型兼容（协议各版本不一致）
//   - 群名三级回退：room_name → nickname → name
//   - 断连 vid 双型兼容（string/number），缺失时 found=false

// TestAsScanReturn 扫码返回：正常解析（大写协议字段名）与畸形 JSON 严格报错。
func TestAsScanReturn(t *testing.T) {
	msg := &CallBackMsg{Json: `{
		"Vid": 1688851000000001,
		"Qrcode_key": "qr-key-1",
		"Corpid": 1970324836974591,
		"Nick_name": "小助手",
		"Icon_url": "http://wx.example/icon.png",
		"Logo": "http://wx.example/logo.png"
	}`}
	p, err := msg.AsScanReturn()
	if err != nil {
		t.Fatalf("正常 JSON 不应报错: %v", err)
	}
	if p.Vid != 1688851000000001 || p.QrcodeKey != "qr-key-1" || p.CorpID != 1970324836974591 ||
		p.NickName != "小助手" || p.IconUrl != "http://wx.example/icon.png" ||
		p.Logo != "http://wx.example/logo.png" {
		t.Errorf("扫码字段解析错误: %+v", p)
	}

	// 畸形 JSON → error（严格语义，由调用方决定跳过）
	if _, err := (&CallBackMsg{Json: `{not-json`}).AsScanReturn(); err == nil {
		t.Error("畸形 JSON 应返回 error（严格语义）")
	}
}

// TestAsLoginSuccess 登录成功：corp_name 有值映射与畸形 JSON 严格报错。
func TestAsLoginSuccess(t *testing.T) {
	msg := &CallBackMsg{Json: `{
		"Corpid": 1970324836974591,
		"corp_name": "某某科技有限公司",
		"Vid": 1688851000000001,
		"nickname": "张三",
		"realname": "张三丰",
		"mobile": "13800138000",
		"avatar": "http://wx.example/avatar.png"
	}`}
	p, err := msg.AsLoginSuccess()
	if err != nil {
		t.Fatalf("正常 JSON 不应报错: %v", err)
	}
	if p.CorpID != 1970324836974591 || p.CorpName != "某某科技有限公司" ||
		p.Vid != 1688851000000001 || p.Nickname != "张三" || p.RealName != "张三丰" ||
		p.Mobile != "13800138000" || p.Avatar != "http://wx.example/avatar.png" {
		t.Errorf("登录成功字段解析错误: %+v", p)
	}

	if _, err := (&CallBackMsg{Json: `{"Corpid": ]`}).AsLoginSuccess(); err == nil {
		t.Error("畸形 JSON 应返回 error（严格语义）")
	}
}

// TestLoginDomainAccessorsStrict 登录域其余访问器：正常 JSON 无错、畸形 JSON 均报错。
func TestLoginDomainAccessorsStrict(t *testing.T) {
	ok := &CallBackMsg{Json: `{"Vid":1688851000000001,"Qrcode_key":"qr-1","Corpid":1970324836974591,"Nick_name":"n","Icon_url":"i","Logo":"l"}`}
	bad := &CallBackMsg{Json: `{bad`}

	cases := []struct {
		name string
		call func(*CallBackMsg) error
	}{
		{"AsCancelReturn", func(m *CallBackMsg) error { _, err := m.AsCancelReturn(); return err }},
		{"AsConfirmReturn", func(m *CallBackMsg) error { _, err := m.AsConfirmReturn(); return err }},
		{"AsNeedInputCode", func(m *CallBackMsg) error { _, err := m.AsNeedInputCode(); return err }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(ok); err != nil {
				t.Errorf("正常 JSON 不应报错: %v", err)
			}
			if err := tc.call(bad); err == nil {
				t.Error("畸形 JSON 应返回 error（严格语义）")
			}
		})
	}
}

// TestAsNewRoom 建群回调：嵌套 member_list 解析与畸形 JSON 严格报错。
func TestAsNewRoom(t *testing.T) {
	msg := &CallBackMsg{Json: `{
		"room_id": "R:10782666091770616",
		"create_user_id": 1688851000000001,
		"total": 3,
		"create_time": 1755600000,
		"notice_content": "群公告内容",
		"member_list": [
			{"acctid": "zhangsan", "uin": 1688851000000001, "nickname": "张三"},
			{"acctid": "lisi", "uin": "1688851000000002", "nickname": "李四"}
		]
	}`}
	p, err := msg.AsNewRoom()
	if err != nil {
		t.Fatalf("正常 JSON 不应报错: %v", err)
	}
	if p.RoomID != "R:10782666091770616" || p.CreateUserID != 1688851000000001 ||
		p.Total != 3 || p.CreateTime != 1755600000 || p.NoticeContent != "群公告内容" {
		t.Errorf("建群字段解析错误: %+v", p)
	}
	if len(p.MemberList) != 2 {
		t.Fatalf("MemberList 长度 = %d, want 2", len(p.MemberList))
	}
	if p.MemberList[0].Acctid != "zhangsan" || p.MemberList[1].Nickname != "李四" {
		t.Errorf("成员字段解析错误: %+v", p.MemberList)
	}

	if _, err := (&CallBackMsg{Json: `{bad`}).AsNewRoom(); err == nil {
		t.Error("畸形 JSON 应返回 error（严格语义）")
	}
}

// TestGroupMemberList 成员引用列表：acctid/uin 双字段、uin 数字/字符串双型、缺失空切片。
func TestGroupMemberList(t *testing.T) {
	msg := &CallBackMsg{Json: `{"member_list":[
		{"acctid":"zhangsan","uin":1688851000000001},
		{"acctid":"lisi","uin":"1688851000000002"},
		{"acctid":"wangwu"}
	]}`}
	refs := msg.GroupMemberList()
	if len(refs) != 3 {
		t.Fatalf("引用数 = %d, want 3", len(refs))
	}
	if refs[0].Acctid != "zhangsan" || refs[0].UIN != 1688851000000001 {
		t.Errorf("数字 uin 解析错误: %+v", refs[0])
	}
	// 字符串数字 uin → gjson Int 宽容转换
	if refs[1].UIN != 1688851000000002 {
		t.Errorf("字符串 uin 应宽容转换, got %+v", refs[1])
	}
	// 仅 acctid 的成员（部分协议版本）→ UIN 零值不报错
	if refs[2].Acctid != "wangwu" || refs[2].UIN != 0 {
		t.Errorf("仅 acctid 成员解析错误: %+v", refs[2])
	}

	// member_list 缺失 → 空切片（非 nil panic）
	if empty := (&CallBackMsg{Json: `{}`}).GroupMemberList(); len(empty) != 0 {
		t.Errorf("缺失 member_list 应返回空切片, got %d", len(empty))
	}
}

// TestAsGroupNameChangeFallback 群名三级回退：room_name → nickname → name。
func TestAsGroupNameChangeFallback(t *testing.T) {
	cases := []struct {
		name    string
		json    string
		want    string
		wantRID string
	}{
		{"room_name 优先", `{"room_id":"R:1","room_name":"甲群","nickname":"乙群","name":"丙群"}`, "甲群", "R:1"},
		{"回退 nickname", `{"room_id":"R:1","nickname":"乙群","name":"丙群"}`, "乙群", "R:1"},
		{"回退 name", `{"room_id":"R:1","name":"丙群"}`, "丙群", "R:1"},
		{"均缺失为空串", `{"room_id":"R:1"}`, "", "R:1"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := (&CallBackMsg{Json: tc.json}).AsGroupNameChange()
			if p.GroupName != tc.want {
				t.Errorf("GroupName = %q, want %q", p.GroupName, tc.want)
			}
			if p.RoomID != tc.wantRID {
				t.Errorf("RoomID = %q, want %q", p.RoomID, tc.wantRID)
			}
		})
	}
}

// TestRoomID 群事件回调群 ID 提取（含缺失空串）。
func TestRoomID(t *testing.T) {
	if got := (&CallBackMsg{Json: `{"room_id":"R:10782666091770616"}`}).RoomID(); got != "R:10782666091770616" {
		t.Errorf("RoomID = %q", got)
	}
	if got := (&CallBackMsg{Json: `{}`}).RoomID(); got != "" {
		t.Errorf("缺失应返回空串, got %q", got)
	}
}

// TestAsFriendChangeList 好友变更列表：user_id 字符串视图与嵌套 source_info.source_type。
func TestAsFriendChangeList(t *testing.T) {
	msg := &CallBackMsg{Json: `{"list":[{
		"status": 9,
		"user_id": 1688851000000001,
		"nickname": "客户A",
		"realname": "A",
		"avatar": "http://wx.example/a.png",
		"add_customer_time": 1755600000,
		"source_info": {"source_type": 1}
	}]}`}
	items := msg.AsFriendChangeList()
	if len(items) != 1 {
		t.Fatalf("条目数 = %d, want 1", len(items))
	}
	it := items[0]
	if it.Status != 9 || it.Nickname != "客户A" || it.RealName != "A" ||
		it.Avatar != "http://wx.example/a.png" || it.AddCustomerTime != 1755600000 {
		t.Errorf("好友变更字段解析错误: %+v", it)
	}
	// user_id 字符串视图：数字 → 十进制串
	if it.UserID != "1688851000000001" {
		t.Errorf("UserID = %q, want 1688851000000001", it.UserID)
	}
	if it.SourceType != 1 {
		t.Errorf("SourceType = %d, want 1", it.SourceType)
	}

	if empty := (&CallBackMsg{Json: `{}`}).AsFriendChangeList(); len(empty) != 0 {
		t.Errorf("缺失 list 应返回空切片, got %d", len(empty))
	}
}

// TestAsInnerLabelChangeList 内部标签变动列表。
func TestAsInnerLabelChangeList(t *testing.T) {
	msg := &CallBackMsg{Json: `{"list":[{
		"vid": 1688851000000001,
		"description": "高意向",
		"remark": "备注文本",
		"timestamp": 1755600000
	}],"seq": 7}`}
	items := msg.AsInnerLabelChangeList()
	if len(items) != 1 {
		t.Fatalf("条目数 = %d, want 1", len(items))
	}
	if items[0].Vid != 1688851000000001 || items[0].Description != "高意向" ||
		items[0].Remark != "备注文本" || items[0].Timestamp != 1755600000 {
		t.Errorf("标签变动字段解析错误: %+v", items[0])
	}

	if empty := (&CallBackMsg{Json: `{}`}).AsInnerLabelChangeList(); len(empty) != 0 {
		t.Errorf("缺失 list 应返回空切片, got %d", len(empty))
	}
}

// TestAsContactEventMsg 联系人事件提示（被删除/被拉黑通知）。
func TestAsContactEventMsg(t *testing.T) {
	if got := (&CallBackMsg{Json: `{"msg":"对方已把你删除"}`}).AsContactEventMsg(); got.Content != "对方已把你删除" {
		t.Errorf("Content = %q", got.Content)
	}
	if got := (&CallBackMsg{Json: `{}`}).AsContactEventMsg(); got.Content != "" {
		t.Errorf("缺失应返回空串, got %q", got.Content)
	}
}

// TestDisconnectVid 断连 vid 双型兼容与缺失语义。
func TestDisconnectVid(t *testing.T) {
	cases := []struct {
		name      string
		json      string
		wantVid   string
		wantFound bool
	}{
		{"number 型", `{"vid":1688851000000001}`, "1688851000000001", true},
		{"string 型", `{"vid":"1688851000000001"}`, "1688851000000001", true},
		{"缺失", `{"other":"x"}`, "", false},
		{"空 JSON", `{}`, "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			vid, found := (&CallBackMsg{Json: tc.json}).DisconnectVid()
			if vid != tc.wantVid || found != tc.wantFound {
				t.Errorf("DisconnectVid() = (%q, %v), want (%q, %v)", vid, found, tc.wantVid, tc.wantFound)
			}
		})
	}
}
