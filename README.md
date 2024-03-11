# WebSocket
WebSocket 是一种与 HTTP 不同的协议。两者都位于 OSI 模型的应用层，并且都依赖于传输层的 TCP 协议。为了与 HTTP 协议兼容，WebSocket 握手使用 HTTP Upgrade 头从 HTTP 协议更改为 WebSocket 协议。
WebSocket 提供全双工通信,并且可以在 TCP 之上启用消息流,而TCP 单独处理字节流，没有固有的消息概念。

聊天室的功能主要有：
* 聊天室在线人数、在线列表
* 昵称设置
* 发送/接收消息
* 进入、退出聊天室
* 敏感词过滤
* 提醒或私聊

相关目录说明如下：
* cmd：用于存放 main.main；
* logic：用于存放项目核心业务逻辑代码；
1)broadcast.go: 负责将用户发送的消息广播给聊天室里的其他人,单例broadcastor在start()协程内利用 for- select 各个channel,达到广播消息的目的.
 同时还包含CanEnterRoom()和GetUserList(),用于判断用户名是否重复和获取用户列表;
2)user.go: 包含user类型,其含有SendMessage(),ReceiveMessage()等方法,以便广播器调用;
* server：服务端按流程调用业务;
1)handle.go 包含RegisterHandle(),用于开启广播器和注册handle;
2)home.go 包含homeHandleFunc()和userListHandleFunc,用于渲染页面和获取用户列表并显示;
3)websocket.go  包含Accept(),从前端接收WebSocket握手并升级连接为WebSocket.负责构建用户实例并提供服务,如启动协程SendMessage()向用户发送消息等
* template：存放前端静态模板文件；
* global：预先准备好环境;
1）init.go 包含inferRootDir函数和init()函数,用来推断根目录,初始化全局变量;

本项目是对《Go 语言编程之旅：一起用 Go 做项目》 第四章聊天室的复现.
通过该项目的练习,可以学习到使用nhooyr网络库的网络编程知识以及协程之间利用channel进行通信的方法.

reference:
github.com/go-programming-tour-book/chatroom
