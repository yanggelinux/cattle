package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Client 封装

type Client struct {
	httpClient *http.Client
}

func New(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
	}
}

// 通用请求方法
func (c *Client) request(ctx context.Context, method, urlStr string, headers map[string]string, body io.Reader) ([]byte, int, error) {
	req, err := http.NewRequestWithContext(ctx, method, urlStr, body)
	if err != nil {
		return nil, 0, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return respBody, resp.StatusCode, nil
}

// GET请求
func (c *Client) Get(ctx context.Context, urlStr string, params map[string]string, headers map[string]string) ([]byte, int, error) {
	if params != nil {
		u, err := url.Parse(urlStr)
		if err != nil {
			return nil, 0, err
		}
		q := u.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
		urlStr = u.String()
	}
	return c.request(ctx, http.MethodGet, urlStr, headers, nil)
}

// POST请求
func (c *Client) Post(ctx context.Context, urlStr string, data interface{}, headers map[string]string) ([]byte, int, error) {
	var body io.Reader
	if data != nil {
		switch v := data.(type) {
		case []byte:
			// 如果是 []byte，直接用作请求体
			body = bytes.NewReader(v)
		default:
			// 其余类型：json.Marshal，再写入 body
			jsonData, err := json.Marshal(data)
			if err != nil {
				return nil, 0, err
			}
			body = bytes.NewReader(jsonData)

			// 如果需要序列化为 JSON，也需要设置 Content-Type
			if headers == nil {
				headers = make(map[string]string)
			}
			if _, exists := headers["Content-Type"]; !exists {
				headers["Content-Type"] = "application/json"
			}
		}
	}
	return c.request(ctx, http.MethodPost, urlStr, headers, body)
}

// PUT请求
func (c *Client) Put(ctx context.Context, urlStr string, data interface{}, headers map[string]string) ([]byte, int, error) {
	var body io.Reader
	if data != nil {
		switch v := data.(type) {
		case []byte:
			// 如果是 []byte，直接用作请求体
			body = bytes.NewReader(v)
		default:
			// 其余类型：json.Marshal，再写入 body
			jsonData, err := json.Marshal(data)
			if err != nil {
				return nil, 0, err
			}
			body = bytes.NewReader(jsonData)

			// 如果需要序列化为 JSON，也需要设置 Content-Type
			if headers == nil {
				headers = make(map[string]string)
			}
			if _, exists := headers["Content-Type"]; !exists {
				headers["Content-Type"] = "application/json"
			}
		}
	}
	return c.request(ctx, http.MethodPut, urlStr, headers, body)
}

// DELETE请求
func (c *Client) Delete(ctx context.Context, urlStr string, data interface{}, headers map[string]string) ([]byte, int, error) {
	var body io.Reader
	if data != nil {
		switch v := data.(type) {
		case []byte:
			// 如果是 []byte，直接用作请求体
			body = bytes.NewReader(v)
		default:
			// 其余类型：json.Marshal，再写入 body
			jsonData, err := json.Marshal(data)
			if err != nil {
				return nil, 0, err
			}
			body = bytes.NewReader(jsonData)
			// 如果需要序列化为 JSON，也需要设置 Content-Type
			if headers == nil {
				headers = make(map[string]string)
			}
			if _, exists := headers["Content-Type"]; !exists {
				headers["Content-Type"] = "application/json"
			}
		}
	}
	return c.request(ctx, http.MethodDelete, urlStr, headers, body)
}
