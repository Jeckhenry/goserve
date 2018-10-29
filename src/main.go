package main

import (
	_ "github.com/go-sql-driver/mysql"
	"mysqlInfo"
	"serveInfo"
	"sqlAddtional"
)

func main()  {
	sqlAddtional.NowDate()
	mysqlInfo.StartSql()
	serveInfo.StartServe()
}
