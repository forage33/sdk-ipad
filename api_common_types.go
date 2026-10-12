// api_common_types.go 跨接口共享的数据结构（被多个接口方法引用；含暂无引用的历史类型）。

package sdk_ipad

// UserData 运行中实例的用户数据。
type UserData struct {
	IsLogin    bool   `json:"isLogin"`
	ClientID   string `json:"clientId"`
	UpdateTime int64  `json:"updateTime"`
	UserID     int64  `json:"userid"`
	UUID       string `json:"uuid"`
	ReqURL     string `json:"requrl"`
	Object     User   `json:"object"`
}

// User 企业微信用户基本信息。
type User struct {
	UnionID      string `json:"unionid"`
	CreateTime   int64  `json:"create_time"`
	Sex          int    `json:"sex"`
	Mobile       string `json:"mobile"`
	CorpFullName string `json:"Corp_full_name"`
	AcctID       string `json:"acctid"`
	SCorpID      string `json:"scorp_id"`
	Avatar       string `json:"avatar"`
	CorpName     string `json:"corp_name"`
	EnglishName  string `json:"english_name"`
	RealName     string `json:"realname"`
	UserID       int64  `json:"user_id"`
	Nickname     string `json:"nickname"`
	Position     string `json:"position"`
	CorpDesc     string `json:"Corp_desc"`
	CorpID       int64  `json:"corp_id"`
}

// CommonAccountResp 账号操作通用返回结果（data 为 null 或简单数据）。
type CommonAccountResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// LabelItem 标签详情。
type LabelItem struct {
	Op             int    `json:"op"`
	BDeleted       int    `json:"bDeleted"`
	CreateTime     int64  `json:"create_time"`
	LabelGroupID   int64  `json:"label_groupid"`
	LabelType      int    `json:"label_type"`
	SourceAppID    int    `json:"source_appid"`
	BusinessType   int    `json:"business_type"`
	Name           string `json:"name"`
	DataType       int    `json:"data_type"`
	ID             int64  `json:"id"`
	ServiceGroupID int64  `json:"service_groupid"`
	Seq            int64  `json:"seq,omitempty"`
	NewOrder       int    `json:"new_order"`
	Order          int    `json:"order"`
}

// UserInfo 搜索结果中的用户信息。
type UserInfo struct {
	HeadImg string `json:"headImg"`
	Ticket  string `json:"ticket"`
	UserId  int64  `json:"user_id"`
	Sex     int    `json:"sex"`
	Name    string `json:"name"`
	State   string `json:"state"`
	CorpId  int64  `json:"corp_id,omitempty"`
	Openid  string `json:"openid,omitempty"`
}

// CommonContactResp 联系人操作通用返回结果（data 为 null 或简单数据）。
type CommonContactResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// SnsMedia 朋友圈媒体文件信息。
type SnsMedia struct {
	Size      int    `json:"size"`
	Width     int    `json:"width"`
	VideoLen  int    `json:"video_len"`
	StoreType int    `json:"store_type"`
	Type      int    `json:"type"`   // 1=图片, 2=视频, 3=视频封面
	FileID    string `json:"fileid"` // CDN 文件 ID
	Md5       string `json:"md5"`
	Height    int    `json:"height"`
}

// SnsComment 朋友圈评论信息。
type SnsComment struct {
	FromID       int64  `json:"from_id"`       // 评论人 VID
	Flag         int    `json:"flag"`          // 标记
	CommentID    int64  `json:"commentid"`     // 评论 ID
	Comment      string `json:"comment"`       // 评论内容
	Time         int64  `json:"time"`          // 评论时间（Unix 时间戳）
	PostID       string `json:"postid"`        // 朋友圈 Post ID
	RefCommentID int64  `json:"ref_commentid"` // 引用评论 ID（0 表示直接评论）
	SID          int64  `json:"sid"`           // 朋友圈 SID
}

