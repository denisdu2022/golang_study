package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器准备就绪")

	//1.开启服务端监听
	sock, err := net.Listen("tcp", "127.0.0.1:8000")
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是>>> ", err)
		return
	}
	//2.延迟注册关闭监听
	defer sock.Close()

	for true {
		//3.阻塞接收客户端消息
		conn, err := sock.Accept()
		//错误信息处理
		if err != nil {
			fmt.Println("错误信息是: ", err)
			continue
		}
		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		fmt.Println("n>>> ", n)

		//4.写入客户端信息
		conn.Write([]byte(`HTTP/1.1 200 OK\r\n\r\n<h1>hello Web!</h1> <img src="https://gd-hbimg.huaban.com/c52266373df6e989a1e9b08931656cf79bc3fe28569b1-V6uBLB_fw658">`))
	}

}
