package main

import (
	_ "github.com/go-sql-driver/mysql"
	"mysqlInfo"
	"serveInfo"
	"session"
	"sqlAddtional"
)

func main()  {
	session.InintSession()
	sqlAddtional.NowDate()
	mysqlInfo.StartSql()
	serveInfo.StartServe()
}
