package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
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

//md5 加解密参数
func CheckSign(params map[string]string, key string, filter string) (sign string) {

	nk := make([]string, 0)
	for k, _ := range params {
		if k == filter {
			continue
		}
		nk = append(nk, k)
	}

	sort.Strings(nk)

	var originStr string
	for _, v := range nk {
		originStr += fmt.Sprintf("%v=%v&", v, params[v])
	}

	originStr += fmt.Sprintf("key=%v", key)
	h := md5.New()
	h.Write([]byte(originStr))
	sign = hex.EncodeToString(h.Sum(nil)[:])
	return
}
