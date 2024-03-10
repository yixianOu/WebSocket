package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/polaris1119/chatroom/global"
	"github.com/polaris1119/chatroom/logic"
)

func homeHandleFunc(w http.ResponseWriter, req *http.Request) {
	//解析浏览器的前端页面并渲染
	tpl, err := template.ParseFiles(global.RootDir + "/template/home.html")
	if err != nil {
		fmt.Fprint(w, "模板解析错误！")
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, "模板执行错误！")
		return
	}
}

func userListHandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//拿到用户组并解码为Byte类型
	userList := logic.Broadcaster.GetUserList()
	b, err := json.Marshal(userList)

	if err != nil {
		fmt.Fprint(w, `[]`)
	} else {
		fmt.Fprint(w, string(b))
	}
}
