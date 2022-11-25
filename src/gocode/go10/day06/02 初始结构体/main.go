package main

import "fmt"

// 学生数据结构 结构体
type Student struct {
	sid    int
	name   string
	age    int8
	couser []string
}

// 客户关系管理系统数据结构
type cms struct {
	Sid, age            int
	name, gender, email string
}

func main() {

	//声明一个结构体对象,值类型.默认开辟空间,字段赋予零值
	var s Student
	fmt.Println("s: ", s) //s:  {0  0 []}
	//要访问结构体成员,需要使用.号操作符
	fmt.Println("---", s.name, "---") // " "
	//修改成员变量的值
	s.name = "daerwen"
	s.age = 22
	//s.couser[0] = "chinese"
	fmt.Println("s: ", s)

}
