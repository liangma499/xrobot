package tron

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"xbase/log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"golang.org/x/crypto/sha3"
)

func Tron_PubkeyToAddress() (string, string, string) {

	// 生成私钥
	privKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		log.Warnf("trc20 failed to generate private key: %v", err)
		return "", "", ""
	}

	// 获取公钥
	pubKey := privKey.Public()

	// 输出私钥
	privKeyBytes := crypto.FromECDSA(privKey)

	// 输出公钥
	pubKeyBytes := crypto.FromECDSAPub(pubKey.(*ecdsa.PublicKey))

	// 生成 TRON 地址
	tronAddress := tron_generateTronAddress(pubKeyBytes)
	return hex.EncodeToString(privKeyBytes), hex.EncodeToString(pubKeyBytes), tronAddress
}

// 生成 TRON 地址
func tron_generateTronAddress(pubKeyBytes []byte) string {
	// Keccak256 哈希
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKeyBytes[1:])                      // 去掉第一个字节
	address := hash.Sum(nil)[len(hash.Sum(nil))-20:] // 取最后20个字节

	// 版本字节加地址
	versionedAddress := append([]byte{0x41}, address...)

	// 计算校验和
	checksum := sha256.Sum256(versionedAddress)
	checksum = sha256.Sum256(checksum[:]) // 双重SHA256
	checksum2 := checksum[:4]             // 取前4字节

	// 组合地址和校验和
	finalAddress := append(versionedAddress, checksum2...)

	// Base58 编码
	return base58.Encode(finalAddress)
}
