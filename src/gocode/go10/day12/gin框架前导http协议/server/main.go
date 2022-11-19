package main

import "net"

func main() {
	//1.开启服务端监听
	listten, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic("连接创建失败")
	}

	listten.Accept()
}
