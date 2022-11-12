package main

import "fmt"

func foo() func() {
	inner := func() {
		fmt.Println("函数inner执行......")
	}
	return inner

}

func main() {

	foo()()

}
