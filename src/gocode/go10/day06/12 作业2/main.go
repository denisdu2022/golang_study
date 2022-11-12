package main

import "fmt"

//func makeFun(i int) func() {
//	return func() {
//		fmt.Println(i)
//	}
//}

func makeFun2() func(i int) {
	return func(i int) {
		fmt.Println(i)
	}
}

func main() {
	////版本1
	//var fn [10]func()
	//
	//for i := 0; i < len(fn); i++ {
	//	fn[i] = func() {
	//		fmt.Println(i)
	//	}
	//}
	//for _, f := range fn {
	//	f()
	//}
	////输出结果打印是10个10

	////版本2
	//var fn [10]func()
	//var i int
	//for i = 0; i < len(fn); i++ {
	//	fn[i] = func() {
	//		fmt.Println(i)
	//	}
	//}
	////for _, f := range fn {
	////	f()
	////}
	//for i = 0; i < len(fn); i++ {
	//	fn[i]()
	//}
	////输出结果从0开始,小于len(fn)==10,打印结果从0-9

	////版本3
	//var fn [10]func()
	//for i := 0; i < len(fn); i++ {
	//	fn[i] = makeFun(i)
	//}
	//
	//for _, f := range fn {
	//	f()
	//}
	////输出结果0-9

	//版本4
	var fn [10]func(int)
	for i := 0; i < len(fn); i++ {
		fn[i] = makeFun2()
	}
	for i, f := range fn {
		f(i)
	}
	//输出结果还是0-9,因为f(i)是makeFun2返回的函数参数的i
}
