package main

import "fmt"

type Student struct {
	sid    int
	name   string
	age    int8
	course []string
}

func NewStudent(sid int, name string, age int8, course []string) *Student {
	return &Student{
		sid:    sid,
		name:   name,
		age:    age,
		course: course,
	}

}

func main() {

	s := NewStudent(1001, "daerwen", 23, nil)
	fmt.Println(s) //&{1001 daerwen 23 []}

}
