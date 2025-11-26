package xtypes

import (
	"xbase/errors"
	"xbase/log"
	"xbase/utils/xrand"
	"xrobot/internal/code"
	"xrobot/internal/utils/xcrypt"
)

// 加密密码
func EncryptPassword(password string) (string, string, error) {
	salt := xrand.Letters(8)

	hashed, err := xcrypt.Encrypt(password, salt)
	if err != nil {
		log.Errorf("encrypt password failed, password = %s salt = %s err = %v", password, salt, err)
		return "", "", errors.NewError(err, code.InternalError)
	}

	return salt, hashed, nil
}
