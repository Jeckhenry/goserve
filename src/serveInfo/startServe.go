package serveInfo

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"mysqlInfo"
	"net/http"
	"session"
	"sqlAddtional"
	"strconv"
	"strings"
)
//操作数据库返回信息
type Reback struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
type Server struct {
	ArticleId int
	LabelId int
}
type ServerSlice struct {
	Ids []Server
}
func StartServe(){
	//标签新增/更新接口
	http.HandleFunc("/addLabel", func(writer http.ResponseWriter, request *http.Request) {
		checkres := session.Test_session_valid(writer,request)
		if checkres == "err" {
			return
		}
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
		checkres := session.Test_session_valid(writer,request)
		if checkres == "err" {
			return
		}
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
		checkres := session.Test_session_valid(writer,request)
		if checkres == "err" {
			return
		}
		request.ParseForm()
		ids := request.FormValue("ids")
		idsmid := ids[1:len(ids)-1]
		idarr := strings.Split(idsmid,",")
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
	//文章信息录入/更新接口
	http.HandleFunc("/addArticle", func(writer http.ResponseWriter, request *http.Request) {
		checkres := session.Test_session_valid(writer,request)
		if checkres == "err" {
			return
		}
		articleName := request.FormValue("articleName")
		articleInfo := request.FormValue("articleInfo")
		labelId := request.FormValue("labelId")
		subdate := sqlAddtional.NowDate() //提交时间
		articleId := request.FormValue("articleId") //文章id，以此判断是否是修改操作
		var err1,err2 error
		var sqlword string
		if (articleId != "") {
			sqlword = "update articles set articleTitle='"+articleName+"'," +
				"articleInfo='"+articleInfo+"',labelId='"+labelId+"',subDate='"+subdate+"' where articleId='"+articleId+"'"
		}else {
			sqlword = "insert into articles(articleTitle,articleInfo,labelId," +
				"subDate) values('"+articleName+"','"+articleInfo+"','"+labelId+"'" +
				",'"+subdate+"')"
		}
		_,err1 = mysqlInfo.DB.Exec(sqlword)
		sqlword = "update labels set isUse=1 where labelId='"+labelId+"'"
		_,err2 = mysqlInfo.DB.Exec(sqlword)
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
		checkres := session.Test_session_valid(writer,request)
		if checkres == "err" {
			return
		}
		sqlword := "SELECT * FROM articles INNER JOIN labels USING (labelId)"
		rows,_ := mysqlInfo.DB.Query(sqlword)
		var err1 error
		res := make(map[string]interface{})
		var arr []interface{}
		for rows.Next()  {
			single := make(map[string]interface{})
			var articletitle,articleinfo,subdate,articlelabel string
			var articleid,labelid,isuse int
			err1 = rows.Scan(&labelid,&articletitle,&articleinfo,&articleid,&subdate,&articlelabel,&isuse)
			if err1 != nil {
				fmt.Println(err1,"哈哈哈")
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
	//文章信息删除接口
	http.HandleFunc("/delArticle", func(writer http.ResponseWriter, request *http.Request) {
		checkres := session.Test_session_valid(writer,request)
		if checkres == "err" {
			return
		}
		request.ParseForm()
		idsArr := request.FormValue("ids")
		idsArr = `{"ids":`+idsArr+`}`
		var s ServerSlice
		json.Unmarshal([]byte(idsArr),&s)
		var sqlword string
		var err error
		for _,val := range s.Ids {
			articleid := strconv.Itoa(val.ArticleId)
			//删除文章
			sqlword = "delete from articles where articleId='"+articleid+"'"
			_,err = mysqlInfo.DB.Exec(sqlword)
			if err != nil {
				log.Fatal(err)
				break
			}
			//检索文章的标签id是否有关联的文章，没有就将标签的isUse置为0
			labelid := strconv.Itoa(val.LabelId)
			sqlword = "select articleTitle from articles where labelId='"+labelid+"'"
			rows,_ := mysqlInfo.DB.Query(sqlword)
			if rows.Next() == false {
				sqlword = "update labels set isUse=0 where labelId='"+labelid+"'"
				mysqlInfo.DB.Exec(sqlword)
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
	//文章模糊查找接口
	http.HandleFunc("/searcharticle", func(writer http.ResponseWriter, request *http.Request) {
		checkres := session.Test_session_valid(writer,request)
		if checkres == "err" {
			return
		}
		title := request.FormValue("title")
		sqlword := "SELECT * FROM articles as a INNER JOIN labels USING (labelId) where a.articleTitle like '%"+title+"%'"
		rows,_ := mysqlInfo.DB.Query(sqlword)
		var err error
		res := make(map[string]interface{})
		var datas []interface{}
		for rows.Next() {
			single := make(map[string]interface{})
			var articletitle,articleinfo,subdate,articlelabel string
			var articleid,labelid,isuse int
			err = rows.Scan(&labelid,&articletitle,&articleinfo,&articleid,&subdate,&articlelabel,&isuse)
			if err != nil {
				fmt.Println(err)
				break
			}
			single["articleTitle"] = articletitle
			single["articleInfo"] = articleinfo
			single["articleId"] = articleid
			single["subDate"] = subdate
			single["labelId"] = labelid
			datas = append(datas,single)
		}
		if err != nil {
			res["code"] = 500
			res["message"] = "操作失败"
		}else{
			res["code"] = 200
			res["message"] = "操作成功"
			if len(datas) == 0 {
				res["data"] = make([]string,0)
			}else{
				res["data"] = datas
			}
		}
		b,_ := json.Marshal(res)
		fmt.Fprint(writer,string(b))
	})
	//用户登陆
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		session.Login(writer,request)
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t,err := template.ParseFiles("index.html")
		if err != nil {
			fmt.Print("未找到文件")
		}
		t.Execute(writer,nil)
	})
	//用户退出
	http.HandleFunc("/logout", func(writer http.ResponseWriter, request *http.Request) {
		session.Logout(writer,request)
	})


	//开始监听
	fmt.Println("端口9095服务已启动.....")
	err := http.ListenAndServe(":9095",nil)
	if err != nil {
		fmt.Println("服务启动失败",err)
	}
}
