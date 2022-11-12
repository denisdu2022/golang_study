package main

import (
	"fmt"
	"time"
)

//func counter(f func()) func() {
//	var conut = 0
//	return func() {
//		f()
//		conut++
//		fmt.Println("函数被调用次数: ", conut)
//	}
//}

func foo() {
	fmt.Println("foo函数功能执行开始...")
	time.Sleep(time.Second * 3)
	fmt.Println("foo函数功能执行结束...")
}

func bar() {
	fmt.Println("bar函数功能执行开始***")
	time.Sleep(time.Second * 5)
	fmt.Println("bar函数功能执行结束***")
}

//统计运行时间的函数
func timer(f func()) func() {
	return func() {
		timeStart := time.Now().Unix()
		f()
		timeEnd := time.Now().Unix()
		fmt.Println("函数功能的运行时间", timeEnd-timeStart)
	}

}

func main() {
	////foo()
	//cou := counter(foo)
	//cou()
	//cou()

	foo := timer(foo)
	foo()
	fmt.Println("-------------------------------------------")
	bar := timer(bar)
	bar()
}
