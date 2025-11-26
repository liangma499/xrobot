package xcrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

//var desKey = []byte{0x16, 0x18, 0xFA, 0xE6, 0xF9, 0x67, 0x98, 0xA2}

func DesEncrypt(origData, desKey []byte) ([]byte, error) {
	block, err := des.NewCipher(desKey)
	if err != nil {
		return nil, err
	}

	origData = PKCS5Padding(origData, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, desKey)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	//加为base64
	encoded := []byte(base64.StdEncoding.EncodeToString(crypted))
	return encoded, nil
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func DesDecrypt(crypted, desKey []byte) ([]byte, error) {

	decrData, err := base64.StdEncoding.DecodeString(string(crypted[:]))
	if err != nil {
		return nil, err
	}
	block, err := des.NewCipher(desKey)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, desKey)
	origData := make([]byte, len(decrData))

	blockMode.CryptBlocks(origData, decrData)
	origData = PKCS5UnPadding(origData)

	return origData, nil
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	if length == 0 {
		return origData
	}
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
