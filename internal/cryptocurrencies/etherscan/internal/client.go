package internal

import (
	"time"
	"xbase/network/http"
)

type Client struct {
	client *http.Client
}

func NewClient(baseUrl string, bSetSjon bool) *Client {
	c := &Client{client: http.NewClient()}
	if bSetSjon {
		c.client.SetContentType(http.ContentTypeJson)
	}

	c.client.SetTimeout(30 * time.Second)
	c.client.SetBaseUrl(baseUrl)
	//c.client.SetHeader("TRON-PRO-API-KEY", apiKey)

	return c
}

// Get 执行Get请求
func (c *Client) Get(url string, req any, resp any, opts ...*http.RequestOptions) error {
	return c.request(http.MethodGet, url, req, resp, opts...)
}

// Post 执行POST请求
func (c *Client) Post(url string, req any, resp any, opts ...*http.RequestOptions) error {
	return c.request(http.MethodPost, url, req, resp, opts...)
}

// Delete 执行DELETE请求
func (c *Client) Delete(url string, req any, resp any, opts ...*http.RequestOptions) error {
	return c.request(http.MethodDelete, url, req, resp, opts...)
}

// 执行请求
func (c *Client) request(method string, url string, req any, resp any, opts ...*http.RequestOptions) error {

	res, err := c.client.Request(method, url, req, opts...)
	if err != nil {
		return err
	}
	return res.ScanBody(resp)
}
