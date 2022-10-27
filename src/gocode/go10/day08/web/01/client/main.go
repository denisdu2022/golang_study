package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("客户端准备开启....")
	//01. 服务端拨号连接客户端
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	//错误处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
		return //如果出现了问题也要返回
	}
	//02. 延迟注册关闭
	defer conn.Close()
	fmt.Println("客户端连接成功...", conn.RemoteAddr().String())

}
