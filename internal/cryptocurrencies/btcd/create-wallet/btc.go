package createwallet

import (
	"encoding/hex"

	"xbase/log"

	"github.com/btcsuite/btcd/btcec/v2"
)

func BTC_PubkeyToAddress() (string, string) {

	// 生成新的私钥
	privKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Warnf("btcd failed to generate private key: %v", err)
		return "", ""
	}

	// 获取公钥
	pubKey := privKey.PubKey()

	// 输出私钥和公钥
	//fmt.Printf("Private Key: %s\n", hex.EncodeToString(privKey.Serialize()))
	//fmt.Printf("Public Key: %s\n", hex.EncodeToString(pubKey.SerializeCompressed()))
	return hex.EncodeToString(privKey.Serialize()), hex.EncodeToString(pubKey.SerializeCompressed())
}
