// http_client.go HTTP 客户端封装。
// 提供 SDK 内部使用的 HTTP POST 请求能力，支持 JSON 序列化、自定义请求头和网络错误自动重试。

package sdk_ipad

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"time"
)

// optimizedTransport 全局共享的 HTTP Transport，针对多协议服务器、高并发场景调优。
//
// 默认 http.DefaultTransport.MaxIdleConnsPerHost = 2，无法满足单台协议服务器
// 承载 500 个号的并发需求，导致大量连接反复建立/关闭，增加延迟和 CPU 开销。
//
// 配置依据（每台协议服务器 ~500 号，共 ~20 台）：
//   - MaxIdleConnsPerHost: 100  — 覆盖 20% 账号同时活跃的峰值，避免 TCP+TLS 重复握手
//   - MaxConnsPerHost:     200  — 安全上限，防止单 host 耗尽进程 fd
//   - MaxIdleConns:       2000  — 20 hosts × 100 空闲连接
var optimizedTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConns:          2000,
	MaxIdleConnsPerHost:   100,
	MaxConnsPerHost:       200,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

// NewOptimizedHTTPClient 返回针对高并发场景调优的 http.Client。
//
// 使用全局共享的 optimizedTransport，多个 SdkIPad 实例复用同一连接池。
func NewOptimizedHTTPClient() *http.Client {
	return &http.Client{
		Timeout:   30 * time.Second,
		Transport: optimizedTransport,
	}
}

// HTTPClient HTTP 客户端接口，定义 SDK 内部所有 HTTP 请求的统一规范。
type HTTPClient interface {
	// Post 发送 JSON POST 请求并将响应反序列化到 target。
	Post(url string, params interface{}, target interface{}) error
	// PostWithHeaders 发送带自定义请求头的 JSON POST 请求。
	PostWithHeaders(url string, params interface{}, headers map[string]string, target interface{}) error
	// PostMultipart 发送 multipart/form-data POST 请求（用于文件上传）。
	PostMultipart(url string, fields map[string]string, fileField, fileName string, fileData io.Reader, target interface{}) error
}

// DefaultHTTPClient 默认 HTTP 客户端实现，内部封装 net/http.Client。
type DefaultHTTPClient struct {
	client *http.Client
}

// NewDefaultHTTPClient 创建默认 HTTP 客户端实例。
//
// 参数:
//   - client: net/http.Client 实例（建议设置 Timeout）
//
// 返回: *DefaultHTTPClient
func NewDefaultHTTPClient(client *http.Client) *DefaultHTTPClient {
	return &DefaultHTTPClient{
		client: client,
	}
}

// Post 发送 JSON POST 请求，等同于 PostWithHeaders(url, params, nil, target)。
//
// 参数:
//   - url: 请求地址
//   - params: 请求参数（将被 JSON 序列化）
//   - target: 响应反序列化目标
//
// 返回: error 请求或解析过程中的错误
func (c *DefaultHTTPClient) Post(url string, params interface{}, target interface{}) error {
	return c.PostWithHeaders(url, params, nil, target)
}

// PostWithHeaders 发送带自定义请求头的 JSON POST 请求。
// 网络错误时自动重试一次。响应 Body 会被完整读取并反序列化到 target。
//
// 参数:
//   - url: 请求地址
//   - params: 请求参数（将被 JSON 序列化）
//   - headers: 自定义请求头（可为 nil）
//   - target: 响应反序列化目标
//
// 返回: error 请求或解析过程中的错误
func (c *DefaultHTTPClient) PostWithHeaders(url string, params interface{}, headers map[string]string, target interface{}) error {
	// 序列化请求数据
	jsonData, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("[SDK %s] marshal params failed: %w", url, err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("[SDK %s] create request failed: %w", url, err)
	}

	// 设置默认头部
	req.Header.Set("Content-Type", "application/json")

	// 设置自定义头部
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求（带重试：网络错误重试1次）
	var resp *http.Response
	for attempt := 0; attempt < 2; attempt++ {
		resp, err = c.client.Do(req)
		if err == nil {
			break
		}
		if attempt == 0 {
			// 第一次失败，重建 request body 重试
			req.Body = io.NopCloser(bytes.NewBuffer(jsonData))
			continue
		}
	}
	if err != nil {
		return fmt.Errorf("[SDK %s] request failed after retry: %w", url, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[SDK %s] bad status: %s", url, resp.Status)
	}

	// 读取响应
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[SDK %s] read response failed: %w", url, err)
	}

	if err := json.Unmarshal(response, target); err != nil {
		// 截取前500字节方便排查
		raw := string(response)
		if len(raw) > 500 {
			raw = raw[:500] + "...(truncated)"
		}
		return fmt.Errorf("[SDK %s] json unmarshal failed: %w, raw: %s", url, err, raw)
	}
	return nil
}

