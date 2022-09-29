package main

import "fmt"

func main() {
	////01. defer语句
	//fmt.Println("test01")
	//defer fmt.Println("test02")
	//fmt.Println("test03")
	////执行结果
	///*	test01
	//	test03
	//	test02*/

	////02 defer语句延迟注册调用
	////打开文件
	//file_open, err := os.Open("纳兰词")
	//if err != nil {
	//	fmt.Println("文件打开失败,错误原因是: ", err)
	//}
	////延迟注册关闭文件
	//defer file_open.Close()

	//03 defer语句的执行顺序
	fmt.Println("test01...")
	defer fmt.Println("test02...")
	fmt.Println("test03...")
	defer fmt.Println("test04...")
	//输出结果
	/*	test01...
		test03...
		test04...
		test02...*/

}
