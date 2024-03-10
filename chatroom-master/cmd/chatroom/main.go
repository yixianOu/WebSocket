package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/polaris1119/chatroom/global"
	"github.com/polaris1119/chatroom/server"
)

var (
	addr = ":2022"
)

func init() {
	//推断出项目的根目录并对各种全局变量进行初始化
	global.Init()
}

func main() {
	fmt.Printf(addr)
	//注册socket并bind
	server.RegisterHandle()
	//监听并服务
	log.Fatal(http.ListenAndServe(addr, nil))
}
