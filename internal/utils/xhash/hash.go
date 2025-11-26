package xhash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
)

func MD5(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

func Hash256Hmac(data string, key ...string) string {
	var h hash.Hash

	if len(key) > 0 {
		h = hmac.New(sha256.New, []byte(key[0]))
	} else {
		h = hmac.New(sha256.New, nil)
	}

	h.Write([]byte(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}
