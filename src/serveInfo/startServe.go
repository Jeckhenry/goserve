package serveInfo

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"log"
	"mysqlInfo"
	"net/http"
	"sqlAddtional"
	"strings"
)
//操作数据库返回信息
type Reback struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
var decoder = schema.NewDecoder()

type delarr struct {
	ids []int
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
		if len(dataArr) == 0 {
			resData["data"] = make([]string,0)
		}else {
			resData["data"] = dataArr
		}
		b,err := json.Marshal(resData)
		if err != nil {
			panic(err)
		}
		fmt.Fprint(writer,string(b))
	})
	//标签删除
	http.HandleFunc("/delLabel", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		ids := request.FormValue("ids")
		ids = ids[1:len(ids)-1]
		idarr := strings.Split(ids,",")
		var sqlword string
		var err error
		for _,val := range idarr {
			sqlword = "delete from labels where labelId='"+val+"'"
			_,err = mysqlInfo.DB.Exec(sqlword)
			if err != nil {
				log.Fatal(err)
				break
			}
		}
		res := make(map[string]interface{})
		if err != nil {
			res["code"] = 500
			res["message"] = "操作失败"
		}else {
			res["code"] = 200
			res["message"] = "操作成功"
		}
		b,_ := json.Marshal(res)
		fmt.Fprint(writer,string(b))
	})
	//文章信息录入接口
	http.HandleFunc("/addArticle", func(writer http.ResponseWriter, request *http.Request) {
		articleName := request.FormValue("articleName")
		articleLabel := request.FormValue("articleLabel")
		articleInfo := request.FormValue("articleInfo")
		labelId := request.FormValue("labelId")
		subdate := sqlAddtional.NowDate() //提交时间
		sqlword := "insert into articles(articleTitle,articleInfo,articleLabel,labelId," +
			"subDate) values('"+articleName+"','"+articleInfo+"','"+articleLabel+"','"+labelId+"'" +
			",'"+subdate+"')"
		_,err1 := mysqlInfo.DB.Exec(sqlword)
		sqlword = "update labels set isUse=1 where labelId='"+labelId+"'"
		_,err2 := mysqlInfo.DB.Exec(sqlword)
		res := make(map[string]interface{})
		if err1 != nil || err2 != nil {
			log.Fatal(err1.Error(),err2.Error())
			res["code"] = 500
			res["message"] = "操作失败"
		}else{
			res["code"] = 200
			res["message"] = "操作成功"
		}
		b,_ := json.Marshal(res)
		fmt.Fprint(writer,string(b))
	})
	//文章信息查询接口
	http.HandleFunc("/articleInfo", func(writer http.ResponseWriter, request *http.Request) {
		sqlword := "select articleTitle,articleInfo,articleId,articleLabel," +
			"subDate,labelId from articles"
		rows,_ := mysqlInfo.DB.Query(sqlword)
		var err1 error
		res := make(map[string]interface{})
		var arr []interface{}
		for rows.Next()  {
			single := make(map[string]interface{})
			var articletitle,articleinfo,subdate,articlelabel string
			var articleid,labelid int
			err1 = rows.Scan(&articletitle,&articleinfo,&articleid,&articlelabel,&subdate,&labelid)
			if err1 != nil {
				fmt.Println(err1)
				break
			}
			single["articleTitle"] = articletitle
			single["articleInfo"] = articleinfo
			single["articleId"] = articleid
			single["articleLabel"] = articlelabel
			single["subDate"] = subdate
			single["labelId"] = labelid
			arr = append(arr,single)
		}
		if err1 != nil {
			res["code"] = 500
			res["message"] = "操作失败"
		}else{
			res["code"] = 200
			res["message"] = "操作成功"
			if len(arr) == 0 {
				res["data"] = make([]string,0)
			}else{
				res["data"] = arr
			}
		}
		b,_ := json.Marshal(res)
		fmt.Fprint(writer,string(b))

	})
	//开始监听
	fmt.Println("端口9095服务已启动.....")
	err := http.ListenAndServe(":9095",nil)
	if err != nil {
		fmt.Println("服务启动失败",err)
	}
}
