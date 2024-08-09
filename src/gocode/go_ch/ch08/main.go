package main

import "fmt"

type Person struct {
	name string
}

func changeName(p Person) {
	p.name = "lili"
}

// * 指针类型 &取址符号
func changeName1(p *Person) {
	p.name = "lili"
}

func main() {
	//指针 : 希望结构体传值的时候,在函数中修改的值可以反映到变量中

	//p := Person{
	//	name: "libai",
	//}

	//changeName(p)   //值并不会被改变
	//changeName1(&p) //值会改变
	//fmt.Println(p.name)

	//var pi *Person = &p
	//fmt.Printf("%p\r\n", pi) //地址
	//fmt.Println(pi)

	//指针的定义
	//var po *Person
	//赋值
	//po = &p
	//fmt.Println(po)

	//po := Person{ //值对象
	//	"libai",
	//}
	//fmt.Println(po)

	po := &Person{ //指针对象
		"libai",
	}
	fmt.Println(po)

	//和其他的语言对比
	//1.可以和其他语言一样直接使用
	po.name = "libai4"
	//2.go语言限制了指针的运算,在C语言中你拿到一个指针之后可以进行+1
	//go的指针是一个阉割版(常规意义),而unsafe包可以使用,一般情况下不使用(常规),但是用unsafe包时会提示不安全

	//(*po).name = "libai3"
	//fmt.Println(po.name)
	po.name = "libai4"
	fmt.Println(po.name)
}
