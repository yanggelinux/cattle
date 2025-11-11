package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/pkg/log"
	"go.uber.org/zap"
	"io"
)

const sharedKeyBase64 = "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODA5MTI="

var sharedKey []byte

func DecryptPassword(ivB64, cipherB64 string) (string, error) {
	// 解码 Base64
	iv, err := base64.StdEncoding.DecodeString(ivB64)
	if err != nil {
		return "", errors.Wrap(err, "invalid iv encoding")
	}
	ct, err := base64.StdEncoding.DecodeString(cipherB64)
	if err != nil {
		return "", errors.Wrap(err, "invalid cipher encoding")
	}
	// 创建 AES-GCM
	block, err := aes.NewCipher(sharedKey)
	if err != nil {
		return "", errors.Wrap(err, "aes.NewCipher")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrap(err, "cipher.NewCipher")

	}
	// 解密
	pt, err := gcm.Open(nil, iv, ct, nil)
	if err != nil {
		return "", errors.Wrap(err, "decrypt failed")
	}
	return string(pt), nil
}

func EncryptPassword(password string) (ivB64, cipherB64 string, err error) {
	// 1. 用 sharedKey 创建 AES block
	block, err := aes.NewCipher(sharedKey)
	if err != nil {
		return "", "", errors.New("aes.NewCipher: " + err.Error())
	}
	// 2. 用 block 创建 GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", errors.New("cipher.NewGCM: " + err.Error())
	}
	// 3. 为 GCM 生成随机 12 字节 IV
	iv := make([]byte, gcm.NonceSize()) // GCM 标准下通常是 12 字节
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", "", errors.New("generate iv: " + err.Error())
	}

	// 4. 对明文 password 进行加密
	ciphertext := gcm.Seal(nil, iv, []byte(password), nil)

	// 5. 对 iv 和 ciphertext 做 Base64 编码后返回
	ivB64 = base64.StdEncoding.EncodeToString(iv)
	cipherB64 = base64.StdEncoding.EncodeToString(ciphertext)
	return ivB64, cipherB64, nil
}

// 生成 sharedKeyBase64 的方法
func genBase64Key() string {
	key := make([]byte, 32) // 32 字节
	if _, err := rand.Read(key); err != nil {
		log.Logger.Error("generate key fail", zap.Error(err))
	}
	return base64.StdEncoding.EncodeToString(key)
}

func init() {
	var err error
	sharedKey, err = base64.StdEncoding.DecodeString(sharedKeyBase64)
	if err != nil {
		log.Logger.Error("invalid shared key", zap.Error(err))
	}
	if len(sharedKey) != 32 {
		log.Logger.Error("shared key must be 32 bytes for AES-256-GCM", zap.Error(err))
	}
}
