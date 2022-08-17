package common

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	tk := GenToken()
	if tk == "" {
		t.Error("token empty")
	}

	len := len(tk)
	if len != 32 {
		fmt.Println(len)
		t.Error("token len not 32")
	}

}

func TestChecksign(t *testing.T) {

	signEx := "fd0b2ade49986d464370d1d29a23decf"
	var testParam = make(map[string]string)
	testParam["name"] = "afei"
	testParam["pwd"] = "e10adc3949ba59abbe56e057f20f883e"
	testParam["stime"] = "1660364991"
	testParam["aa"] = "vv"

	sign := CheckSign(testParam, "123", "aa")

	if sign == "" {
		t.Error("error")
	}
	if sign != signEx {
		fmt.Println(sign)
		t.Error("error")
	}
}
