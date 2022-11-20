package main

import (
	"fmt"
	"net"
)

func main() {

	//服务端
	listen, err := net.Listen("tcp", "127.0.0.1:8090")
	//错误信息处理
	if err != nil {
		panic("创建失败!")
	}
	for true {
		fmt.Println("Server is wating!!!")
		//获取客户端连接套接字对象,三次握手,如果客户端未连接,呈现阻塞
		conn, err := listen.Accept()
		//错误信息处理
		if err != nil {
			panic("连接失败!")
		}

		//基于conn进行通信
		//接收来自客户端浏览器的请求信息
		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		//data符合http请求格式
		fmt.Println("data:\n", string(data[:n]))
		//响应必须符合http协议
		//res := "hello world"
		res := "HTTP/1.1 200 OK\r\ncontent-type:text/plain\r\n\r\nhello world"
		conn.Write([]byte(res))
		conn.Close()
	}
}
