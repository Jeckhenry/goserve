package serveInfo

import (
	"encoding/json"
	"fmt"
	"log"
	"mysqlInfo"
	"net/http"
)
//操作数据库返回信息
type Reback struct {
	Code int `json:"code"`
	Message string `json:"message"`
}


func StartServe(){
	//标签新增/更新接口
	http.HandleFunc("/addLabel", func(writer http.ResponseWriter, request *http.Request) {
		var labelname string
		labelname = request.FormValue("labelName")
		labelId := request.FormValue("id")
		var sqlword string
		if (labelId != "") {
			sqlword = "update labels set labelName='"+labelname+"' where labelId='"+labelId+"'"
		}else {
			sqlword = "insert into labels(labelName) values('"+labelname+"')"
		}
		_,err := mysqlInfo.DB.Exec(sqlword)
		var res Reback
		if err != nil {
			res = Reback{Code: 500,Message:"操作失败"}
		}else {
			res = Reback{Code:200,Message:"操作成功"}
		}
		b,_ := json.Marshal(res)
		fmt.Fprint(writer,string(b))
		return
	})
	//查询标签信息
	http.HandleFunc("/labelInfo", func(writer http.ResponseWriter, request *http.Request) {
		sqlWord := "select labelId,labelName,isUse from labels"
		rows,err := mysqlInfo.DB.Query(sqlWord)
		var res Reback
		if err != nil {
			panic(err)
			res = Reback{Code: 500,Message:"操作失败"}
			b,_ := json.Marshal(res)
			fmt.Fprint(writer,string(b))
			return
		}
		resData := make(map[string]interface{})
		var dataArr []interface{}
		for rows.Next() {
			single := make(map[string]interface{})
			var labelname string
			var isUse bool
			var labelId int
			err := rows.Scan(&labelId,&labelname,&isUse)
			if err != nil {
				log.Fatal(err)
			}
			single["labelname"] = labelname
			single["isUse"] = isUse
			single["id"] = &labelId
			dataArr = append(dataArr,single)
		}
		resData["code"] = 200
		resData["message"] = "操作成功"
		resData["data"] = dataArr
		b,err := json.Marshal(resData)
		if err != nil {
			panic(err)
		}
		fmt.Fprint(writer,string(b))
	})
	//标签删除
	http.HandleFunc("/delLabel", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Print("llllll")
		idarr := request.FormValue("ids")
		fmt.Println(idarr)
	})
	//开始监听
	fmt.Println("端口9095服务已启动.....")
	err := http.ListenAndServe(":9095",nil)
	if err != nil {
		fmt.Println("服务启动失败",err)
	}
}
