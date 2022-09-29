package main

import (
	"fmt"
	"os"
)

func main() {

	////打开文件
	//file_open, err := os.OpenFile("蝶恋花", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	////判断文件是否存在
	//if err != nil {
	//	fmt.Println("文件打开失败,错误原因是: ", err)
	//}
	////关闭文件
	//defer file_open.Close()
	//
	//reader := bufio.NewReader(file_open)
	//wirter := bufio.NewWriter(file_open)
	//
	//for true {
	//	//1.按行读字符串
	//	strs, err := reader.ReadString('\n') //读取到换行符为止,读取内容包括换行符
	//	counter := strings.Trim(strs, "\n")
	//	s := fmt.Sprintf("\n改行的长度为%d,内容为: %s", len([]rune(counter)), counter)
	//
	//	//2.将改行数据追加进文件
	//	wirter.WriteString(s)
	//	wirter.Flush()
	//
	//	if err == io.EOF {
	//		break
	//	}
	//}

	////删除文件
	//os.Remove("蝶恋花1")

	////创建目录
	//dName := "DIELIANHUA"
	//os.Mkdir(dName, os.ModeDir|os.ModePerm)

	//获取文件信息
	fileInfo, err := os.Stat("蝶恋花2")
	if err == nil {
		fmt.Println("name: ", fileInfo.Name())
		fmt.Println("size: ", fileInfo.Size())
		fmt.Println("is dir: ", fileInfo.IsDir())
		fmt.Println("mod: ", fileInfo.Mode())
		fmt.Println("modTime: ", fileInfo.ModTime())
	}
	//输出结果
	//name:  蝶恋花2
	//size:  1441
	//is dir:  false
	//mod:  -rw-r--r--
	//modTime:  2022-09-05 17:11:26.318 +0800 CST

}
