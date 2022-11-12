package main

import "fmt"

//声明结构体
type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

func initSid(s Student) {
	s.sid = 0
}

func initSid1(p *Student) {
	//(*p).sid = 0
	//可以直接这样写  语法糖
	p.sid = 0
}

func main() {
	//声明并赋值
	var s1 = Student{sid: 1001, name: "dafenqi", age: 18}
	fmt.Println(s1)

	var s2 = s1 //值拷贝
	s2.age = 20
	fmt.Println(s1.age) //18

	//初始化学号
	initSid(s1)
	fmt.Println(s1.sid) //1001 没有改变

	initSid(s1)
	fmt.Println()

	//初始化学号为0
	initSid1(&s1)
	fmt.Println(s1.sid)

	//直接取址
	var s3 = &Student{sid: 1003, name: "bijiasuo", age: 17}
	fmt.Println(s3) //&{1003 bijiasuo 17 []}
	initSid1(s3)
	fmt.Println(s3.sid) //0

}
