package authenticator

import (
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"image/png"
	"sync"
	"xbase/etc"
	"xbase/log"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

var (
	once     sync.Once
	instance *Authenticator
)

type Config struct {
	Issuer       string `json:"issuer"`       // 发行方
	QrcodeWidth  int    `json:"qrcodeWidth"`  // 二维码宽度
	QrcodeHeight int    `json:"qrcodeHeight"` // 二维码高度
}

var b64 = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
var b32 = base32.StdEncoding.WithPadding(base32.NoPadding)

func Instance() *Authenticator {
	once.Do(func() {
		instance = NewInstance("etc.authenticator.default")
	})

	return instance
}

// NewInstance 新建实例
func NewInstance[T string | Config | *Config](config T) *Authenticator {
	var (
		conf *Config
		v    any = config
	)

	switch c := v.(type) {
	case string:
		conf = &Config{}
		if err := etc.Get(c).Scan(conf); err != nil {
			log.Fatalf("load google authenticator config failed: %v", err)
		}
	case Config:
		conf = &c
	case *Config:
		conf = c
	}

	return &Authenticator{conf: conf}
}

type Authenticator struct {
	conf *Config
}

// Validate 校验验证码
func (a *Authenticator) Validate(passcode, secret string) bool {
	return totp.Validate(passcode, secret)
}

// Generate 生成验证码
func (a *Authenticator) Generate(account string, secret ...string) *Generator {
	opts := totp.GenerateOpts{
		Issuer:      a.conf.Issuer,
		AccountName: account,
	}

	g := &Generator{a: a}

	if len(secret) > 0 {
		opts.Secret, g.err = b32.DecodeString(secret[0])
		if g.err != nil {
			return g
		}
	}

	g.key, g.err = totp.Generate(opts)

	return g
}

type Generator struct {
	a   *Authenticator
	err error
	key *otp.Key
}

// Secret 获取密钥
func (g *Generator) Secret() (string, error) {
	if g.err != nil {
		return "", g.err
	}

	return g.key.Secret(), nil
}

// Url 获取链接
func (g *Generator) Url() (string, error) {
	if g.err != nil {
		return "", g.err
	}

	return g.key.URL(), nil
}

// Qrcode 获取二维码
func (g *Generator) Qrcode() (string, error) {
	if g.err != nil {
		return "", g.err
	}

	img, err := g.key.Image(g.a.conf.QrcodeWidth, g.a.conf.QrcodeHeight)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}

	err = png.Encode(buf, img)
	if err != nil {
		return "", err
	}

	return "data:image/png;base64," + b64.EncodeToString(buf.Bytes()), nil
}
