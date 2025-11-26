package xdrand

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"strconv"
	"sync"
	"time"

	"github.com/drand/drand/client"
	"github.com/drand/drand/client/http"
)

const (
	chainHashStr = "52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971"
)
const (
	DrandVerifyUrl = "https://api.drand.sh/52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971/public/%d"
)

var chainHash, _ = hex.DecodeString(chainHashStr)

var urls = []string{
	"https://api.drand.sh",
	"https://drand.cloudflare.com",
}
var (
	once     sync.Once
	instance *Drand
)

type Drand struct {
}
type DrandRes struct {
	Round      int64  `json:"round"`
	Randomness string `json:"randomness"`
	Signature  string `json:"signature"`
	ChainHash  string `json:"chainHash"`
}

func Instance() *Drand {
	once.Do(func() {
		instance = newInstance()
	})
	return instance
}
func newInstance() *Drand {
	return &Drand{}
}

func (d *Drand) GetResult(seed, signature string) (int64, string) {

	hash := d.generateHMAC([]byte(seed), []byte(signature))
	result := uint64(0)
	for i := 0; i < 8; i++ {
		dist16Int64, _ := strconv.ParseUint(hash[i*8:i*8+8], 16, 64)
		result += dist16Int64 / 2
	}
	return int64(result), hash
}
func (d *Drand) GetDrandRandomness() *DrandRes {
	c, err := client.New(
		client.From(http.ForURLs(urls, chainHash)...),
		client.WithChainHash(chainHash),
	)
	if err != nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	data, err := c.Get(ctx, 0)
	if err != nil {
		return nil
	}
	return &DrandRes{
		Round:      int64(data.Round()),
		Randomness: hex.EncodeToString(data.Randomness()),
		Signature:  hex.EncodeToString(data.Signature()),
		ChainHash:  chainHashStr,
	}
}

func (d *Drand) generateHMAC(key, data []byte) string {
	// 创建一个新的HMAC实例，使用SHA-512哈希算法
	h := hmac.New(sha512.New, key)

	// 写入要计算HMAC的数据
	h.Write(data)

	// 计算HMAC值
	hmacValue := h.Sum(nil)

	// 将HMAC值转换为十六进制字符串
	return hex.EncodeToString(hmacValue)
}
