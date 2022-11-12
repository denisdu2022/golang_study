package main

import "fmt"

//声明结构体
type student struct {
	sId    int
	name   string
	gender string
	age    int8
	course []string
}

func main() {

	//声明一个结构体对象

	var info student
	info.sId = 1001
	info.name = "daerwen"
	info.gender = "male"
	info.age = 22
	info.course = []string{"语文", "数学", "英语"}

	fmt.Println(info)      //{1001 daerwen male 22 [语文 数学 英语]}
	fmt.Println(info.name) //daerwen

	info.name = "halibote"
	fmt.Println(info.name)

	//fmt.Println(&info)
	//fmt.Println(&info.sId)
	//fmt.Println(&info.name)
	//fmt.Println(&info.gender)
	//fmt.Println(&info.age)
	//fmt.Println(&info.course)

}
