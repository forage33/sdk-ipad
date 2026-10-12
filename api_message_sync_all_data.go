// api_message_sync_all_data.go SyncAllData 同步消息记录。

package sdk_ipad

import "errors"

// SyncAllDataResp 同步消息记录的返回结果。
type SyncAllDataResp struct {
	Data struct {
		IsSelect int    `json:"is_select"`
		Seq      int    `json:"seq"`
		Content  string `json:"content"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SyncAllData 同步消息记录。
// 端点: POST /wxwork/SyncAllData
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - limit: 单次拉取消息条数
//   - seq: 消息序列号（从上次同步的位置继续）
//
// 返回: *SyncAllDataResp 包含同步到的消息内容和新的 seq
func (s *SdkIPad) SyncAllData(host, uuid string, limit, seq int) (*SyncAllDataResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"limit": limit,
		"seq":   seq,
	}
	var resp SyncAllDataResp
	if err := s.httpClient.Post(host+"/wxwork/SyncAllData", reqParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
