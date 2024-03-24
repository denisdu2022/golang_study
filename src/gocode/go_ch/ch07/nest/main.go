package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Student struct {
	//结构体嵌套
	//第一种嵌套方式
	//p     Person
	//第二种嵌套方式: 匿名嵌套
	Person
	score float32
	//优先级,比内部Person的高(会覆盖)
	name string
}

// 结构体的方法
// (p Person)  方法接收器
// 接收器有两种形态:(p Person) 值传递 ,引用传递 (p *Person)
func (p *Person) print() {
	//有可能该函数中有想要改变p的值,就是person对象很大,数据较大的时候考虑使用指针方式
	//对age重新赋值
	p.age = 19
	fmt.Printf("name: %s, age: %d\r\n", p.name, p.age)
}

func (p Person) print2() {
	fmt.Printf("name: %s, age: %d\r\n", p.name, p.age)
}

//空结构体不占用内存空间

type MySet map[string]struct {
}

func main() {

	p := &Person{
		"libai", 18,
	}

	//p := Person{
	//	"libai", 18,
	//}
	//内置函数print
	//print(p.age)

	////调用方式
	//p.print()
	//fmt.Println(p.age)

	p.print2()

	//s := Student{
	//	Person{
	//		"libai", 18,
	//	},
	//	95.6,
	//	"lili",
	//}
	////嵌套的结构体也继承了方法
	//s.print()

	//s  := Student{}
	//s.p.name = "libai"
	//fmt.Println(s.p.age)

	//s := Student{
	//	Person{
	//		"libai", 18,
	//	},
	//	95.6,
	//	"lili",
	//}
	////s.name = "lili"
	//
	//fmt.Println(s.name)
	//fmt.Println(s)

}
