package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// 处理文件上传
func upload(conn net.Conn, params []string) {

	//接收文件信息
	fileName := params[0]
	fileSize, _ := strconv.Atoi(params[1])
	fmt.Println("fileName: ", fileName)
	fmt.Println("fileSize: ", fileSize)

	//循环接收文件
	//创建文件句柄
	file, _ := os.OpenFile("uploads/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	//获取writer对象
	writer := bufio.NewWriter(file)

	var readSize = 0
	for readSize < fileSize {
		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		readSize += n
		//将每行数据写入到writer中
		writer.WriteString(string(data[:n]))
	}
	writer.Flush()
	fmt.Println("文件上传成功")

}

// 处理文件下载
func download(conn net.Conn, params []string) {

}

func main() {

	//1. 创建服务端套接字对象
	sock, err := net.Listen("tcp", "127.0.0.1:8000")
	//错误处理
	if err != nil {
		fmt.Println("errors: ", err)
	}
	//2. 延迟注册关闭套接字
	defer sock.Close()

	for true {
		//等待客户端连接提示
		fmt.Println("Server is Wating...")

		//3.接收客户端套接字对象
		conn, err1 := sock.Accept()
		//错误处理
		if err1 != nil {
			fmt.Println("errors: ", err)
		}
		fmt.Println(conn)
		for true {
			//服务器文件上传接收的功能
			//接收客户端传过来的info
			infoBytes := make([]byte, 1024)
			n, _ := conn.Read(infoBytes)
			info := string(infoBytes[:n])
			infoSlice := strings.Split(info, " ")
			cmd := infoSlice[0]
			params := infoSlice[1:]
			switch cmd {
			case "put":
				upload(conn, params)
			case "get":
				download(conn, params)
			}
		}
	}

}
