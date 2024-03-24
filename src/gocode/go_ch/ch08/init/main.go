package main

import (
	"fmt"
)

type Person struct {
	name string
}

//接収者: 指针接收器
//命名规范,单个单词的首字母小写,多个单词的首字母小写缩写

func (p *Person) SayHello() {

}

//通过指针交换两个值

func swap(a, b *int) {
	//a中存放b中的值,b中的值存放a中的值, 这样看是没有问题的,但是还是值传递
	//a, b = b, a

	//临时值
	t := *a //把a的地址赋值给了临时变量t

	*a = *b //把b的地址赋值给了a ,a就变成了指向b的值

	*b = t

}

func main() {
	////指针的定义
	//var a int = 10
	//b := &a //在变量上取址
	//fmt.Println(b)
	//
	////指针需要初始化
	//var c *int
	//fmt.Println(c) //nil
	//
	////var p *Person //这样会报错:指针初始化会出现nil pointer ,指针需要初始化
	////var ps Person
	//ps := &Person{} //第一种: 初始化方式
	////fmt.Println(p.name)
	//fmt.Println(ps.name)
	//
	//var emptyPerson Person //第二种: 初始化方式 结构体有默认值
	//pi := &emptyPerson
	//
	//fmt.Println(pi)
	//
	////初始化两个关键字: map ,channel ,slice 初始化推荐使用make方法
	////指针初始化推荐使用new函数,否则指针初始化会出现nil pointer
	////map必须初始化
	//
	////var pp *Person = new(Person)
	//var pp = new(Person) //第三种: 初始化方式
	//fmt.Println(pp.name)
	//

	//swap交换
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)

}
