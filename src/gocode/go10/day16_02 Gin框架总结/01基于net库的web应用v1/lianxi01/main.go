package main

import (
	"fmt"
	"net"
)

func main() {

	//1.开启服务端监听
	sock, err := net.Listen("tcp", "127.0.0.1:8080")
	//错误信息处理
	if err != nil {
		panic("开启服务端监听失败!")
	}

	//2.for循环
	for true {
		fmt.Println("Server is Waiting!!!")
		//end: 延迟注册关闭客户端监听
		defer sock.Close()
		//3.阻塞接收客户端请求
		conn, err := sock.Accept()
		//错误信息处理
		if err != nil {
			panic("接收客户端请求失败!")
		}
		//读取客户端请求信息
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		//错误信息处理
		if err != nil {
			panic("读取客户端请求发来的信息")
		}
		//打印客户端请求信息
		fmt.Println("n>>> ", string(data[:n]))
		//4.服务端响应客户端请求信息
		//conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n<h1>Hello Web!!!</h1>"))
		//conn.Write([]byte(`HTTP/1.1 200 OK\r\n\r\n<h1>风景图片</h1><img src="https://cdn.pixabay.com/photo/2023/01/14/19/50/flower-7718952_1280.jpg">`))
		//content-type
		//conn.Write([]byte("HTTP/1.1 200 OK\r\ncontent-type:text\r\n\r\n<h1>Hello Web!!!</h1>"))
		conn.Write([]byte("HTTP/1.1 200 OK\r\ncontent-type:text/html\r\n\r\n<h1>Hello Web!!!</h1>"))

	}

}
