package main

import (
	"fmt"
	"net"
)

func main() {
	//基于net包的web应用的content-type扩展

	//1.开启服务端监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	//捕获错误
	if err != nil {
		panic("服务端开启监听失败!")
	}
	//2.延迟注册关闭客户端监听
	defer listener.Close()

	//循环
	for true {
		fmt.Println("Server is Waiting!!!")
		//3.阻塞接受客户端请求
		conn, err := listener.Accept()
		//捕获错误
		if err != nil {
			panic("接受客户端请求失败!")
		}
		//4.读取客户端请求
		//data一定要符合客户端请求格式
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		//捕获错误
		if err != nil {
			panic("读取客户端请求失败!")
		}
		//打印客户端请求
		fmt.Println("n>>> " + string(data[:n]))

		//5.服务端相应客户端请求
		//content-type
		//conn.Write([]byte("HTTP/1.1 200 OK\r\ncontent-type:text/plain\r\n\r\nhello web"))
		conn.Write([]byte("HTTP/1.1 200 OK\r\ncontent-type:text/html\r\n\r\nhello web"))
	}
}
