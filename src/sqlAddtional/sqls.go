package sqlAddtional

import (
	"fmt"
	"time"
)

//获取服务器当前时间
func NowDate() string{
	nowtime := time.Now()
	fmt.Println(nowtime.Format("2006-01-02 15:04:05"))
	return nowtime.Format("2006-01-02 15:04:05")
}
