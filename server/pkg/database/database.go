package database

import (
	"fmt"
	"pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conn() (db *gorm.DB, er error) {

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := config.GetDatabaseDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Printf("Conn db=%v, err=%v\n", db, err)
	er = err
	return db, er
}

type Account struct {
	Accountid  string `json:accountid`
	Username   string `json:username` //用户名',
	Password   string `json:password` //密码',
	Corpid     int    `json:corpid`   //企业编号',
	CreateTime string `gorm:"type:datetime" json:"createTime"`
	LastLogin  string `json:lastLogin`
	Isdelete   string `json:isdelete`
}

func GetUser(name string, pwd string) {

	db, err1 := Conn()
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	account := new(Account)
	res := db.Where("username=?", name).Where("password=?", pwd).First(account)

	fmt.Println(res.Error)        //判断返回值的错误
	fmt.Println(res.RowsAffected) //查看返回条数
	fmt.Println("查询到的对象为", *account)
	fmt.Println("查询到的对象为", account) //相对于上面
}
