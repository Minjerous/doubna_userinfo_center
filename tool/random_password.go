package tool

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"time"
)

func GenerateRandomPassword() string {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(10) + 8
	var b = make([]byte, 0)
	for i := 0; i < n; i++ {
		b = append(b, byte(rand.Intn(128)))
	}
	hash := hmac.New(sha256.New, []byte("MJ_happy"))
	hash.Write(b)
	return base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(hash.Sum(nil))))[0:8]
}
