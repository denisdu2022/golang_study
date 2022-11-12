package main

import "fmt"

//计算n的阶乘功能函数
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {

	//计算n的阶乘,及n!
	var ret = factorial(4)
	fmt.Println(ret) //24

}
