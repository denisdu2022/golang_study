package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器准备开启...")

	//1.开启服务端监听
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	//错误处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
		return //如果出现了错误也要返回
	}
	//2.延迟注册关闭监听
	defer listener.Close()

	//服务端循环接收
	for {
		//3.阻塞接收
		conn, err := listener.Accept()
		//错误处理
		if err != nil {
			fmt.Println("错误信息是: ", err) //err == EOF客户端退出
		}
		//打印服务端的本地地址
		//fmt.Println(conn.LocalAddr())
		//打印服务端的远程连接地址
		fmt.Println(conn.RemoteAddr()) //127.0.0.1:50464 客户端的随机端口

		//14.使用协程
		go proc(conn)
	}

}

// 11.协程方法
func proc(conn net.Conn) {
	//12.延迟注册关闭conn
	defer conn.Close()

	for {
		//13.客户端接收 >>> 一收
		//初始化接收buff
		buf := make([]byte, 1024)
		//读取客户端发来的信息
		n, nerr := conn.Read(buf)
		//错误信息处理
		if nerr != nil {
			fmt.Println("错误信息是: ", nerr)
		}
		//n是长度,这里会有一个粘包粘包的问题
		fmt.Println(string(buf[:n]))

	}

}
