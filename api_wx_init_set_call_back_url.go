// api_wx_init_set_call_back_url.go SetCallBackURL 设置协议服务的回调推送地址。

package sdk_ipad

import "errors"

// 回调类型常量
const (
	// CallbackTypeHTTP HTTP 回调
	CallbackTypeHTTP = "HTTP"
	// CallbackTypeRabbitMQ RABBITMQ 回调
	CallbackTypeRabbitMQ = "RABBITMQ"
)

// CallbackConfig 设置回调地址的请求参数。
type CallbackConfig struct {
	// CallbackType 消息回调类型: HTTP 或 RABBITMQ
	CallbackType string
	// URL HTTP 回调地址
	URL string
	// MqExchange RABBITMQ 交换机（direct 类型）
	MqExchange string
	// MqRoutingKey RABBITMQ 路由键
	MqRoutingKey string
	// ExtraContent 自定义回调额外内容，可为空
	ExtraContent string
}

// CallbackData 设置回调地址返回的数据部分。
type CallbackData struct {
	CallbackType string `json:"callbackType"`
	URL          string `json:"url"`
	MqExchange   string `json:"mqExchange"`
	MqRoutingKey string `json:"mqRoutingKey"`
	ExtraContent string `json:"extraContent"`
	Msg          string `json:"msg"`
	StatusCode   int    `json:"statusCode"`
}

// CallbackResp 设置回调地址的返回结果。
type CallbackResp struct {
	Data    CallbackData `json:"data"`
	ErrCode int          `json:"errcode"`
	ErrMsg  string       `json:"errmsg"`
}

// SetCallBackURL 设置协议服务的回调推送地址。
// 端点: POST /wxwork/SetCallbackUrl
//
// 参数:
//   - host: 协议服务地址
//   - uuid: 实例唯一标识
//   - config: 回调配置，CallbackType 为空时默认 HTTP 回调
//
// 返回: *CallbackResp 设置结果
func (s *SdkIPad) SetCallBackURL(host, uuid string, config CallbackConfig) (*CallbackResp, error) {
	if uuid == "" {
		return nil, errors.New("uuid not set,please set uuid")
	}
	// 请求参数
	reqParams := map[string]interface{}{
		"uuid": uuid,
	}
	switch config.CallbackType {
	case CallbackTypeRabbitMQ:
		reqParams["callbackType"] = CallbackTypeRabbitMQ
		reqParams["mqExchange"] = config.MqExchange
		reqParams["mqRoutingKey"] = config.MqRoutingKey
		if config.ExtraContent != "" {
			reqParams["extraContent"] = config.ExtraContent
		}
	default: // HTTP
		reqParams["callbackType"] = CallbackTypeHTTP
		reqParams["url"] = config.URL
	}

	var callbackResp CallbackResp

	if err := s.httpClient.Post(host+"/wxwork/SetCallbackUrl", reqParams, &callbackResp); err != nil {
		return nil, err
	}

	return &callbackResp, nil
}
