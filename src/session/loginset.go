package session

import (
	"encoding/json"
	"fmt"
	"mysqlInfo"
	"net/http"
)

type UserInfo struct {
	UserId int
	UserName string
	UserPasswd string
}
var sessionMgr *SessionMgr //session管理器
var resmsg map[string]interface{}
//然后在init函数中初始化
func InintSession() {
	//创建session管理器,”TestCookieName”是浏览器中cookie的名字，3600是浏览器cookie的有效时间（秒）
	sessionMgr = NewSessionMgr("TestCookieName", 3600)
	fmt.Println("session启动")
}
//处理登录
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()

		//可以使用template.HTMLEscapeString()来避免用户进行js注入
		username := r.FormValue("username")
		password := r.FormValue("passwd")

		//在数据库中得到对应数据
		var userID int = 0

		userRow := mysqlInfo.DB.QueryRow("select userid from users where username=? and password=?", username, password)
		userRow.Scan(&userID)

		//TODO:判断用户名和密码
		if userID != 0 {
			//创建客户端对应cookie以及在服务器中进行记录
			var sessionID = sessionMgr.StartSession(w, r)
			var loginUserInfo = UserInfo{UserId: userID, UserName: username, UserPasswd: password}

			//踢除重复登录的
			var onlineSessionIDList = sessionMgr.GetSessionIDList()

			for _, onlineSessionID := range onlineSessionIDList {
				if userInfo, ok := sessionMgr.GetSessionVal(onlineSessionID, "UserInfo"); ok {
					if value, ok := userInfo.(UserInfo); ok {
						if value.UserId == userID {
							sessionMgr.EndSessionBy(onlineSessionID)
						}
					}
				}
			}

			//设置变量值
			sessionMgr.SetSessionVal(sessionID, "UserInfo", loginUserInfo)
			//TODO 设置其它数据

			//TODO 转向成功页面
			resmsg = make(map[string]interface{})
			resmsg["code"] = 200
			resmsg["message"] = "登陆成功"
			resmsg["data"] = loginUserInfo
		}else {
			resmsg = make(map[string]interface{})
			resmsg["code"] = 404
			resmsg["message"] = "用户不存在"
		}
	}
	b,_ := json.Marshal(resmsg)
	fmt.Fprint(w,string(b))
	return
}
//处理退出
func Logout(w http.ResponseWriter, r *http.Request) {
	sessionMgr.EndSession(w, r) //用户退出时删除对应session
	return
}
//在每个页面中进行用户合法性验证
func Test_session_valid(w http.ResponseWriter, r *http.Request) string{
	var sessionID = sessionMgr.CheckCookieValid(w, r)
	if sessionID == "" {
		resmsg = make(map[string]interface{})
		resmsg["code"] = 401
		resmsg["message"] = "用户未登录"
		data := make([]string,0)
		resmsg["data"] = data
		b,_ := json.Marshal(resmsg)
		fmt.Fprint(w,string(b))
		return "err"
	}
	return  ""
}