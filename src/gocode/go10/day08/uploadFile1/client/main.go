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

// 文件上传功能函数
func upload(conn net.Conn, path string) {
	fmt.Println("------------------文件上传开始------------------")
	//fmt.Println("path是: ", path)
	//16.获取文件信息
	f, err2 := os.Stat(path)
	if err2 != nil {
		fmt.Println("错误信息是: ", err2)
	}
	fileName := f.Name()
	fileSize := f.Size()
	fmt.Println("文件的名字: ", fileName)
	fmt.Println("文件的大小: ", fileSize)
	//17.发送给服务端,拼接字符串
	uploadInfo := "put" + " " + fileName + " " + strconv.FormatInt(fileSize, 10)
	//fmt.Println(uploadInfo)
	//将拼接好的字符串发送给服务端
	conn.Write([]byte(uploadInfo))

	//13. 打开上传的文件
	uploadFile, err := os.Open(path)
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
	}

	//14.读取文件
	reader := bufio.NewReader(uploadFile) //获取文件阅读器

	//循环读取文件每一行
	for true {
		//读取文件每一行
		lineContent, err1 := reader.ReadBytes('\n') ///n读到换行
		//错误信息处理
		if err1 != nil {
			fmt.Println("错误信息是: ", err1)
		}

		//15.发送给服务端
		conn.Write(lineContent)

		//读取文件结束标志结束读取
		if err1 == io.EOF {
			fmt.Println("------------------文件上传成功------------------")
			break
		}

	}

}

// 文件下载功能函数
func download(conn net.Conn, path string) {

	fmt.Println("------------------文件下载开始------------------")
	//fmt.Println("path是: ", path)
	//23.字符串拼接
	getInfo := "get" + " " + path
	fmt.Println(getInfo)
	//发送给服务端:文件命令和文件路径
	conn.Write([]byte(getInfo))
	//接收服务端发来的文件信息
	fileInfoBytes := make([]byte, 1024)
	n, err := conn.Read(fileInfoBytes)
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
	}
	//fmt.Println(string(fileInfoBytes[:n]))
	//将接收的字节串转为字符串
	fileInfoString := string(fileInfoBytes[:n])
	f := strings.Split(fileInfoString, " ")

	fileName := f[0]
	fmt.Println(fileName)
	//fmt.Println(f[1])
	fileSize := f[1]
	fmt.Println(fileSize)
	fileSize1, _ := strconv.Atoi(fileSize)

	//fmt.Println("fileName: ", fileName, reflect.TypeOf(fileName))
	//fmt.Println("fileSize: ", fileSize, reflect.TypeOf(fileSize))
	//28.创建文件句柄
	file, err1 := os.OpenFile("imgs/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	//错误信息处理
	if err1 != nil {
		fmt.Println("错误信息是: ", err1)
	}
	//获取write对象
	writer := bufio.NewWriter(file)
	//29.循环接收文件
	var readsize = 0
	for readsize < fileSize1 {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		//错误信息处理
		if err != nil {
			fmt.Println("错误信息是: ", err)
		}
		readsize += n
		writer.WriteString(string(data[:n]))
		writer.Flush()
	}

}

func main() {

	//6. 客户端拨号连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
	}
	fmt.Println(conn)

	//7.客户端接收来自服务端的信息
	serverWel := make([]byte, 1024)
	n, err1 := conn.Read(serverWel)
	//错误信息处理
	if err1 != nil {
		fmt.Println("错误信息是: ", err1)
	}
	fmt.Println(string(serverWel[:n]))

	//8. 客户端发消息给服务端
	conn.Write([]byte("Connect to Client***"))

	for true {

		//10.从终端控制台读取输入信息
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("请输入命令>>> ")
		content, err2 := reader.ReadString('\n') //读到换行为止
		//错误信息处理
		if err2 != nil {
			fmt.Println("错误信息是: ", err2)
		}
		content = strings.TrimSpace(content) //去除可能多输入的空格
		fmt.Println(content)
		//11.对输入的信息进行分割
		cmd := strings.Split(content, " ")[0]    //分割第1个元素取命令
		params := strings.Split(content, " ")[1] //分割第2个元素取路径

		//12.对输入的命令进行判断
		switch cmd {
		case "put":
			//文件上传
			upload(conn, params)
		case "get":
			//文件下载
			download(conn, params)
		}

	}

}
