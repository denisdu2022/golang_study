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

//文件上传函数
func upload(conn net.Conn, path string) {
	fmt.Println("文件上传开始,path: ", path)

	//打开上传的文件
	uploadFile, _ := os.Open(path)
	fmt.Println("uploadFile是: ", uploadFile)

	//读取文件,发送给服务端

	//获取文件信息
	f, _ := os.Stat(path)
	fileSize := f.Size()
	fileName := f.Name()
	fmt.Println("文件的名字: ", fileName)
	fmt.Println("文件的大小: ", fileSize)
	//定义传给服务端的文件信息,拼接字符串,需要把fileSize的类型是int64强转为字符串才可以拼接
	Info := "put" + " " + fileName + " " + strconv.FormatInt(fileSize, 10)
	fmt.Println("fileInfo: ", Info)
	//上传文件名称和大小
	conn.Write([]byte(Info))

	//读取文件,获取文件阅读器
	reader := bufio.NewReader(uploadFile)
	//fmt.Println(reader)

	//循环上传文件数据
	for true {
		//读取文件每一行
		lineContent, err := reader.ReadBytes('\n')

		//发送给服务端
		conn.Write(lineContent)

		//读取文件结束后退出循环
		if err == io.EOF {
			break
		}
	}
	fmt.Println("文件上传成功!!!")

}

//文件下载函数
func download(path string) {

}

func main() {

	//1.客户端拨号
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	//错误处理
	if err != nil {
		fmt.Println("errors: ", err)
	}
	fmt.Println(conn)

	for true {
		//客户端文件上传
		//var cmdStr string
		//fmt.Println("请输入路径>>>")
		//
		//		fmt.Scan(&cmdStr)
		///*	输出结果,因为你Scan或者Scanln是以空格为分隔符的,所以输出结果是乱的
		//		请输入路径>>>
		//			put a.png
		//		put
		//		请输入路径>>>
		//			a.png
		//		请输入路径>>>*/
		//			fmt.Println(cmdStr)
		reader := bufio.NewReader(os.Stdin) //从终端控制行读取
		fmt.Println("请输入命令: ")
		content, _ := reader.ReadString('\n') //读到换行为止
		content = strings.TrimSpace(content)  //去除空格
		fmt.Println("content是: ", content)
		/*	输出结果
			请输入命令:
				put a.png
			content是:  put a.png*/
		cmd := strings.Split(content, " ")[0]    //分割取命令
		params := strings.Split(content, " ")[1] //分割取路径

		switch cmd {
		case "put":
			//文件上传
			upload(conn, params)
		case "get":
			//文件下载
			download(params)
		}

	}
}