// PostMultipart 发送 multipart/form-data POST 请求（用于文件上传）。
// 网络错误时自动重试一次。响应 Body 会被完整读取并反序列化到 target。
//
// 参数:
//   - url: 请求地址
//   - fields: 文本字段（key-value）
//   - fileField: 文件字段名
//   - fileName: 文件名
//   - fileData: 文件内容（io.Reader）
//   - target: 响应反序列化目标
//
// 返回: error 请求或解析过程中的错误
func (c *DefaultHTTPClient) PostMultipart(url string, fields map[string]string, fileField, fileName string, fileData io.Reader, target interface{}) error {
	// 将 fileData 读入内存，以便重试时可以重用；限制最大 500MB 防止 OOM
	const maxFileSize = 500 << 20 // 500MB
	limitedReader := io.LimitReader(fileData, maxFileSize+1)
	fileBytes, err := io.ReadAll(limitedReader)
	if err != nil {
		return fmt.Errorf("[SDK %s] read file data failed: %w", url, err)
	}
	if int64(len(fileBytes)) > maxFileSize {
		return fmt.Errorf("[SDK %s] file size exceeds limit (%dMB)", url, maxFileSize>>20)
	}

	// buildBody 构建 multipart body
	buildBody := func() (*bytes.Buffer, string, error) {
		body := bytes.NewBuffer(make([]byte, 0, len(fileBytes)+4096))
		writer := multipart.NewWriter(body)

		// 写入文本字段
		for key, value := range fields {
			if err := writer.WriteField(key, value); err != nil {
				return nil, "", fmt.Errorf("[SDK %s] write field %s failed: %w", url, key, err)
			}
		}

		// 写入文件字段
		part, err := writer.CreateFormFile(fileField, fileName)
		if err != nil {
			return nil, "", fmt.Errorf("[SDK %s] create form file failed: %w", url, err)
		}
		if _, err := io.Copy(part, bytes.NewReader(fileBytes)); err != nil {
			return nil, "", fmt.Errorf("[SDK %s] copy file data failed: %w", url, err)
		}

		if err := writer.Close(); err != nil {
			return nil, "", fmt.Errorf("[SDK %s] close multipart writer failed: %w", url, err)
		}
		return body, writer.FormDataContentType(), nil
	}

	// 发送请求（带重试：网络错误重试1次）
	var resp *http.Response
	for attempt := 0; attempt < 2; attempt++ {
		body, contentType, err := buildBody()
		if err != nil {
			return err
		}

		req, err := http.NewRequest("POST", url, body)
		if err != nil {
			return fmt.Errorf("[SDK %s] create request failed: %w", url, err)
		}
		req.Header.Set("Content-Type", contentType)

		resp, err = c.client.Do(req)
		if err == nil {
			break
		}
		if attempt == 1 {
			return fmt.Errorf("[SDK %s] request failed after retry: %w", url, err)
		}
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[SDK %s] bad status: %s", url, resp.Status)
	}

	// 读取响应
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[SDK %s] read response failed: %w", url, err)
	}

	if err := json.Unmarshal(response, target); err != nil {
		raw := string(response)
		if len(raw) > 500 {
			raw = raw[:500] + "...(truncated)"
		}
		return fmt.Errorf("[SDK %s] json unmarshal failed: %w, raw: %s", url, err, raw)
	}
	return nil
}
