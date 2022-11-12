package main

import "fmt"

type Student struct {
	sid, age     int
	name, gender string
	course       []string
}

func main() {
	//方式1:先声明在赋值
	s1 := Student{}
	s1.sid = 1001
	s1.name = "daer"
	fmt.Println(s1) //{1001 0 daer  []}

	//方式2:键值对赋值
	s2 := Student{sid: 1001, name: "dafen", gender: "male", course: []string{"chinese", "math", "english"}}
	//{1001 0 dafen male [chinese math english]}
	fmt.Println(s2)

	//方式3: 多值赋值
	s3 := Student{1003, 22, "daerwen", "male", []string{"chinese", "math", "english"}}
	//{1003 22 daerwen male [chinese math english]}
	fmt.Println(s3)

}
