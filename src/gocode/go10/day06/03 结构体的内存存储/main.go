package main

import "fmt"

type Student struct {
	sid, age int
	name     string
	course   []string
}

func main() {

	//声明一个结构体对象,值类型,默认开辟空间,字段赋予零值
	var s Student
	fmt.Println("s: ", s)
	s.sid = 1001
	s.name = "daerwen"
	s.age = 23
	s.course = []string{"chinese", "math", "english"}
	fmt.Println(s.course)
	fmt.Printf("%p\n", &s)          //0x14000022080
	fmt.Printf("%p\n", &(s.sid))    //0x14000022080
	fmt.Printf("%p\n", &(s.name))   //0x14000022090
	fmt.Printf("%p\n", &(s.age))    //0x14000022088
	fmt.Printf("%p\n", &(s.course)) //0x140000220a0

}
