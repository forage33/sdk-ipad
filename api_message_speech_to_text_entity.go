// api_message_speech_to_text_entity.go SpeechToTextEntity 语音转文字。

package sdk_ipad

import "errors"

// TextEntityResp 语音转文字的返回结果。
type TextEntityResp struct {
	Data struct {
		Text string `json:"text"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// SpeechToTextEntity 语音转文字。
// 端点: POST /wxwork/SpeechToTextEntity
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - msgID: 语音消息 ID
//
// 返回: *TextEntityResp 包含转换后的文本
func (s *SdkIPad) SpeechToTextEntity(host, uuid string, msgID int) (*TextEntityResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid":  uuid,
		"msgid": msgID,
	}
	var entity TextEntityResp
	if err := s.httpClient.Post(host+"/wxwork/SpeechToTextEntity", reqParams, &entity); err != nil {
		return nil, err
	}
	return &entity, nil
}
