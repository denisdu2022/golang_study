package main

import "fmt"

func main() {
	fmt.Println()

	//fmt占位符
	var n = 100
	//查看数据类型
	fmt.Printf("%T\n", n)
	//查看二进制
	fmt.Printf("%b\n", n)
	//查看值
	fmt.Printf("%v\n", n)
	//查看十进制
	fmt.Printf("%d\n", n)
	//查看八进制
	fmt.Printf("%o\n", n)
	//查看16进制
	fmt.Printf("%x\n", n)

	var s = "哈哈哈哈哈"
	//查看字符串
	fmt.Printf("%s\n", s)
	fmt.Printf("%#v\n", s) //#增加一个数据类型的描述符

}
