package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("客户端准备开启...")
	//4.客户端拨号连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
	}

	//10.使用协程
	go proc(conn)
	if <-exitChan {
		fmt.Println("客户端退出连接...")
		return
	}

}

// 退出协程标记
var exitChan chan bool = make(chan bool, 1)

// 9.协程方法
func proc(conn net.Conn) {
	//5.延迟注册关闭客户端拨号
	defer conn.Close()
	//7.向服务端写数据
	//从当前命令行读取输入
	fmt.Print("请输入信息>>> ")
	//标准输入流
	reader := bufio.NewReader(os.Stdin)

	//15,循环多行读写
	for {
		//每次读一行,读到换行\n结束
		line, lerr := reader.ReadString('\n')
		//处理输入的字符串
		strings.Trim(line, "\r\n")
		//错误信息处理
		if lerr != nil {
			fmt.Println("错误信息是: ", lerr)
		}
		//16.客户端退出连接判断
		if line == "bye" {
			exitChan <- true
			break
		}
		//8.通过conn写入服务端 ,>>>>>> 一发
		n, nerr := conn.Write([]byte(line))
		//错误信息处理
		if nerr != nil {
			fmt.Println("错误信息是: ", nerr)
		}
		//n是长度
		//打印输入的信息
		//fmt.Println(string(line[:n]))
		//6.客户端连接成功
		fmt.Println("客户端连接成功: ", conn.RemoteAddr().String(), "并且写了:", n, "个字节...")

	}

}
