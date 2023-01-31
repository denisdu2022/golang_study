package main

import (
	"fmt"
	"net"
)

func main() {

	//服务端
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		//捕获错误
		panic("创建监听失败!")
	}

	//循环
	for true {
		//延迟注册关闭监听
		defer listen.Close()
		fmt.Println("Server is Waiting!!!")
		//获取客户端连接对象,三次握手,客户端未连接,呈现阻塞
		//阻塞接受
		//客户端套接字对象conn
		conn, err := listen.Accept()
		if err != nil {
			//捕获错误
			panic("连接失败!")
		}

		//基于conn进行通信
		//接收来自客户端浏览器的请求信息
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			//捕获错误
			panic("客户端连接失败!")
		}
		//data一定是要符合请求格式
		//打印客户端浏览器的请求信息
		fmt.Println("客户端发送来的数据: ", string(data[:n]))

		//服务端响应客户端
		//conn.Write([]byte("HTTP/1.1 200 ok\r\ncontent-type:text/plain\r\n\r\nhello Web"))
		//content-type:text/plain 并不会显示H1标签,而是一个完整的<h1>hello Web</h1>
		//conn.Write([]byte("HTTP/1.1 200 ok\r\ncontent-type:text/plain\r\n\r\n<h1>hello Web</h1>"))
		//修改为content-type:text/html则可以显示H1一级标题,客户端浏览器通过content-type来解析
		conn.Write([]byte("HTTP/1.1 200 ok\r\ncontent-type:text/html\r\n\r\n<h1>hello Web</h1>"))
		fmt.Println("end....")
	}
}
