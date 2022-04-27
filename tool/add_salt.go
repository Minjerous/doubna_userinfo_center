package tool

import (
	"crypto/md5"
	"encoding/hex"
)

//加盐加密
func HashWithSalted(plain, Salt string) string {
	m5 := md5.New()
	m5.Write([]byte(plain))
	m5.Write([]byte(Salt))
	src := m5.Sum(nil)
	return hex.EncodeToString(src)
}

//解密判断
func Match(secret, plain, Salt string) bool {
	m5 := md5.New()
	m5.Write([]byte(plain))
	m5.Write([]byte(Salt))
	src := m5.Sum(nil)
	hashPwd := hex.EncodeToString(src)
	return hashPwd == secret
}
