package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

// 处理文件上传的功能函数
func handleUpload(conn net.Conn, params []string) {
	//19:接收文件信息
	fileNmae := params[0]
	fileSize, err := strconv.Atoi(params[1]) //params文件大小是string，转为int
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是：", err)
	}
	//fmt.Println("文件名：", fileNmae)
	//fmt.Println("文件大小：", fileSize)
	//20. 创建文件句柄
	file, err1 := os.OpenFile("uploads/"+fileNmae, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	//错误信息处理
	if err1 != nil {
		fmt.Println("错误信息是：", err1)
	}
	//21.获取write对象,存储文件
	writer := bufio.NewWriter(file)

	//22.循环读取文件
	//定义初始的redsize值
	var readsize = 0
	for readsize < fileSize {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		//错误信息处理
		if err != nil {
			fmt.Println("错误信息是：", err)
		}
		readsize += n
		//将读取到的每一行写入writer对象中
		writer.WriteString(string(data[:n]))
		writer.Flush()
	}

	fmt.Println("文件上传成功")
}

// 处理文件下载的功能函数
func handleDownload(conn net.Conn, params []string) {
	//24.获取文件信息
	//[uploads/1.png]
	var path string
	for _, v := range params {
		path = v
	}
	//fmt.Println(path, reflect.TypeOf(path))
	//打开文件
	f, err := os.Stat(path)
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是：", err)
	}
	fileName := f.Name()
	fileSize := f.Size()
	fmt.Println("文件的名字: ", fileName)
	fmt.Println("文件的大小: ", fileSize)
	fileSize1 := strconv.FormatInt(fileSize, 10)
	fileInfo := fileName + " " + fileSize1
	//25.将获取到的文件信息发送给客户端
	conn.Write([]byte(fileInfo))

	//26.打开文件
	downloadFile, err1 := os.Open(path)
	//错误信息处理
	if err1 != nil {
		fmt.Println("错误信息是：", err1)
	}
	//获取文件阅读器
	reader := bufio.NewReader(downloadFile)
	//27.循环读取每一行
	for true {
		lineContent, err := reader.ReadBytes('\n')
		//错误信息处理
		if err != nil {
			fmt.Println("错误信息是：", err)
		}
		//发送给客户端
		conn.Write(lineContent)
		if err == io.EOF {
			fmt.Println("文件下载成功...")
			break
		}
	}

}

func main() {

	//1.创建服务端套接字对象
	sock, err := net.Listen("tcp", "127.0.0.1:8080")
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是：", err)
	}

	//2.延迟注册 关闭套接字对象
	defer sock.Close()

	//3.等待客户端连接信息
	fmt.Println("Server is Wating...")

	//4.接受客户端套接字对象
	conn, err1 := sock.Accept()
	//错误信息处理
	if err1 != nil {
		fmt.Println("错误信息是：", err1)
	}
	//fmt.Println(conn)

	//5.发消息给客户端
	conn.Write([]byte("Connect to Server!!!"))

	//9.接受客户端发来的消息
	clientWel := make([]byte, 1024)
	n, err2 := conn.Read(clientWel)
	if err2 != nil {
		fmt.Println("错误信息是：", err2)
	}
	fmt.Println(string(clientWel[:n]))

	for true {
		//18.服务端接收客户端拼接好的传送过来的字节串
		fileInfoBytes := make([]byte, 1024)
		n, err3 := conn.Read(fileInfoBytes)
		//错误信息处理
		if err3 != nil {
			fmt.Println("错误信息是：", err3)
		}
		//将接收的字节串转为字符串
		info := string(fileInfoBytes[:n])
		//fmt.Println(info) //put 1.png 71908
		//将字符串进行分割，以空格分割
		infoSlice := strings.Split(info, " ")
		//分割后是[]string
		//从索引0拿到执行的命令
		cmd := infoSlice[0]
		//从索引1拿到文件名和文件大小
		params := infoSlice[1:]
		//执行命令的Switch判断
		switch cmd {
		case "put":
			//处理文件上传
			handleUpload(conn, params)
		case "get":
			//处理文件下载
			handleDownload(conn, params)
		}

	}
	fmt.Println("test")
}
