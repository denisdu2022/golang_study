package main

import (
	"fmt"
)

// foo函数功能
func foo() {
	fmt.Println("执行fool 函数功能...")
}

// bar函数功能
func bar() {
	fmt.Println("执行bar 函数功能***")
}

func main() {
	//1. defer延迟注册
	//fmt.Println("test01")
	//foo()
	//fmt.Println("test02")
	//fmt.Println("test01")
	//defer foo() //延迟调用,当执行return后在执行defer 语句
	//fmt.Println("test02")
	///*	输出结果
	//	test01
	//	test02
	//	执行fool 函数功能...
	//*/

	////打开文件
	//file_obj, err := os.Open("满江红")
	////关闭文件
	//defer file_obj.Close()
	////判断文件是否存在
	//if err != nil {
	//	fmt.Println("文件打开失败,错误原因: ", err)
	//}
	//fmt.Println("文件打开成功: ", file_obj)

	//	//2. 多个defer调用
	//	fmt.Println("test01")
	//	defer foo() //先注册的最后调用
	//	fmt.Println("test02")
	//	defer bar() //延迟注册
	///*	输出结果
	//	test01
	//	test02
	//	执行bar 函数功能***
	//	执行fool 函数功能...
	//*/

	////3. defer的拷贝机制
	//foo := func() {
	//	fmt.Println("这是foo功能函数...")
	//}
	//defer foo() //会把第一次注册的语句体拷贝
	//foo = func() {
	//	fmt.Println("这是new foo功能函数...")
	//}
	////输出结果:fmt.Println("这是foo功能函数...")

	////拷贝机制案例1:
	//x := 10
	//defer func(a int) {
	//	fmt.Println(a)
	//}(x)
	//x++
	////输出结果: 10

	//拷贝机制案例2:
	//闭包or作用域
	x := 10
	defer func() {
		fmt.Println(x)
	}()
	x++
	//输出结果:11
}
