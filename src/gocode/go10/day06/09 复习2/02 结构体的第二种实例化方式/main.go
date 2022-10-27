package main

import "fmt"

type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

func main() {

	//声明并赋值
	//键值对赋值不用按照顺序
	var s = Student{name: "daerwen", age: 23, sid: 1002}
	//{1002 daerwen 23 []}
	fmt.Println(s)
	//以下是必须要按照顺序赋值
	//var s1 = Student{1003, "dafenqi", 22, } 这样是错的,所有属性必须都要按照顺序赋值
	var s2 = Student{1003, "dafenqi", 22, nil}
	var s3 = Student{1004, "daerwen", 18, []string{"chines", "math"}}
	//{1003 dafenqi 22 []}
	fmt.Println(s2)
	//{1004 daerwen 18 [chines math]}
	fmt.Println(s3)

}
