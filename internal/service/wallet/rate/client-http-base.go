package rate

import (
	"xbase/network/http"
)

type clientBase struct {
	client *http.Client
}
type RateReq struct {
	Tokens     string `json:"tokens"`     // 查询
	Currencies string `json:"currencies"` //对应
}
type CryptoCompareRateReq struct {
	Fsym  string `json:"fsym"`  // 查询
	Tsyms string `json:"tsyms"` //对应tsyms=usd,ton,usdt,vnd,trx,btc,eth,bnb,sol,usdc
}

func newCryptoCompareRatesClient() *clientBase {
	c := &clientBase{client: http.NewClient()}
	c.client.SetBaseUrl("https://min-api.cryptocompare.com")
	c.client.SetHeaders(map[string]string{
		"Accept": "/",
	})
	return c
}

// Get 执行Get请求
func (c *clientBase) get(url string, req, resp any) error {
	return c.request(http.MethodGet, url, req, resp)
}

/*
// Post 执行POST请求
func (c *clientBase) post(url string, req, resp any) error {
	return c.request(http.MethodPost, url, req, resp)
}

// Delete 执行DELETE请求
func (c *clientBase) delete(url string, req, resp any) error {
	return c.request(http.MethodDelete, url, req, resp)
}
*/
// 执行请求
func (c *clientBase) request(method string, url string, req, resp any) error {
	res, err := c.client.Request(method, url, req)
	if err != nil {
		return err
	}

	return res.ScanBody(resp)
}
