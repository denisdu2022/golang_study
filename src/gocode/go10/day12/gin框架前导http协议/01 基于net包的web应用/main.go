package main

import (
	"fmt"
	"net"
)

func main() {
	//基于net包的web应用程序
	fmt.Println("服务端已开启...")

	//1.开启监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	//错误信息处理
	if err != nil {
		panic("开启失败")
	}

	//2.延迟注册关闭监听
	defer listener.Close()

	for {
		//3.阻塞等待接收
		conn, err1 := listener.Accept()
		//错误信息处理
		if err1 != nil {
			fmt.Println("错误信息是: ", err1)
			continue
		}
		//4.读取请求信息
		data := make([]byte, 1024)
		n, err2 := conn.Read(data)
		//错误信息处理
		if err2 != nil {
			fmt.Println("错误信息是: ", err2)
			continue
		}
		fmt.Println("客户端的请求信息是:\n", string(data[:n]))

		//5.响应客户端,向客户端发送信息
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n<h1>Hello web!!!</h1> "))
	}

}
