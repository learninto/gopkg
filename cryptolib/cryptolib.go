package cryptolib

import (
	"crypto/sha256"
	"fmt"
)

// Sha256Encode
func Sha256Encode(param string) string {
	hash := sha256.New()
	hash.Write([]byte(param))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum)
}
