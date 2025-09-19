package dto

import (
	"crypto/hmac"
	"encoding/base64"
	"hash"
	"strings"
)

// HashMac
func HashMac(f func() hash.Hash, secret string, rawArr ...string) string {
	raw := strings.Join(rawArr, "\n") + "\n"
	h := hmac.New(f, []byte(secret))
	h.Write([]byte(raw))
	sign := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(sign)
}
