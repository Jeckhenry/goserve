package main

import (
	_ "github.com/go-sql-driver/mysql"
	"mysqlInfo"
	"serveInfo"
	"session"
	"sqlAddtional"
)
var globalSessions *session.Manager
//然后在init函数中初始化
func init() {
	//var err error
	//globalSessions, err = session.NewManager("memory","gosessionid",3600)
	//fmt.Println(globalSessions,err)
	}
func main()  {
	sqlAddtional.NowDate()
	mysqlInfo.StartSql()
	serveInfo.StartServe()
}
