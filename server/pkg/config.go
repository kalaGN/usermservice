package config

import (
	"fmt"

	"github.com/tietang/props/ini"
)

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func LoadPort() (string, error) {
	file := "../v1/config.ini"
	conf := ini.NewIniFileConfigSource(file)
	return conf.Get("production.server.port")
}

func GetProduction() {
	file := "./config.ini"
	conf := ini.NewIniFileConfigSource(file)
	fmt.Println(conf.Get("production"))
}
