package main

import (
	"fmt"
	"reflect"
)

func main() {

	//var x = 10
	////&变量  :获取变量地址
	//fmt.Println("赋值之前X对应地址: ",&x) //0x1400012c008
	//x = 100
	//fmt.Println("赋值之后X对应地址: ",&x) //0x1400012c008

	var x = 10 //x是整形变量
	var p *int //p是一个整形的指针类型
	p = &x
	fmt.Println(p) //0x1400012c008
	*p = 100
	fmt.Println(*p, reflect.TypeOf(*p))
	fmt.Println(x)

}
