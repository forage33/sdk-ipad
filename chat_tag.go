// chat_tag.go 聊天标签操作相关接口。
// 提供聊天标签的增删改等功能。

package sdk_ipad

import "errors"

// ---------------------------------------------------------------------------
// 返回数据结构
// ---------------------------------------------------------------------------

// SessionTagInfo 聊天标签信息。
type SessionTagInfo struct {
	Flag       int    `json:"flag"`
	PrevTagID  int64  `json:"prev_tagid"`
	AfterTagID int64  `json:"after_tagid"`
	TagID      int64  `json:"tagid"`
	Name       string `json:"name"`
	Cmd        int    `json:"cmd"`
	Order      int    `json:"order"`
}

// ModSessionTagResp 增删改聊天标签的返回结果。
type ModSessionTagResp struct {
	Data struct {
		GlobalVer int64            `json:"global_ver"`
		SyncKey   int64            `json:"synckey"`
		TagInfo   []SessionTagInfo `json:"taginfo"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// ---------------------------------------------------------------------------
// 调用接口
// ---------------------------------------------------------------------------

// ModSessionTag 增删改聊天标签。
// 端点: POST /wxwork/ModSessionTag
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - cmd: 操作类型（1=新增, 2=修改, 3=删除）
//   - tagName: 标签名称（删除时可为空）
//   - tagID: 标签 ID（新增时传 0）
//   - globalVer: 全局版本号（用于并发控制）
//
// 返回: *ModSessionTagResp 包含操作后的标签信息和新版本号
func (s *SdkIPad) ModSessionTag(host, uuid string, cmd int, tagName string, tagID int64, globalVer int) (*ModSessionTagResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":      uuid,
		"cmd":       cmd,
		"tagName":   tagName,
		"tagid":     tagID,
		"globalver": globalVer,
	}
	var resp ModSessionTagResp
	if err := s.httpClient.Post(host+"/wxwork/ModSessionTag", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ---------------------------------------------------------------------------
// 待补全接口（参数待从文档确认后完善）
// ---------------------------------------------------------------------------

// TODO: GetSessionTagList 获取聊天标签列表
// 预估路径: /wxwork/GetSessionTagList
// 预估参数: uuid, globalver(int)

// TODO: GetSessionTagMembers 获取标签下的联系人
// 预估路径: /wxwork/GetSessionTagMembers
// 预估参数: uuid, tagid(int64)

// TODO: ModSessionTagMembers 新增/移除联系人到聊天标签
// 预估路径: /wxwork/ModSessionTagMembers
// 预估参数: uuid, tagid(int64), add_userid([]int64), del_userid([]int64)
