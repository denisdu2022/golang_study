package main

import (
	"fmt"
	"net"
)

func main() {

	//1. 客户端拨号
	//拨号获取客户端的套接字对象,建立连接
	//通过拨号把客户端的套接字对象conn传给了服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8000") //协议和address和服务端一致
	//错误处理
	if err != nil {
		fmt.Println("errors: ", err)
	}
	fmt.Println(conn)

	//2. 收发消息
	//make初始化 1KB
	data := make([]byte, 1024)
	//接收消息
	n, err := conn.Read(data) //n是长度
	//错误处理
	if err != nil {
		fmt.Println("errors: ", err)
	}
	fmt.Println(n)        //n是长度
	fmt.Println(data[:n]) //data是内容,这是字节  data[:n]  []byte
	fmt.Println(string(data[:n]))

	//2.客户端持续输入
	for true {
		//2.客户发送一个英文名服务端转成大写
		var name string
		fmt.Print("请输入一个英文名称 [q: 退出]: ")
		fmt.Scanln(&name)

		//2发
		conn.Write([]byte(name))

		//2发 判断name == "q" 退出
		if name == "q" {
			break //退出当前循环
		}

		//2发,接收客户端发来的转大写的英文名字
		data2 := make([]byte, 1024)
		n2, err2 := conn.Read(data2)
		//错误处理
		if err2 != nil {
			fmt.Println("errors: ", err)
		}
		fmt.Println("upperName: ", string(data2[:n2]))
	}

}
