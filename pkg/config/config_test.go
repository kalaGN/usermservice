package config

import (
	"strconv"
	"strings"
	"testing"
)

// test demo  go test -v
func TestGetProduction(t *testing.T) {

	dsn := GetDatabaseDsn()

	if dsn == "" {
		t.Error("test not null")
	}
}

// TODO dsn format
func TestGetProductionFormat(t *testing.T) {

	dsn := GetDatabaseDsn()
	arr := strings.Split(dsn, "@")

	if arr[0] == "" || arr[1] == "" {
		t.Error("dsn error")
	}

	user := strings.Split(arr[0], ":")

	if user[0] == "" || user[1] == "" {
		t.Error("user or pwd get error")
	}

}

//不为空
func TestLoadPortGet(t *testing.T) {
	port, _ := LoadPort()
	if port == "" {
		t.Error("error LoadPort")
	}
}

//是正确的端口范围
func TestLoadPortRang(t *testing.T) {
	port, _ := LoadPort()
	portNum, err := strconv.Atoi(port)
	if err != nil {
		t.Error(err)
	}
	if portNum < 1 || portNum > 65535 {
		t.Error("port range error")
	}
}
