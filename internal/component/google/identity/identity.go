package identity

import (
	"context"
	"sync"
	"xbase/encoding/json"
	"xbase/etc"
	"xbase/log"

	"google.golang.org/api/idtoken"
)

var (
	once     sync.Once
	instance *Identity
)

type Profile struct {
	ID            string `json:"sub"`            // ID
	Name          string `json:"name"`           // 名称
	Email         string `json:"email"`          // 邮箱
	EmailVerified bool   `json:"email_verified"` // 邮箱是否已验证
	Picture       string `json:"picture"`        // 头像
	FamilyName    string `json:"family_name"`    // 姓
	GivenName     string `json:"given_name"`     // 名
}

type Config struct {
	Audience string `json:"audience"` // 对应着客户端ID
}

func Instance() *Identity {
	once.Do(func() {
		instance = NewInstance("etc.google.identity.default")
	})

	return instance
}

// NewInstance 新建实例
func NewInstance[T string | Config | *Config](config T) *Identity {
	var (
		conf *Config
		v    any = config
	)

	switch c := v.(type) {
	case string:
		conf = &Config{}
		if err := etc.Get(c).Scan(conf); err != nil {
			log.Fatalf("load google identity config failed: %v", err)
		}
	case Config:
		conf = &c
	case *Config:
		conf = c
	}

	return &Identity{conf: conf}
}

type Identity struct {
	conf *Config
}

// Validate 校验id token
func (i *Identity) Validate(ctx context.Context, idToken string) (*Profile, error) {
	payload, err := idtoken.Validate(ctx, idToken, i.conf.Audience)
	if err != nil {
		return nil, err
	}

	buf, err := json.Marshal(payload.Claims)
	if err != nil {
		return nil, err
	}

	profile := &Profile{}

	err = json.Unmarshal(buf, profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