// SnsLinkInfo 朋友圈链接信息。
type SnsLinkInfo struct {
	Title      string `json:"title"`
	ContentUrl string `json:"content_url"`
}

// SnsLikeItem 朋友圈点赞信息。
type SnsLikeItem struct {
	Xid    int64  `json:"xid"`
	Time   int64  `json:"time"`
	PostID string `json:"postid"`
	SID    int64  `json:"sid"`
}

// SnsPoiInfo 朋友圈位置信息。
type SnsPoiInfo struct {
	City               string `json:"city"`
	Latitude           string `json:"latitude"`
	Longitude          string `json:"longitude"`
	PoiAddress         string `json:"poiAddress"`
	PoiClassifyId      string `json:"poiClassifyId"`
	PoiClassifyType    string `json:"poiClassifyType"`
	PoiClickableStatus string `json:"poiClickableStatus"`
	PoiName            string `json:"poiName"`
	PoiScale           string `json:"poiScale"`
}

// SnsItem 朋友圈条目详情。
type SnsItem struct {
	Flag              int           `json:"flag"`
	TaskSID           int64         `json:"task_sid"`
	PostID            string        `json:"postid"`
	Media             []SnsMedia    `json:"media"`
	Type              int           `json:"type"`
	Content           string        `json:"content"`
	SID               int64         `json:"sid"`
	UpdateTime        int64         `json:"update_time"`
	Xid               []int64       `json:"xid"`
	AuthorVid         int64         `json:"author_vid"`
	DeleteCommentList []SnsComment  `json:"delete_comment_list"`
	CommentList       []SnsComment  `json:"comment_list"`
	LinkInfo          SnsLinkInfo   `json:"link_info"`
	VisibleType       int           `json:"visible_type"`
	Time              int64         `json:"time"`
	Seq               int64         `json:"seq"`
	SnsLikeInfo       []SnsLikeItem `json:"SnsLikeInfo"`
	PoiInfo           *SnsPoiInfo   `json:"poiInfo,omitempty"`
}

// CommonSnsResp 朋友圈操作通用返回结果。
type CommonSnsResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// CreateRoomResp 创建群聊的返回结果。
type CreateRoomResp struct {
	Data struct {
		Roomname   string      `json:"roomname"`
		CreateTime int64       `json:"createTime"`
		Createid   int64       `json:"createid"`
		Roomid     interface{} `json:"roomid"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// CommonGroupResp 群操作通用返回结果。
type CommonGroupResp struct {
	Data    interface{} `json:"data"`
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
}

// RoomMember 群成员详细信息。
type RoomMember struct {
	Unionid      string `json:"unionid"`
	CreateTime   int64  `json:"create_time"`
	Sex          int    `json:"sex"`
	Mobile       string `json:"mobile"`
	Acctid       string `json:"acctid"`
	JoinScene    int    `json:"join_scene"`
	Avatar       string `json:"avatar"`
	EnglishName  string `json:"english_name"`
	Realname     string `json:"realname"`
	RoomNotes    string `json:"room_notes"`
	Jointime     int64  `json:"jointime"`
	Nickname     string `json:"nickname"`
	RoomNickname string `json:"room_nickname"`
	Position     string `json:"position"`
	Uin          int64  `json:"uin"`
	InviteUserId int64  `json:"invite_user_id"`
	CorpId       int64  `json:"corp_id"`
}

// VidOpenIdPair 企业用户 id 与开放平台 id（external_userid）的对应关系。
// 用于 Vid2UserId / UserId2Vid / Vid2OpenUserTripartiteId 的返回数据。
type VidOpenIdPair struct {
	UserID int64  `json:"user_id"`
	OpenID string `json:"openid"`
}

// RoomChatIdPair 群 id 与开放平台 chatid 的对应关系。
// 用于 RoomIdToChatId / ChatIdToRoomId 的返回数据。
type RoomChatIdPair struct {
	RoomID int64  `json:"room_id"`
	ChatID string `json:"chatid"`
}
