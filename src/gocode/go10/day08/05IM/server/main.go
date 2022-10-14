package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("服务器准备开启...")

	//1.创建服务端监听
	sock, err := net.Listen("tcp", "127.0.0.1:8081")
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
		return
	}
	//2.延迟注册关闭监听
	defer sock.Close()

	//循环接收客户端发来的消息
	for true {
		//4.服务端阻塞接收客户端发来的消息
		conn, cerr := sock.Accept()
		//错误信息处理
		if cerr != nil {
			fmt.Println("错误信息是: ", cerr)
			return
		}
		//6.接收客户端发来的信息 >>> 一收
		clToSe := make([]byte, 1024)
		n, cerr1 := conn.Read(clToSe)
		//错误信息处理
		if cerr1 != nil {
			fmt.Println("错误信息是: ", cerr1)
			return
		}
		fmt.Println(string(clToSe[:n]))

		//7.发送消息给客户端
		conn.Write([]byte("Server connect client..."))

		//使用服务端处理函数
		go ClientConn(conn)

	}

}

// 服务端处理函数
func ClientConn(conn net.Conn) {
	//获取客户端地址
	ipAddr := conn.RemoteAddr().String()
	fmt.Println(ipAddr, "连接成功...")

	//创建接收缓冲区
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		//错误信息处理
		if err != nil {
			fmt.Println("错误信息是: ", err)
			return
		}
		//接收收到的数据
		//fmt.Println(string(buf[:n]))

		//切出有效的接收数据
		result := buf[:n]
		fmt.Printf("接收的数据,来自%-8s %d:%s\n", ipAddr, n, string(result))

		//接收到如果是输入exit 退出客户端连接
		if string(result) == "exit" {
			fmt.Println(ipAddr, "退出连接...")
			return
		}

		//回复消息给客户端
		conn.Write([]byte(strings.ToUpper(string(result))))
	}

}
