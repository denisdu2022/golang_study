package main

import "fmt"

/*
	函数功能:代码的一种组织形式,实现模块化,避免重复

	声明函数
	func 函数名(){
		函数体  功能代码
	}

	函数的调用
	func()
*/

func printLing() {
	// 层数
	var n int = 6

	for i := 1; i <= n; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := n - 1; i >= 1; i-- {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func main() {

	printLing()

}
