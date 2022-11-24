package main

import "fmt"

// 高阶函数返回值
func foo() func() int {
	var inner = func() int {
		fmt.Println("一个新的函数功能....")
		return 100
	}
	return inner
}

func main() {
	var f func() int
	f = foo() //返回inner函数体复制给f变量
	f()       //函数调用
	//直接调用
	foo()()

}
