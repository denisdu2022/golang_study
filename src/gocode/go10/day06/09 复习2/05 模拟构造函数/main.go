package main

import "fmt"

// 声明结构体
type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

func NewStudent(sid int, name string, age int, course []string) Student {
	return Student{
		sid:    sid,
		name:   name,
		age:    age,
		course: course,
	}
}

func main() {

	//s1 := Student{
	//	sid:    1001,
	//	name:   "daerwen",
	//	age:    22,
	//	course: []string{"chinese", "math"},
	//}
	////{1001 daerwen 22 [chinese math]}
	//fmt.Println(s1)

	s1 := NewStudent(1001, "daeren", 23, []string{"chinese", "math"})
	fmt.Println(s1) //{1001 daeren 23 [chinese math]}

}
