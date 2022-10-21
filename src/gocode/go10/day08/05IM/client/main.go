package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("客户端准备开启...")

	//2.客户端拨号连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)

	}

	for {
		//3. 延迟注册关闭conn
		defer conn.Close()
		//5.发送信息给服务端  >>>>一发
		conn.Write([]byte("Client to server...."))

		//8. 客户端接收服务端发来的消息
		seToCl := make([]byte, 1024)
		n, serr := conn.Read(seToCl)
		//错误信息处理
		if serr != nil {
			fmt.Println("错误信息是: ", serr)

		}
		fmt.Println(string(seToCl[:n]))

		//计数器
		//wg.Add(1)

		//聊天室功能
		//go im(conn)
		//wg.Wait()
		im(conn)
	}

}

func im(conn net.Conn) {
	//defer wg.Done()
	for {
		//创建缓冲区
		buf := make([]byte, 1024)
		fmt.Print("请输入发送的内容: ")
		fmt.Scan(&buf)
		fmt.Printf("发送的内容是: %s\n", string(buf))
		//发送数据
		conn.Write(buf)
		//如果输入的是exit就退出
		if string(buf) == "exit" {
			fmt.Println("您已退出聊天....")
			break
		}
		//接收服务端的回复
		res := make([]byte, 1024)
		n, rerr := conn.Read(res)
		//错误信息处理
		if rerr != nil {
			fmt.Println("错误信息是: ", rerr)
			return
		}
		result := res[:n]
		fmt.Printf("接收到数据:%s\n", string(result))

	}

}
