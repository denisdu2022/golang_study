package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

// 同步锁
var wg sync.WaitGroup

func write(conn net.Conn) {
	defer wg.Done()
	//4.写消息
	for true {
		//从控制台读取标准输入
		//引导用户输入聊天内容
		reader := bufio.NewReader(os.Stdin)
		//读取到换行截止
		content, err := reader.ReadBytes('\n')
		//错误信息处理
		if err != nil {
			fmt.Println("错误信息是: ", err)
			return
		}
		//fmt.Println("输入的是>>> ", string(content))
		//发送数据
		conn.Write(content)

	}
}

func read(conn net.Conn) {
	defer wg.Done()
	//6.收消息
	for true {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		//错误信息处理
		if err != nil {
			fmt.Println("错误信息是: ", err)
			return
		}
		fmt.Println(string(data[:n]))
	}

}

func main() {

	//3.客户端创建套接字
	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
		return
	}
	//延迟注册关闭conn
	defer conn.Close()

	//fmt.Println(conn)

	wg.Add(2)

	//读消息
	go read(conn)

	//写消息
	go write(conn)

	wg.Wait()

}
