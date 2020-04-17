package formatting

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(pwd string) string {
	hash := sha256.New()
	defer hash.Reset()
	hash.Write([]byte(pwd))
	return hex.EncodeToString(hash.Sum(nil))
}
