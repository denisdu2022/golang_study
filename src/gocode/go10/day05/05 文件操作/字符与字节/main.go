package main

import (
	"fmt"
	"reflect"
)

func main() {
	//byte类型
	var b1 byte
	b1 = 'A' //必须是单引号
	//b1 = 98
	fmt.Println(b1, reflect.TypeOf(b1)) //65 uint8
	fmt.Println("%c\n", b1)             //ascll数字
	fmt.Println("%d\n", b1)             //65

	//uint8类型
	var b2 uint8
	b2 = 65
	fmt.Printf("%c\n", b2) //A
	fmt.Printf("%d\n", b2) //ascll数字
	fmt.Println(b2)        //65

	//var b3 byte
	var b3 rune
	b3 = 'A'
	fmt.Println(b3, reflect.TypeOf(b3)) //65 int32

}
