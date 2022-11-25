package main

import "fmt"

func main() {
	//十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%o\n", i1) //把十进制转为八进制
	fmt.Printf("%x\n", i1) //把十进制转为十六进制
	fmt.Printf("%b\n", i1) //把十进制转为二进制

	//八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	//十六进制
	i3 := 0x121212
	fmt.Printf("%d\n", i3)

	//查看变量类型
	fmt.Printf("%T\n", i1)

	//声明int8类型的变量
	i4 := int8(9) //明确指定数据类型,否则就是int类型
	fmt.Printf("%T\n", i4)
}
