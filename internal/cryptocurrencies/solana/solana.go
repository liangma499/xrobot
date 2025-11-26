package solana

import (
	"tron_robot/internal/cryptocurrencies/solana/internal"
)

func SolanaPubkeyToAddress() (string, string) {

	// 生成新的密钥对
	kp := internal.NewWallet()

	// 获取公钥和私钥
	publicKey := kp.PublicKey()
	privateKey := kp.PrivateKey

	// 输出公钥和私钥
	return privateKey.String(), publicKey.String()
}
