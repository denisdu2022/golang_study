package main

import "fmt"

func main() {
	//基本数据类型属于值类型
	//值类型的特点(整形,浮点型,字符串,布尔型,数组,结构体):当声明未赋值之前存在一个默认值(zero value)
	var x int
	fmt.Println(x)

	var name string
	fmt.Println(name)

	//指针类型属于引用类型,包括切片,map,channel都属于引用类型
	//引用类型当声明未赋值的时候,是没有开辟空间以及没有默认值
	/*var p *int
	*p = 10
	//执行会报错
	*/

	var p *int
	p = new(int)

	*p = 10

	var c *int = new(int)
	*c = 20
	fmt.Println(*p) //10
	fmt.Println(*c) //20

}
