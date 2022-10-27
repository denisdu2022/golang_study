package main

import "fmt"

//声明结构体
type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

//模拟构造函数
func NewStudent(sid int, name string, age int, course []string) Student {
	return Student{
		sid:    sid,
		name:   name,
		age:    age,
		course: course,
	}

}

//Student类型的方法接收器
func (s Student) read() {
	fmt.Println("学生正在读书")
}

//Student类型的方法接收器带参数
func (s Student) read1(book string) {
	fmt.Printf("%s正在读%s\n", s.name, book)
}

func (s Student) learn() {
	//fmt.Println("学生正在学习")
	fmt.Printf("%s正在学习\n", s.name)
}

func main() {
	s1 := NewStudent(1001, "daerwen", 22, []string{"chines", "math"})
	fmt.Println(s1.name)
	fmt.Println(s1.age)
	//调用Student里边的方法
	//s1.learn() //学生正在学习  s = s1
	s1.learn() //daerwen正在学习

	s2 := NewStudent(1002, "bijiasuo", 18, []string{"chines", "math"})
	s2.learn()

	//传参数调用方法接收器
	s2.read1("本草纲目") //bijiasuo正在读本草纲目

}
