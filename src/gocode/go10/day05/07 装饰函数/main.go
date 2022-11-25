package main

import (
	"fmt"
	"time"
)

////装饰函数
////统计函数调用的次数
//func counter(f func()) func() int {
//	var sumCounter = 0
//	return func() int {
//		f()
//		sumCounter++
//		fmt.Println(f, "函数的调用次数: ", sumCounter)
//		return sumCounter
//	}
//
//}

////函数计数器
//func counter(f func()) func() int {
//	i := 0 //数据隔离
//	return func() int {
//		f()
//		i++
//		return i
//	}
//}
//
////测试调用的函数
//func foo() {
//	fmt.Println("执行fool函数的功能...")
//}
//
//func bar() {
//	fmt.Println("执行bar函数的功能***")
//}

// 计算函数运行时间的装饰函数
func timer(f func(x, y int)) func(int, int) {

	return func(x, y int) {
		timeS := time.Now().Unix()
		f(x, y)
		timeE := time.Now().Unix()
		fmt.Println("函数运行时间: ", timeE-timeS)
	}

}

// 测试调用的函数
// 相加
func add(x, y int) {
	time.Sleep(time.Second * 2)
	fmt.Println(x + y)
}

// 相乘
func call(x, y int) {
	time.Sleep(time.Second * 3)
	fmt.Println(x * y)
}

func main() {

	//foo := counter(foo)
	//foo()
	//foo()
	//foo()
	//
	//bar := counter(bar)
	//bar()
	//bar()

	//foo := counter(foo)
	//foo()
	//foo()
	//foo()
	//count := foo()
	//fmt.Println(count)

	//add(11, 12)
	//call(10, 20)

	//调用装饰函数
	add := timer(add)
	add(12, 23)

	call := timer(call)
	call(10, 20)

}
