package main

import "fmt"

func main() {

	var msg = "abc哈"

	//字符串转换为字节串
	fmt.Println([]byte(msg)) //最常用的方式
	//[97 98 99 229 147 136]

	fmt.Println([]rune(msg))
	//[97 98 99 21704] Unicode的值
	fmt.Println(len([]rune(msg))) //统计字符的时候常用   4

	//字节串转换成字符串
	info1 := []byte(msg)
	info2 := []byte{97, 98, 99, 229, 147, 136}
	fmt.Println(info1)
	fmt.Println(info2)
	fmt.Println(string(info1))
	fmt.Println(string(info2))

}
