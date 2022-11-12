package main

import (
	"fmt"
	"reflect"
)

//声明结构体
type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

func initSid(p *Student) {
	p.sid = 1001
}

func main() {
	var s1 = new(Student) //&Student{} 先开辟空间,在取内存地址
	//new()开辟一块空间,把地址返回
	fmt.Println(s1) //&{0  0 []}
	s1.name = "bijiasuo"
	fmt.Println(s1.name)            //== (*s1).name 语法糖  "bijiasuo"
	fmt.Println(reflect.TypeOf(s1)) //*main.Student

	initSid(s1)
	fmt.Println(s1.sid) //1001

	var s = make([]int, 3)
	//make() 初始化一个类型,把类型本身返回
	fmt.Println(reflect.TypeOf(s)) //[]int

	s1.name = "daerwen"
	fmt.Println(s1.name)

}
