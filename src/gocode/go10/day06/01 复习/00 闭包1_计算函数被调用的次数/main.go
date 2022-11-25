package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// 统计函数调用次数的功能函数
func countFunc(f func()) func() {
	//定义函数调用次数变量,初始值为0
	var ciShu = 0
	return func() {
		f()
		//函数调用一次自加1
		ciShu++
		//把调用的函数名称赋值为一个变量fn
		fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		fmt.Printf("%s函数调用的次数:%d\n", fn, ciShu)
	}

}

// 测试函数
func foo() {
	fmt.Println("foo 函数功能开始...")
	fmt.Println("foo 函数功能结束...")
}

func bar() {
	fmt.Println("bar 函数功能开始...")
	fmt.Println("bar 函数功能结束...")
}

func main() {
	//把需要调用函数的当做实参传给统计函数调用次数的形参在赋值,赋值变量的名字使用调用函数的名字
	foo := countFunc(foo)
	foo()
	foo()
	foo()

	bar := countFunc(bar)
	bar()
	bar()
}
