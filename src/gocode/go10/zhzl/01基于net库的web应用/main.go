package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("服务器准备就绪...")

	//1.开启服务端监听 端口:8000
	sock, err := net.Listen("tcp", "127.0.0.1:8080")
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
		return
	}
	//2.延迟注册关闭监听
	defer sock.Close()

	//3.循环
	for true {
		//阻塞接受客户端连接
		conn, err := sock.Accept()
		//错误信息处理
		if err != nil {
			fmt.Println("错误信息是: ", err)
			continue
		}
		//接收客户端发来的数据
		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		fmt.Println("n>>> ", n)

		//4.写入客户端信息
		conn.Write([]byte(`HTTP/1.1 200 OK \r\n\r\n hello Web!`))
		fmt.Println("-------")
		//test
	}
}
