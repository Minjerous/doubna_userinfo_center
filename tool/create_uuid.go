package tool

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func CreateUUID() string {
	rand.Seed(time.Now().Unix())
	var randomBytes = make([]byte, 16)
	for i := 0; i < 16; i++ {
		randomBytes[i] = byte(rand.Intn(128))
	}
	randomBytes[6] &= 0x0f /* clear version        */
	randomBytes[6] |= 0x40 /* set to version 4     */
	randomBytes[8] &= 0x3f /* clear variant        */
	randomBytes[8] |= 0x80 /* set to IETF variant  */
	return hex.EncodeToString(randomBytes)
}
