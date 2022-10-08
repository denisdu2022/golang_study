package main

import (
	"fmt"
	"net"
)

func main() {

	//6. 客户端拨号连接服务端
	conn, err := net.Dial("tcp", "114.67.92.28:8080")
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
	}
	fmt.Println(conn)

	//7.客户端接收来自服务端的信息
	serverWel := make([]byte, 1024)
	n, err1 := conn.Read(serverWel)
	//错误信息处理
	if err1 != nil {
		fmt.Println("错误信息是: ", err1)
	}
	fmt.Println(string(serverWel[:n]))

	//8. 客户端发消息给服务端
	conn.Write([]byte("Connect to Client***"))
}
