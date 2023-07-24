package middlewares

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Check(content, encrypted string) bool {
	return strings.EqualFold(Encode(content), encrypted)
}
func Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
