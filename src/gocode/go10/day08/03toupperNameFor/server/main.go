package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	//1. 服务端创建套接字对象
	//这个sock是服务端创建的套接字对象
	sock, err := net.Listen("tcp", "127.0.0.1:8000") //network是协议,address是IP地址:端口
	//错误处理
	if err != nil {
		fmt.Println("errors: ", err)
	}

	//2. 延迟注册关闭套接字 ,使用完毕后释放资源
	defer sock.Close()

	//3. 操作sock对象

	//(1). 等待客户端连接
	fmt.Println("Server is Waiting...")
	//conn是客户端的套接字对象

	for true {
		conn, err := sock.Accept() //如果没有客户端连接,会一直卡在这里,这是阻塞,会一直等待客户端连接
		//通过sock.Accept()来接收客户端的套接字对象
		if err != nil {
			fmt.Println("errors: ", err)
		}
		fmt.Println(conn)

		//(2). 收发消息: conn.read()收  conn.Write()发
		//time.Sleep(time.Second * 10)
		conn.Write([]byte("welcome to server!"))

		//2.服务端持续接收
		for true {
			//2收,接收客户端姓名
			data2 := make([]byte, 1024)
			n2, err2 := conn.Read(data2)
			//错误处理
			if err2 != nil {
				fmt.Println("errors: ", err)
			}
			fmt.Println(string(data2[:n2]))
			//判断name == "q" 退出
			if string(data2[:n2]) == "q" {
				break //退出当前循环
			}
			//定义转大写的变量接收
			var upperName = strings.ToUpper(string(data2[:n2]))
			fmt.Println("upperName: ", upperName)
			//在把转大写的发送给客户端
			//2服务端转为大写后发
			conn.Write([]byte(upperName))
		}

	}

}
