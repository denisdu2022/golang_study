package main

import (
	"fmt"
	"reflect"
)

type Msg struct {
	content string
	from    string
}

func main() {

	//chan 是go的协程之间通信的数据类型: 引用类型
	//引用类型:切片 map chan
	//var s []int
	var ch01 = make(chan int, 3) //使用make()初始化,int类型,容量为3
	//插入值 ch <- 值
	ch01 <- 11
	ch01 <- 12
	ch01 <- 23

	//取值
	fmt.Println(<-ch01)
	fmt.Println(<-ch01)
	fmt.Println(<-ch01)
	//fmt.Println(<-ch01) //取不到的时候,死锁 fatal error: all goroutines are asleep - deadlock!

	var ch02 = make(chan interface{}, 3)
	ch02 <- "hello"
	ch02 <- true
	ch02 <- Msg{"hello world", "daerwen"}

	//fmt.Println(<-ch02, reflect.TypeOf(<-ch02)) 这样是报错的
	//需要
	ch02t := <-ch02 //先赋值给一个变量,在打印类型
	fmt.Println(ch02t, reflect.TypeOf(ch02t))
	y := <-ch02
	fmt.Println(y, reflect.TypeOf(y))
	//第三个是一个结构体对象,是否可以.
	//fmt.Println((<-ch02).content) 是报错的不可以的
	//需要类型断言
	fmt.Println((<-ch02).(Msg).content)

}
