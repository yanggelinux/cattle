package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

const prefix = "Fh@d%12"

func Md5Salt(str string) string {
	str += prefix
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func HashBase64(input string) string {
	sum := sha256.Sum256([]byte(input))
	return hex.EncodeToString(sum[:])
}
