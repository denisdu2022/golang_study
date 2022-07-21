package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	//字符串
	//\本来是具有特殊含义的,我应该告诉程序我写的\就是一个单纯的\
	path := "D:\\Go\\src\\code\\main\\main.go"
	fmt.Println(path)

	//多行的字符串
	s2 := `
		哈哈哈哈
		嘿嘿嘿
		呵呵呵呵
	`

	fmt.Println(s2)
	s3 := `D:\Go\src\code\main\main.go`
	fmt.Println(s3)

	//字符串相关的操作
	//求字符串长度
	fmt.Println(len(s3))

	//字符串拼接
	//+号拼接
	fmt.Println(s3 + s2)
	//Sprintf拼接
	ss1 := fmt.Sprintf("%s%s", s2, s3)
	//fmt.Printf("%s%s", s2, s3)
	fmt.Println(ss1)

	//字符串分割
	ret := strings.Split(s3, "\\") //以\分割
	fmt.Println(ret)

	//字符串包含
	fmt.Println(strings.Contains(ss1, "嘿嘿"))
	fmt.Println(strings.Contains(ss1, "沙河"))

	//前缀
	fmt.Println(strings.HasPrefix(ss1, "哈"))

	//后缀
	fmt.Println(strings.HasSuffix(ss1, "go"))

	//字符串出现的位置
	ss2 := "abcdeb"
	fmt.Println(strings.Index(ss2, "d"))
	fmt.Println(strings.LastIndex(ss2, "b"))

	//拼接
	fmt.Println(strings.Join(ret, "+"))

}
