package common

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// 32 bits token
func GenToken() string {
	rand.Seed(time.Now().UnixNano())
	uLen := 32
	b := make([]byte, uLen)
	rand.Read(b)
	return hex.EncodeToString(b)[0:uLen]
}
