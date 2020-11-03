package cryptox

import (
	"crypto/sha256"
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

func Crc32IEEE(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
