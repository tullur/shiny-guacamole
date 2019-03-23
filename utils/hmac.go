package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// HMAC -> wrapper arount the crypto/hmac package
type HMAC struct {
	hmac hash.Hash
}

// NewHMAC -> return a new HMAC object
func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))

	return HMAC{
		hmac: h,
	}
}

// Hash -> hash input string
func (h HMAC) Hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}
