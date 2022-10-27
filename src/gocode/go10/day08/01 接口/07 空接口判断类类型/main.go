package main

import (
	"fmt"
	"reflect"
)

//Person 接口
type Person interface {
	eat() //吃的方法
}

//学生类
type Stu struct {
	name string
	age  int
}

//实现吃的方法
func (s Stu) eat() {
	fmt.Printf("%s正在吃饭,今年%d岁了...\n", s.name, s.age)
}

//Dog接口
type Dog interface {
	spark()
}

func main() {

	var x interface{}
	x = Stu{"dafenqi", 22}
	v, ok := x.(Stu)
	fmt.Println(v, ok) //{dafenqi 22} true
	v1, ok1 := x.(Person)
	fmt.Println(v1, ok1) //{dafenqi 22} true

	v2, ok2 := x.(Dog)
	fmt.Println(v2, ok2) //<nil> false
	//x本身的类型
	fmt.Println(reflect.TypeOf(x)) //main.Stu

}
