package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tietang/props/ini"
)

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func LoadPort() (string, error) {
	apath := getCurrentAbPath()
	file := apath + "/../../v1/config.ini"
	conf := ini.NewIniFileConfigSource(file)
	return conf.Get("production.server.port")
}

func GetProduction() {
	file := "./config.ini"
	conf := ini.NewIniFileConfigSource(file)
	fmt.Println(conf.Get("production"))
}

// root:pwd@tcp(ip:3306)/test
func GetDatabaseDsn() string {
	file := "../v1/config.ini"
	conf := ini.NewIniFileConfigSource(file)
	users, _ := conf.Get("production.database.user")
	passwd, _ := conf.Get("production.database.password")
	host, _ := conf.Get("production.database.host")
	port, _ := conf.Get("production.database.port")
	db, _ := conf.Get("production.database.dbname")
	return users + ":" + passwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
}

//   获取当前路径
//   参考 https://www.cnblogs.com/zhaoyingjie/p/16038348.html
//   最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
