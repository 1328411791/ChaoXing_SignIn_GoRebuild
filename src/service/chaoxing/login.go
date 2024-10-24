package chaoxing

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"log"
	"math"
)

type UserCookieType map[string]string

func userLogin(phone string, password string) (UserCookieType, error) {
	// 将key转换为byte数组
	key := "u2oh6Vu^HWe40fj"
	// 使用DES算法加密密码
	encryptedPassword, err := encryptDES(password, key)

	if err != nil {
		log.Fatal()
	}

	log.Println(encryptedPassword)

	return nil, nil
}

// ceil 向上取整函数
func ceil(n float64) int {
	return int(math.Ceil(n))
}

// padPKCS7 根据DES的块大小进行PKCS7填充
func padPKCS7(src []byte) []byte {
	padding := des.BlockSize - len(src)%des.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

// encryptDES 使用DES算法和ECB模式加密密码
func encryptDES(password, key string) (string, error) {
	keyBytes := []byte(key)
	if len(keyBytes) != 8 {
		return "", fmt.Errorf("key must be 8 bytes long")
	}

	block, err := des.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	paddedPassword := padPKCS7([]byte(password))
	ciphertext := make([]byte, len(paddedPassword))
	for i := 0; i < len(paddedPassword); i += des.BlockSize {
		block.Encrypt(ciphertext[i:i+des.BlockSize], paddedPassword[i:i+des.BlockSize])
	}

	return hex.EncodeToString(ciphertext), nil
}
