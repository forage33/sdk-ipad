package sdk_ipad

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// callback_http_test.go 回调 HTTP 传输层（callback_http.go）的行为契约测试。
//
// 应答协议铁律：
//   - 无论 handle 成功与否，均立即应答 HTTP 200（协议服务以非 200 视为
//     推送失败会重复推送，业务侧以异步重试保证可靠）
//   - 信封解析失败：200 + {"errorcode":-1,"errmsg":"解析参数错误"}
//   - 其余情况（含 handle 失败转异步重试）：200 + {"errorcode":0,"errmsg":"ok"}
//
// 重试语义：handle 失败后按指数退避 + 随机抖动异步重试 MaxRetries 次，
// 不阻塞应答；MaxRetries=0 时不重试。

// newCallbackRequest 构造一次回调推送请求（协议服务信封格式）。
func newCallbackRequest(body string) *http.Request {
	return httptest.NewRequest(http.MethodPost, "/callback", strings.NewReader(body))
}

// decodeCbResp 解析 handler 应答体为 map。
func decodeCbResp(t *testing.T, w *httptest.ResponseRecorder) map[string]interface{} {
	t.Helper()
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("应答体不是合法 JSON: %v, body=%s", err, w.Body.String())
	}
	return resp
}

// waitForTrue 轮询等待条件成立（异步重试断言用）。
func waitForTrue(t *testing.T, cond func() bool, timeout time.Duration) bool {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if cond() {
			return true
		}
		time.Sleep(2 * time.Millisecond)
	}
	return false
}

// TestCallbackHTTPHandlerSuccess handle 成功的完整应答流程。
func TestCallbackHTTPHandlerSuccess(t *testing.T) {
	var got *CallBackMsg
	handler := CallbackHTTPHandler(func(msg *CallBackMsg) error {
		got = msg
		return nil
	}, DefaultCallbackHandlerOptions())

	body := `{"type":102000,"uuid":"uuid-1","json":"{\"content\":\"hi\"}"}`
	w := httptest.NewRecorder()
	handler(w, newCallbackRequest(body))

	if w.Code != http.StatusOK {
		t.Fatalf("状态码 = %d, want 200", w.Code)
	}
	resp := decodeCbResp(t, w)
	if resp["errorcode"] != float64(0) || resp["errmsg"] != "ok" {
		t.Errorf("应答体错误: %v", resp)
	}
	if got == nil {
		t.Fatal("handle 未被调用")
	}
	if got.Type != 102000 || got.UUID != "uuid-1" {
		t.Errorf("信封字段错误: type=%d uuid=%q", got.Type, got.UUID)
	}
	if got.Envelope().MsgType != 0 { // json 字段原样透传，由访问器解析
		t.Errorf("Json 字段未透传: %q", got.Json)
	}
}

// TestCallbackHTTPHandlerInvalidJSON 信封解析失败：仍应答 200 + errorcode=-1，且不调用 handle。
func TestCallbackHTTPHandlerInvalidJSON(t *testing.T) {
	called := false
	handler := CallbackHTTPHandler(func(msg *CallBackMsg) error {
		called = true
		return nil
	}, DefaultCallbackHandlerOptions())

	w := httptest.NewRecorder()
	handler(w, newCallbackRequest(`{bad`))

	if w.Code != http.StatusOK {
		t.Fatalf("状态码 = %d, want 200（解析失败也须应答 200 防推送积压）", w.Code)
	}
	resp := decodeCbResp(t, w)
	if resp["errorcode"] != float64(-1) || resp["errmsg"] != "解析参数错误" {
		t.Errorf("应答体错误: %v", resp)
	}
	if called {
		t.Error("解析失败时不应调用 handle")
	}
}

