package utility

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

// padPKCS7 根据DES的块大小进行PKCS7填充
func EncryptDES(password string, keyBytes []byte) (string, error) {

	if len(keyBytes) != 8 {
		return "", fmt.Errorf("key must be 8 bytes long")
	}

	block, err := des.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	paddedPassword := padPKCS7([]byte(password))
	ciphertext := make([]byte, len(paddedPassword))
	// 使用ECB模式加密
	for i := 0; i < len(paddedPassword); i += des.BlockSize {
		block.Encrypt(ciphertext[i:i+des.BlockSize], paddedPassword[i:i+des.BlockSize])
	}

	return hex.EncodeToString(ciphertext), nil
}

// padPKCS7 进行PKCS7填充
func padPKCS7(data []byte) []byte {
	blockSize := 8 // DES 加密块大小为 8 字节
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}
