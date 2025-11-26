package xcrypt

import (
	"database/sql/driver"
	"xbase/errors"
	"xbase/log"
)

const (
	CollectionCfgDesKeyHex = "************"
)

var collectionCfgDesKey = []byte{0x26, 0x98, 0xA1, 0x6E, 0xF9, 0x67, 0x98, 0xA2}

type CollectionCfgByte []byte

func (e CollectionCfgByte) Value() (driver.Value, error) {
	if e == nil {
		return nil, nil
	}
	if len(e) == 0 {
		return nil, nil
	}
	return DesEncrypt(e, collectionCfgDesKey)
}

func (e CollectionCfgByte) Clone() CollectionCfgByte {
	b2 := make([]byte, len(e))
	copy(b2, e)
	return b2
}

func (e *CollectionCfgByte) Scan(value any) error {
	if value == nil {
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	*e = make(CollectionCfgByte, 0)
	*e = append(*e, s...)
	return nil
}

func (e CollectionCfgByte) DesToString() string {
	if len(e) == 0 {
		return ""
	}
	byteData, err := DesDecrypt(e, collectionCfgDesKey)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return ""
	}

	return string(byteData)
}

// 前端需要显示的
func (e CollectionCfgByte) Encrypt() string {
	return CollectionCfgDesKeyHex
}
