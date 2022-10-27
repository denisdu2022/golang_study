package main

import (
	"fmt"
	"net"
	"strings"
)

// 定义客户端的注册信息的数据结构
var onlineClients = make(map[string]net.Conn)

// 定义同步锁
//var wg sync.WaitGroup

// 消息广播函数
func broadCast(publisher string, content string) {

	//循环所有的客户端对象,返回信息
	for _, client := range onlineClients {
		//client.Write(data[:n])
		//client.Write([]byte(ip + ":" + string(data[:n])))
		client.Write([]byte(publisher + ":" + content))
		//fmt.Println(ip) //127.0.0.1:53638
	}

}

func read(conn net.Conn, addr string) {
	//defer wg.Done()
	for true {
		//5.收消息
		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		////错误信息处理
		//if err != nil {
		//	fmt.Println("错误信息是: ", err)
		//	return
		//}

		//下线处理
		if string(data[:n]) == "" {
			//下线操作
			delete(onlineClients, addr)
			//下线通知
			broadCast("[系统消息]", addr+"已经下线")
			return //不能使用break 因为下边的代码还要执行
		}

		//去除多余的空格
		content := strings.TrimSpace(string(data[:n]))
		fmt.Println(content)
		//调用消息广播
		broadCast(addr, content)
	}

}

func main() {

	fmt.Println("服务器准备就绪...")

	//1.创建服务端套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8899")
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
		return
	}

	//延迟注册关闭套接字
	defer listener.Close()

	for true {
		//2.阻塞等待客户端连接
		conn, cerr := listener.Accept()
		//错误信息处理
		if cerr != nil {
			fmt.Println("错误信息是: ", cerr)
			return
		}
		//注册conn到onlineClients
		addr := conn.RemoteAddr().String() //获取客户端IP
		fmt.Printf("来自客户端[%s]的连接\n", addr)
		//将客户端IP存入注册信息
		onlineClients[addr] = conn

		//fmt.Println(conn)

		//上线提醒
		broadCast("[系统消息]", addr+"已经上线了")

		//wg.Add(100)
		go read(conn, addr)
		//wg.Wait()

	}

}
