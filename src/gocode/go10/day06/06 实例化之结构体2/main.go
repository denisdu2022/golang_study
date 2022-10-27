package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	sid, age     int
	name, gender string
	course       []string
}

func main() {
	s := new(Student)              //&Student{}
	fmt.Println(reflect.TypeOf(s)) //*main.Student
	fmt.Println(s)                 //&{0 0   []}
	s.name = "dafenqi"
	fmt.Println((*s).name) //dafenqi
	fmt.Println(s.name)    //dafenqi

}
