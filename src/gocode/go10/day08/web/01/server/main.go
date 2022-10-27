package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器准备开启....")
	//01. 开启服务端监听
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	//错误处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
		return //如果出现了问题也要返回
	}

	//02. 延迟注册关闭
	defer listener.Close()

	//03.循环阻塞接收
	for {
		conn, err := listener.Accept()
		//错误处理
		if err != nil {
			fmt.Println("错误信息是: ", err)
			return //如果出现了问题也要返回
		}
		fmt.Println(conn.LocalAddr())
	}
	fmt.Println("test")

}
