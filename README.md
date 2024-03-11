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
cmd：用于存放 main.main；
logic：用于存放项目核心业务逻辑代码
server：存放 server 相关代码，
template：存放静态模板文件；

本项目是对《Go 语言编程之旅：一起用 Go 做项目》 第四章聊天室的复现.
通过该项目的练习,可以学习到使用nhooyr网络库的网络编程知识以及协程之间利用channel进行通信的方法.

reference:
github.com/go-programming-tour-book/chatroom
