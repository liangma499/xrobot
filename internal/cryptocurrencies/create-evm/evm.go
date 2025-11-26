package createwallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"xbase/log"

	"github.com/ethereum/go-ethereum/crypto"
)

func EvmPubkeyToAddress() (string, string) {
	// 生成新的私钥
	priv, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		log.Warnf("evm failed to generate private key: %v", err)
		return "", ""
	}

	// 获取公钥
	pub := priv.Public()

	// 获取地址
	address := crypto.PubkeyToAddress(*pub.(*ecdsa.PublicKey))

	// 输出私钥和地址
	privBytes := crypto.FromECDSA(priv)
	privBytesStr := fmt.Sprintf("%x", privBytes) // 输出私钥（十六进制格式）
	return privBytesStr, address.Hex()
}
