package sdk_ipad

// SdkInterface 定义业务侧使用的 SDK 方法集合（ipad-and-customer 服务层、
// admin 群管理 Handler 等调用方均以此接口类型持有 SDK 实例）。
// SdkIPad 已实现所有这些方法，Go 隐式接口自动满足。
type SdkInterface interface {
	// 消息撤回
	RevokeMsg(host, uuid string, msgID, roomID int64) (*RevokeResp, error)
	// 消息发送
	SendTextMsg(host, uuid string, sendUserID int64, content string, isRoom bool) (*SendTextResp, error)
	SendTextAtMsg(host, uuid string, sendUserId int64, atIds []int64, content string) (*SendTextAtResp, error)
	SendCDNImage(host, uuid string, sendUserId int64, isRoom bool, cdnKey, aesKey, md5 string, fileSize int64) (*SendCDNImageResp, error)
	SendCDNFile(host, uuid string, sendUserId int64, isRoom bool, cdnKey, aesKey, md5, fileName string, fileSize int64) (*SendCDNFileResp, error)
	SendCDNVideoMsg(host, uuid string, sendUserId int64, isRoom bool, cdnKey, aesKey, md5 string, videoDuration int, fileSize, videoImgSize int64, videoWidth, videoHeight int) (*SendCDNVideoResp, error)
	SendCDNVoiceMsg(host, uuid string, sendUserId int64, isRoom bool, cdnKey, aesKey, md5 string, voiceTime int, fileSize int64) (*SendCDNVoiceResp, error)
	SendLinkMsg(host, uuid string, sendUserId int64, isRoom bool, url, title, content, imgurl string) (*SendLinkMsgResp, error)
	SendLocationMsg(host, uuid string, sendUserId int64, isRoom bool, longitude, latitude float64, address, detailedAddress string) (*SendLocationResp, error)
	SendAppMsg(host, uuid string, sendUserId int64, isRoom bool, desc, appName, title, weappIconUrl, pagepath, username, appid, cdnKey, md5, aesKey string, fileSize int64) (*SendAppMsgResp, error)
	MarkAsRead(host, uuid string, sendUserId int64, isRoom bool) (*MarkAsReadResp, error)
	SpeechToTextEntity(host, uuid string, msgID int) (*TextEntityResp, error)

	// CDN 操作
	CdnUploadImgLinkIntegration(host, uuid, url string) (*UploadImgResp, error)
	UploadCdnLinkFileIntegration(host, uuid, url, fileName string) (*UploadFileResp, error)
	UploadCdnVideoLinkIntegration(host, uuid, url, fileName, videoImgUrl string) (*UploadVideoResp, error)
	DownloadWeChatFile(host, uuid, url, aesKey, authKey, fileName string, size int64) (*DownloadWechatFileResp, error)
	DownloadFile(host, uuid, fileId, aesKey, fileName string, fileType int, size int64) (*DownloadFileResp, error)

	// 联系人
	GetUserInfoByVids(host, uuid string, vIDs []int64) (*UserInfoResp, error)
	GetInnerContacts(host, uuid string, limit int, strSeq string) (*InnerResp, error)
	SearchContactByPhone(host, uuid, phoneNum string) (*SearchPhoneResp, error)
	AddSearch(host, uuid string, vid int64, optionID, phone, content, ticket string) (*AddSearchResp, error)
	AddShareUser(host, uuid string, vid int64, content string) (*AddShareUserResp, error)
	AddCard(host, uuid string, vid int64, recommendId int64, roomID int64, content string) (*AddCardResp, error)
	GetExternalContacts(host, uuid string, limit, seq int) (*ExternalResp, error)

	// 好友管理
	AddSearchWxwork(host, uuid string, vid, corpid int64, ticket, content string) (*CommonContactResp, error)
	DeleteContact(host, uuid string, vid int64) (*CommonContactResp, error)
	AgreeUser(host, uuid string, vid, corpid int64) (*CommonContactResp, error)
	GetApplyUserList(host, uuid string, limit, seq int) (*ApplyUserListResp, error)
	AddWxUser(host, uuid string, vid int64, content string) (*CommonContactResp, error)
	OperCustomer(host, uuid, thisVID, recommendVID, sendVID string) (*OperCustomerResp, error)

	// 群组
	GetRoomDesUserList(host, uuid string, roomID int64) (*GetRoomDesUserListResp, error)
	GetSessionRoomList(host, uuid string, limit, starIndex int) (*RoomSessionResp, error)
	GetRoomUserList(host, uuid string, roomID int64) (*RoomUserListResp, error)
	GetRoomInfo(host, uuid string, roomID int64) (*GetRoomInfoResp, error)
	WxRoomHeader(host, uuid string, roomID int64) (*WxRoomHeaderResp, error)
	CreateRoomWx(host, uuid, roomName string, vids []int64) (*CreateRoomResp, error)
	CreateRoomNei(host, uuid, roomName string, vids []int64) (*CreateRoomResp, error)
	UpdateRoomName(host, uuid string, roomID int64, roomName string) (*UpdateRoomNameResp, error)
	SendNotice(host, uuid string, roomID int64, msg string) (*CommonGroupResp, error)
	WxRoomInvite(host, uuid string, roomID int64) (*WxRoomInviteResp, error)
	InvitationToRoom(host, uuid string, roomID int64, vids []int64) (*CommonGroupResp, error)
	InvitationToRoomLink(host, uuid string, roomID int64, vids []int64) (*CommonGroupResp, error)
	DelRoomUsers(host, uuid string, roomID int64, vids []int64) (*CommonGroupResp, error)
	AddRoomFriends(host, uuid string, roomID int64, corpId int, vid int64, content string) (*CommonGroupResp, error)
	TransferChatroomOwner(host, uuid string, roomID, vid int64) (*CommonGroupResp, error)
	DissolutionRoom(host, uuid string, roomID int64) (*CommonGroupResp, error)
	OutRoomReq(host, uuid string, roomID int64) (*OutRoomReqResp, error)
	GetWelcomeMsgList(host, uuid, offset string, limit int) (*GetWelcomeMsgListResp, error)
	AddWelcomeMsg(host, uuid string, list []string, url, title, content, headImg string) (*AddWelcomeMsgResp, error)
	SetWelcomeMsg(host, uuid string, indexId, msgType int, roomID int64) (*CommonGroupResp, error)
	DelWelcomeMsg(host, uuid string, roomID int64) (*CommonGroupResp, error)
	AgreeCardAddRoom(host, uuid, code string) (*AgreeCardAddRoomResp, error)

	// 名片/消息
	SendBusinessCardMsg(host, uuid string, sendUserId int64, isRoom bool, headImg string, username int64, nickname, enterpriseName string) (*SendBusinessCardResp, error)

	// 朋友圈
	GetSnsList(host, uuid string, seq int64, limit int) (*SnsListResp, error)
	GetSnsInfo(host, uuid string, sid int64) (*SnsInfoResp, error)
	PostPyq(host, uuid, content string, fileContent []SnsMedia, vids []int64, linkInfo *SnsLinkInfo) (*PostPyqResp, error)
	DeleteSns(host, uuid string, sid int64) (*CommonSnsResp, error)
	CommentSns(host, uuid string, sid int64, comment string, refCommentID int64) (*CommentSnsResp, error)
	DeleteSnsComment(host, uuid string, sid int64, commentID int64) (*CommonSnsResp, error)
	LikeOrUnLikeSns(host, uuid string, sid int64, operation int) (*CommonSnsResp, error)

	// 登录
	GetQrCode(host, uuid string) (*GetQRCodeResp, error)
	CheckCode(host, uuid, qrCodeKey, code string) (*InputCodeResp, error)
	AutomaticLogin(host, uuid string) (*AutoLoginResp, error)
	LoginOut(host, uuid string) (*LoginOutResp, error)
	SecondaryValidation(host, uuid string) (*SecondaryValidation, error)

	// 账号管理
	GetBanDetail(host, uuid, url string) (*BanResp, error)
	AddImage(host, uuid, url string) (*SetImageResp, error)
	UpdateUserInfo(host, uuid, emile, name, englishName, phone, position, mobile string, gender int) (*CommonAccountResp, error)
	GetThisQRCode(host, uuid string) (*GetQrCodeResp, error)
	AutomaticConsent(host, uuid string, state int) (*CommonAccountResp, error)
	GetProfile(host, uuid string) (*ProfileResp, error)

	// 企微初始化与运行状态
	InItWXWork(host, vid, proxyType string, ip string, port int, userName, password, deviceType string, proxySituation int) (*InitResp, error)
	SetCallBackURL(host, uuid string, config CallbackConfig) (*CallbackResp, error)
	CloseConnect(host, uuid string) (*CloseConnResp, error)
	GetRunClient(host string) (*ClientResponse, error)
	GetRunClientByUuid(host, uuid string) (*ClientRespByUUID, error)
	GetRunningByVid(host, vid string) (*UserData, bool)
}
