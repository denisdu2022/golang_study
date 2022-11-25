package main

import "fmt"

func fib(n int) int {
	if n == 2 || n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	//计算n的阶乘,即n!
	ret := fib(5)
	fmt.Println(ret)
}
