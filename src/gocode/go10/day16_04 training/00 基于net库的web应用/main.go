package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("Server is waiting !!!")

	//1.开启监听
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	//错误信息处理
	if err != nil {
		panic("开启服务器监听失败!!!")
	}

	//2.延迟注册关闭监听
	defer listen.Close()

	//循环
	for true {
		//3.服务端阻塞接收请求
		conn, err := listen.Accept()
		//错误信息处理
		if err != nil {
			panic("接收客户端信息失败!!!")
		}

		//打印客户端发来的请求
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		//错误信息处理
		if err != nil {
			panic("读取客户端信息失败!!!")
		}
		fmt.Println("客户端发来的请求数据: ", string(data[:n]))

		//响应  服务端响应客户端
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n\r\n hello web"))
		fmt.Println("end ............")
	}

}
