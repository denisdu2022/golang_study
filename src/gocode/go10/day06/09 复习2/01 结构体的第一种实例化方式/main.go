package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	sid    int
	name   string
	age    int
	course []string
	nianJi int
}

func main() {

	////声明变量
	//var x int
	//x = 10

	//声明实例化结构体对象
	var s Student                  //s是Student的实例对象
	fmt.Println(s)                 //{0  0 []}
	fmt.Println(reflect.TypeOf(s)) //main.Student

	s.sid = 1001 //通过.的方式对成员变量重新赋值
	s.name = "daer"
	fmt.Println(s.sid, s.name)
	//s.course = "chinese" 这样赋值是错误的
	s.course = make([]string, 3, 5)
	//通过make初始化切片的长度和容量,默认值是[]类型的默认值
	fmt.Println(s.course, len(s.course), cap(s.course), reflect.TypeOf(s.course))
	//[  ] 3 5 []string
	fmt.Printf("%p\n", &s)          //0x1400012a040
	fmt.Printf("%p\n", &(s.sid))    //0x1400012a040 int占8个字节
	fmt.Printf("%p\n", &(s.name))   //0x1400012a048 字符串占16个字节
	fmt.Printf("%p\n", &(s.age))    //0x1400012a058
	fmt.Printf("%p\n", &(s.course)) //0x1400012a060  切片占24个字节,起始地址8个字节,长度8个字节,容量8个字节
	fmt.Printf("%p\n", &(s.nianJi)) //0x1400012a078

}
