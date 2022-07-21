package main

import (
	"fmt"
)

func main() {
	//byte和run类型

	//go语言中为了处理非ASCII码类型的字符  定义了新的rune类型

	s := "HELLO沙河안녕하십니까"
	//len()求的是byte字节的数量
	n := len(s) //求字符串s的长度,把长度保存到变量n中
	fmt.Println(n)

	// for i := 0; i < len(s); i++ {
	// 	//fmt.Println(s[i])
	// 	fmt.Printf("%c\n", s[i]) // %c:字符
	// }

	// for _, c := range s {
	// 	fmt.Printf("%c\n", c) //%c:字符
	// }

	//字符串修改
	s2 := "白萝卜"      // [白 萝 卜] 相当于转换为一个列表
	s3 := []rune(s2) //把字符串强制转换成了一个rune切片
	//s3[0] = "红"
	s3[0] = '红'
	fmt.Println(string(s3)) //把rune切片强制转换为字符串

}
