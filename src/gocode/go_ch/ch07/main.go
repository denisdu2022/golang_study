package main

import (
	"fmt"
	"strconv"
)

type MyInt int //自定义类型   基于已有的类型自定义一个类型

func (mi MyInt) string() string {
	return strconv.Itoa(int(mi))
}

func main() {
	//type 关键字
	/*
		1.定义结构体
		2.定义接口
		3.定义类型别名
		4.类型定义
		5.类型判断
	*/
	////别名实际上是为了更好的理解代码
	//type MyInt = int //类型别名
	//var i MyInt = 12
	//var j int = 8
	//fmt.Printf("%T\r\n", i)        //int
	//fmt.Println(reflect.TypeOf(i)) //int
	//fmt.Println(i + j) //在编译的时候类型别名会被直接替换为int

	var i MyInt = 12
	fmt.Println(i.string())
	var j int = 8           //即希望是int类型,又希望可以增加方法
	fmt.Printf("%T\r\n", i) //main.MyInt
	fmt.Println(int(i) + j)

	//类型判断
	var a interface{} = "abc"
	switch a.(type) {
	case string:
		fmt.Println("字符串")
	}

}