// TestCallbackHTTPHandlerRetrySuccess handle 首次失败后异步重试成功，应答即时 200。
func TestCallbackHTTPHandlerRetrySuccess(t *testing.T) {
	var calls int32
	handler := CallbackHTTPHandler(func(msg *CallBackMsg) error {
		if atomic.AddInt32(&calls, 1) == 1 {
			return errors.New("首次分发失败（模拟瞬时拥塞）")
		}
		return nil
	}, CallbackHandlerOptions{MaxRetries: 3, BaseBackoff: time.Millisecond})

	w := httptest.NewRecorder()
	handler(w, newCallbackRequest(`{"type":104001,"uuid":"uuid-2","json":"{}"}`))

	// 应答不等待重试：立即 200
	if w.Code != http.StatusOK {
		t.Fatalf("状态码 = %d, want 200", w.Code)
	}
	if resp := decodeCbResp(t, w); resp["errorcode"] != float64(0) {
		t.Errorf("应答体错误: %v", resp)
	}

	// 异步重试后成功（总调用 2 次）
	if !waitForTrue(t, func() bool { return atomic.LoadInt32(&calls) >= 2 }, 2*time.Second) {
		t.Fatalf("重试后仍未再次调用 handle, calls=%d", atomic.LoadInt32(&calls))
	}
}

// TestCallbackHTTPHandlerRetryExhausted handle 始终失败：重试 MaxRetries 次后放弃
// （总调用 = 1 + MaxRetries），应答仍为 200。
func TestCallbackHTTPHandlerRetryExhausted(t *testing.T) {
	var calls int32
	handler := CallbackHTTPHandler(func(msg *CallBackMsg) error {
		atomic.AddInt32(&calls, 1)
		return errors.New("持续失败")
	}, CallbackHandlerOptions{MaxRetries: 2, BaseBackoff: time.Millisecond})

	w := httptest.NewRecorder()
	handler(w, newCallbackRequest(`{"type":115005,"uuid":"uuid-3","json":"{}"}`))
	if w.Code != http.StatusOK {
		t.Fatalf("状态码 = %d, want 200", w.Code)
	}

	// 等待重试耗尽：1 次首发 + 2 次重试 = 3
	if !waitForTrue(t, func() bool { return atomic.LoadInt32(&calls) == 3 }, 2*time.Second) {
		t.Fatalf("重试次数不符, calls=%d, want 3", atomic.LoadInt32(&calls))
	}
	// 再多等一拍确认无第 4 次（不无限重试）
	time.Sleep(50 * time.Millisecond)
	if got := atomic.LoadInt32(&calls); got != 3 {
		t.Errorf("重试次数超限, calls=%d, want 3", got)
	}
}

// TestCallbackHTTPHandlerNoRetry MaxRetries=0：失败不重试（仅 1 次调用），应答 200。
func TestCallbackHTTPHandlerNoRetry(t *testing.T) {
	var calls int32
	handler := CallbackHTTPHandler(func(msg *CallBackMsg) error {
		atomic.AddInt32(&calls, 1)
		return errors.New("失败")
	}, CallbackHandlerOptions{MaxRetries: 0})

	w := httptest.NewRecorder()
	handler(w, newCallbackRequest(`{"type":102000,"uuid":"uuid-4","json":"{}"}`))
	if w.Code != http.StatusOK {
		t.Fatalf("状态码 = %d, want 200", w.Code)
	}
	if resp := decodeCbResp(t, w); resp["errorcode"] != float64(0) {
		t.Errorf("应答体错误: %v", resp)
	}

	time.Sleep(50 * time.Millisecond)
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Errorf("MaxRetries=0 不应重试, calls=%d, want 1", got)
	}
}

// TestDefaultCallbackHandlerOptions 默认配置与历史业务行为一致（重试 3 次、退避 200ms）。
func TestDefaultCallbackHandlerOptions(t *testing.T) {
	opts := DefaultCallbackHandlerOptions()
	if opts.MaxRetries != 3 {
		t.Errorf("MaxRetries = %d, want 3", opts.MaxRetries)
	}
	if opts.BaseBackoff != 200*time.Millisecond {
		t.Errorf("BaseBackoff = %v, want 200ms", opts.BaseBackoff)
	}
}
