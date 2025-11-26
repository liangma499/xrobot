package helper

import (
	"net/http"
	"xbase/utils/jwt"
	jwtcomp "xrobot/internal/component/jwt"
)

// GenerateToken 生成Token
func GenerateToken(uid int64, data ...map[string]any) (*jwt.Token, error) {
	ins := jwtcomp.Instance()

	payload := jwtcomp.Payload{ins.IdentityKey(): uid}
	if len(data) > 0 {
		for k, v := range data[0] {
			payload[k] = v
		}
	}

	return ins.GenerateToken(payload)
}

// DestroyToken 销毁Token
func DestroyToken(r *http.Request) error {
	return jwtcomp.Instance().Http().DestroyToken(r)
}
