package main

import "fmt"

func main() {
	////1. channel的关闭
	//chan01 := make(chan int, 3)
	//chan01 <- 1
	//chan01 <- 2
	//close(chan01)
	////chan01 <- 3  //panic: send on closed channel 报错,关掉就不可以写了,只可以读
	//fmt.Println(<-chan01)
	//fmt.Println(<-chan01)

	////2. 判断管道是否关闭
	//chan01 := make(chan int, 3)
	//chan01 <- 1
	//chan01 <- 2
	////chan01 <- 3
	//close(chan01)
	//
	//x, ok := <-chan01
	//fmt.Println(x, ok)

	//3. 循环取值
	chan01 := make(chan int, 3)
	chan01 <- 1
	chan01 <- 2
	chan01 <- 3

	close(chan01)
	//遍历管道之前一定要先close管道
	for i := range chan01 {
		fmt.Println(i)
	}

}
