package mysqlInfo

import (
	"database/sql"
	"fmt"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "123456"
	ip = "47.104.227.155"
	port = "3306"
	dbName = "blogserve"
)
//数据库连接池
var DB *sql.DB
//启动数据库
func StartSql(){
	//构建数据库连接
	path := strings.Join([]string{userName,":",password,"@tcp(",ip,":",port,")/",dbName,"?charset=utf8"},"")
	//打开数据库
	var err error
	DB,err = sql.Open("mysql",path)
	if err != nil {
		fmt.Println("启动数据库失败",err)
		return
	}
	//设置数据库最大连接
	DB.SetConnMaxLifetime(100)
	//设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err = DB.Ping(); err != nil {
		fmt.Println("数据库连接失败：",err)
		return
	}
	fmt.Println("数据库初始化成功.......")
}
