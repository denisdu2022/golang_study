package main

import (
	"fmt"
	"net"
)

func main() {

	//基于net包的web应用

	//1.开启服务端监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	//错误信息处理
	if err != nil {
		panic("开启服务端监听失败!")
	}

	//2.延迟注册关闭客户端监听
	defer listener.Close()

	//循环
	for true {
		fmt.Println("Server is waiting!")
		//3.阻塞接收客户端请求
		//获取客户端连接对象,三次握手,客户端未连接呈现阻塞
		//客户端套接字对象conn
		conn, err := listener.Accept()
		//错误信息处理
		if err != nil {
			panic("接收客户端请求失败!")
		}
		//4.读取客户端请求
		//基于客户端套接字对象conn进行通信
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		//错误信息处理
		if err != nil {
			panic("读取客户端请求失败!")
		}
		//打印客户端请求信息
		fmt.Println(string(data[:n]))

		//5.服务端相应客户端请求
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n<h1>Hello WEB!</h1>"))
	}
}
