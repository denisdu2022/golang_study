package main

import "fmt"

//函数拥有开辟作用域的能力

func foo() {
	var x = 20
	fmt.Println(x)
}

func bar() {
	x = 1000
	fmt.Println(x)
	fmt.Println(&x) //0x292338

}

//全局变量
var x = 1

func main() {
	foo()           //20
	bar()           //1000
	fmt.Println(x)  //1000
	fmt.Println(&x) //0x292338

}
