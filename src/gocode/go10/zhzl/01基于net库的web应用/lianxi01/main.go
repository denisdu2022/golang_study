package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("Server is Waiting!!!")

	//1.开启服务端监听
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	//错误信息处理
	if err != nil {
		panic("创建监听失败!")
	}
	//2.延迟注册关闭服务端监听
	defer listen.Close()

	//3.for循环
	for true {
		//4.服务端阻塞接受客户请求
		conn, err := listen.Accept()
		//错误信息处理
		if err != nil {
			panic("接收客户端请求失败")
		}
		//读取客户端的请求信息
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		//错误信息处理
		if err != nil {
			panic("接收客户端请求数据失败")
		}
		//打印客户端的请求数据
		fmt.Println("n>>> ", string(data[:n]))

		//服务端响应客户端请求
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n<h1>hello WEB!!!</h1>"))
	}

}
