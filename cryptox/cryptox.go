package cryptox

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash/crc32"
)

// Sha256Encode
func Sha256Encode(param string) string {
	hash := sha256.New()
	hash.Write([]byte(param))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum)
}

// Crc32IEEE
func Crc32IEEE(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

// Md5Encode
func Md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
