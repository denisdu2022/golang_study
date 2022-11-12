package main

import "fmt"

func foo(ch chan int) {
	fmt.Println(<-ch)
}

func main() {

	//channel的引用类型
	chan01 := make(chan int, 3)
	chan01 <- 1
	chan01 <- 2
	chan01 <- 3
	chan02 := chan01 //拷贝地址
	fmt.Println(<-chan02)
	fmt.Println(<-chan02)
	chan02 <- 100
	fmt.Println(<-chan01) //3 ,先进先出
	//mt.Println(<-chan01) //100
	foo(chan01) //100
}
