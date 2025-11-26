package httpclient

import (
	"time"
	"xbase/network/http"
)

type HttpClient struct {
	client *http.Client
}

func NewClient(baseUrl string, bSetSjon bool) *HttpClient {
	c := &HttpClient{client: http.NewClient()}
	if bSetSjon {
		c.client.SetContentType(http.ContentTypeJson)
	}

	//c.client.SetHeader("x-api-key",)
	c.client.SetTimeout(40 * time.Second)
	c.client.SetBaseUrl(baseUrl)
	//c.client.SetHeader("TRON-PRO-API-KEY", apiKey)

	return c
}
func (c *HttpClient) SetHeader(key, value string) {
	c.client.SetHeader(key, value)
}

// Get 执行Get请求
func (c *HttpClient) Get(url string, req any, resp any, opts ...*http.RequestOptions) error {
	return c.request(http.MethodGet, url, req, resp, opts...)
}

// Post 执行POST请求
func (c *HttpClient) Post(url string, req any, resp any, opts ...*http.RequestOptions) error {
	return c.request(http.MethodPost, url, req, resp, opts...)
}

// Delete 执行DELETE请求
func (c *HttpClient) Delete(url string, req any, resp any, opts ...*http.RequestOptions) error {
	return c.request(http.MethodDelete, url, req, resp, opts...)
}

// 执行请求
func (c *HttpClient) request(method string, url string, req any, resp any, opts ...*http.RequestOptions) error {

	res, err := c.client.Request(method, url, req, opts...)
	if err != nil {
		return err
	}
	return res.ScanBody(resp)
}
