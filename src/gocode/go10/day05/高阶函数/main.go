package main

import (
	"fmt"
	"time"
)

////函数作为形参
//func foo() {
//	fmt.Println("foor函数功能....")
//}
//
//func bar(f func()) {
//	fmt.Println(f) //打印形参函数地址 0x86b4e0
//	f()            //调用形参函数  foor函数功能....
//}

//foo功能函数
func foo() {
	fmt.Println("foo 功能开始.....")
	time.Sleep(time.Second * 2)
	fmt.Println("foo 功能结束.....")
}

//bar功能函数
func bar() {
	fmt.Println("bar 功能开始.....")
	time.Sleep(time.Second * 3)
	fmt.Println("bar 功能结束.....")
}

//运行时间函数
func timer(f func()) {
	start := time.Now().Unix()
	f()
	end := time.Now().Unix()
	fmt.Println(end - start)
}

func main() {
	//调用运行时间函数
	timer(foo)
	timer(bar)

	////调用bar函数
	//bar(foo)
	//
	////时间戳
	//t := time.Now().Unix()
	//fmt.Println(t, reflect.TypeOf(t)) //1661483023 int64
}
