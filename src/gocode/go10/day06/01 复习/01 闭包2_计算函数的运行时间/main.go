package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

//计算函数运行时间的函数
func countTimer(f func(t time.Duration)) func(t time.Duration) {
	return func(t time.Duration) {
		//时间戳开始
		timeStart := time.Now().Unix()
		f(t)
		//时间戳结束
		timeEnd := time.Now().Unix()
		fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		fmt.Printf("%s函数的运行时间是:%d\n", fn, timeEnd-timeStart)
	}

}

//测试函数
func foo(t time.Duration) {
	fmt.Println("foo 功能函数运行...")
	time.Sleep(time.Second * t)
	fmt.Println("foo 功能函数结束...")
}

func bar(t time.Duration) {
	fmt.Println("bar 功能函数运行***")
	time.Sleep(time.Second * t)
	fmt.Println("bar 功能函数结束***")
}

func main() {
	foo := countTimer(foo)
	foo(3)

	bar := countTimer(bar)
	bar(2)
}
