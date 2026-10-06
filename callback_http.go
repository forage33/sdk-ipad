// callback_http.go 回调接收的传输层适配。
//
// 提供与 Web 框架解耦的标准 net/http handler：接收协议服务 POST 推送的
// 回调信封（CallBackMsg），提交给调用方的分发函数，失败时异步重试，
// 无论处理结果如何始终向协议服务应答 200（协议服务以非 200 视为推送失败
// 会重复推送，业务侧以异步重试保证可靠，应答保持即时 200 防止推送积压）。
//
// gin/echo 等框架通过各自的包装器挂载（如 gin.WrapF）。

package sdk_ipad

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// CallbackHandlerOptions 回调 HTTP handler 的可选配置。
type CallbackHandlerOptions struct {
	// MaxRetries handle 失败后的异步重试次数，0 表示不重试（默认重试建议 3）
	MaxRetries int
	// BaseBackoff 重试初始退避时长，0 时使用默认 200ms；每次失败后翻倍
	BaseBackoff time.Duration
	// Logger 重试过程日志（zap Field 风格）；nil 时静默
	Logger *zap.Logger
}

// DefaultCallbackHandlerOptions 返回与历史业务行为一致的默认配置
// （异步重试 3 次、初始退避 200ms）。
func DefaultCallbackHandlerOptions() CallbackHandlerOptions {
	return CallbackHandlerOptions{
		MaxRetries:  3,
		BaseBackoff: 200 * time.Millisecond,
	}
}

// CallbackHTTPHandler 返回接收协议服务回调推送的标准 http.HandlerFunc。
//
// handle 为回调分发入口（通常为业务侧 Adapter 的 HandleCallback 方法值）；
// handle 返回错误时按 opts 异步重试（指数退避 + 随机抖动），重试日志经
// opts.Logger 输出（保持与历史业务实现相同的日志文案）。
//
// 应答协议：
//   - 信封解析失败：200 + {"errorcode":-1,"errmsg":"解析参数错误"}
//   - 其余情况（含 handle 失败后转异步重试）：200 + {"errorcode":0,"errmsg":"ok"}
func CallbackHTTPHandler(handle func(*CallBackMsg) error, opts CallbackHandlerOptions) http.HandlerFunc {
	if opts.BaseBackoff <= 0 {
		opts.BaseBackoff = 200 * time.Millisecond
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var msg CallBackMsg
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			writeCallbackResp(w, -1, "解析参数错误")
			return
		}

		if err := handle(&msg); err != nil {
			// 首次分发失败记录错误详情（含 Submit 超时等原因，排障关键信息）
			if opts.Logger != nil {
				if opts.MaxRetries > 0 {
					opts.Logger.Error("HandleCallback failed, will retry",
						zap.Int64("type", int64(msg.Type)),
						zap.String("wx_uuid", msg.UUID),
						zap.Error(err))
				} else {
					opts.Logger.Error("HandleCallback failed",
						zap.Int64("type", int64(msg.Type)),
						zap.String("wx_uuid", msg.UUID),
						zap.Error(err))
				}
			}
			if opts.MaxRetries > 0 {
				// 异步重试：指数退避+随机抖动最多 MaxRetries 次，
				// 避免分发队列瞬时拥塞时回调静默丢失（与历史业务实现一致）
				go retryCallback(handle, &msg, opts)
			}
		}

		writeCallbackResp(w, 0, "ok")
	}
}

// retryCallback 异步重试：指数退避+随机抖动。
func retryCallback(handle func(*CallBackMsg) error, msg *CallBackMsg, opts CallbackHandlerOptions) {
	backoff := opts.BaseBackoff
	for attempt := 1; attempt <= opts.MaxRetries; attempt++ {
		jitter := time.Duration(rand.Int63n(int64(backoff / 2)))
		time.Sleep(backoff + jitter)
		if retryErr := handle(msg); retryErr == nil {
			if opts.Logger != nil {
				opts.Logger.Info("HandleCallback retry success",
					zap.Int64("type", int64(msg.Type)),
					zap.String("wx_uuid", msg.UUID),
					zap.Int("attempt", attempt))
			}
			return
		}
		backoff *= 2
	}
	if opts.Logger != nil {
		opts.Logger.Error("HandleCallback retry exhausted",
			zap.Int64("type", int64(msg.Type)),
			zap.String("wx_uuid", msg.UUID))
	}
}

// writeCallbackResp 应答协议服务（始终 HTTP 200）。
func writeCallbackResp(w http.ResponseWriter, errcode int, errmsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"errorcode": errcode,
		"errmsg":    errmsg,
	})
}
