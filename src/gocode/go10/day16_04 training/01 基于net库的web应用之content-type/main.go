package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("Server is waiting!!!")

	//1.开启服务端监听
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	//错误处理
	if err != nil {
		panic("服务端开启监听失败...")
	}

	//2.延迟关闭监听
	defer listen.Close()

	//循环
	for true {
		//3.阻塞接収客户端消息
		conn, err := listen.Accept()
		//错误处理
		if err != nil {
			panic("接收客户端请求失败...")
		}
		//读取客户端信息,并打印
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		//错误处理
		if err != nil {
			panic("读取客户端请求失败...")
		}
		//打印客户端请求信息
		fmt.Println("客户端请求信息: ", string(data[:n]))

		//响应客户端请求
		//conn.Write([]byte(`HTTP/1.1 200 OK\r\ncontent-type:text/plain\r\n\r\n\r\n<h1>hello web111</h1>`))
		conn.Write([]byte(`HTTP/1.1 200 OK\r\ncontent-type:text/html\r\n\r\n<h1>hello web111</h1>`))
		fmt.Println("end................")
	}

}
