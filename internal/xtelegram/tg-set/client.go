package tgSet

import (
	"xbase/network/http"
)

const tgBaseUrl = "https://api.telegram.org"

type client struct {
	client *http.Client
}

func NewClient(token string) *client {
	baseUrl := tgBaseUrl + "/bot" + token
	c := &client{client: http.NewClient()}
	c.client.SetBaseUrl(baseUrl)
	c.client.SetHeaders(map[string]string{
		"Accept": "/",
	})
	return c
}

// Get 执行Get请求
func (c *client) Get(url string, req, resp any) error {
	return c.request(http.MethodGet, url, req, resp)
}

// 执行请求
func (c *client) request(method string, url string, req, resp any) error {
	res, err := c.client.Request(method, url, req)
	if err != nil {
		return err
	}

	return res.ScanBody(resp)
}
