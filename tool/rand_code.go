package tool

import (
	"fmt"
	"math/rand"
	"time"
)

// 6位随机验证码
func RandCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println(rndCode)
	return rndCode
}
