package xtypes

import (
	"database/sql/driver"
	"encoding/json"
	"xbase/errors"
	"xbase/log"
	"xrobot/internal/utils/xcrypt"
)

type PrivateKeyCfg struct {
	FromAddress string            `json:"from_address,omitempty"` // trc20转转账地址
	PrivateKey  string            `json:"private_key,omitempty"`  // 私钥
	MaxFeeLimit int64             `json:"feeLimit,omitempty"`     //最高手续续
	AppID       string            `json:"appID,omitempty"`
	Testing     bool              `json:"testing,omitempty"`
	Status      OptionStatus      `json:"status,omitempty"`
	ExtraCfg    *PlatformExtraCfg `json:"extra_cfg,omitempty"` //连接地址
}

type PrivateKeyCfgList struct {
	PrivateKeyCfg map[Currency]*PrivateKeyCfg `json:"privateKeyCfg,omitempty"`
}

var configDesKey = []byte{0x16, 0x18, 0xFA, 0xE6, 0xF9, 0x67, 0x98, 0xA2}

type EncryptByte []byte

func (e EncryptByte) Value() (driver.Value, error) {
	if e == nil {
		return nil, nil
	}
	if len(e) == 0 {
		return nil, nil
	}
	return xcrypt.DesEncrypt(e, configDesKey)
}

func (e EncryptByte) Clone() EncryptByte {
	b2 := make([]byte, len(e))
	copy(b2, e)
	return b2
}

func (e *EncryptByte) Scan(value any) error {
	if value == nil {
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	*e = make(EncryptByte, 0)
	*e = append(*e, s...)
	return nil
}
func (e EncryptByte) DesDecryptToByte() []byte {
	if len(e) == 0 {
		return nil
	}
	//解密
	byteData, err := xcrypt.DesDecrypt(e, configDesKey)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil
	}

	return byteData
}

func (e EncryptByte) Config() *PrivateKeyCfgList {
	if len(e) == 0 {
		return nil
	}
	value := new(PrivateKeyCfgList)
	err := json.Unmarshal([]byte(e), value)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil
	}

	return value
}
func (e EncryptByte) DesConfig() *PrivateKeyCfgList {
	if len(e) == 0 {
		return nil
	}
	byteData, err := xcrypt.DesDecrypt(e, configDesKey)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil
	}
	value := new(PrivateKeyCfgList)

	err = json.Unmarshal(byteData, value)
	if err != nil {
		log.Warnf("unmarshal:%v", err)

		return nil
	}

	return value
}

func (e *PrivateKeyCfgList) ToJsonByte() EncryptByte {
	data, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return data
}
