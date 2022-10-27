package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.OpenFile("渔家傲03", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败...")
	} else {
		fmt.Println("打开文件成功***")
	}
	fmt.Println("file: ", file)

	////写入文件
	////1. 按字符串或者字节写操作
	//file.Write([]byte("我报路长嗟日暮，学诗谩有惊人句。\n"))
	//file.WriteString("仿佛梦魂归帝所，闻天语，殷勤问我归何处。\n")

	////2. 缓冲写
	//writer := bufio.NewWriter(file)
	//writer.WriteString("我报路长嗟日暮，学诗谩有惊人句。\n")
	//writer.Flush()

	//3. 写入整个文件
	str := "仿佛梦魂归帝所，闻天语，殷勤问我归何处。\n"
	ioutil.WriteFile("渔家傲03", []byte(str), 0666)
}
