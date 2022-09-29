package main

import (
	"fmt"
	"reflect"
)

func main() {
	////byte类型
	//var b1 byte
	//b1 = 'A' //必须是单引号
	////b1 = 98
	//fmt.Println(b1, reflect.TypeOf(b1)) //65 uint8  bete就是unit8类型  type byte = uint8
	//fmt.Println("%c\n", b1)             //ascll数字
	//fmt.Println("%d\n", b1)             //65
	//
	////uint8类型
	//var b2 uint8
	//b2 = 65
	//fmt.Printf("%c\n", b2) //A
	//fmt.Printf("%d\n", b2) //ascll数字
	//fmt.Println(b2)        //65
	//
	////var b3 byte
	//var b3 rune
	//b3 = 'A'
	////rune,占用4个字节byte,共32比特位bit,所以它和int32本质上是没有区别,他表示一个Unicode字符
	//fmt.Println(b3, reflect.TypeOf(b3)) //65 int32
	//fmt.Println(string(b3))             //'A'

	//y := 'a'
	//fmt.Println(y, reflect.TypeOf(y)) //97  int32

	var y byte                        //byte 字节类型 声明字符
	y = 'a'                           //""双引号表示字符串,单引号表示字符
	fmt.Println(y, reflect.TypeOf(y)) //97 uint8  type byte = uint8
	fmt.Printf("%c\n", y)             //格式化输出 字符 'a'
	fmt.Printf("%d\n", y)             //十进制,还是ASCII 数字 97
	fmt.Printf("%b\n", y)             //97的二进制 1100001
	fmt.Printf("%x\n", y)             //97的十六进制,逢16进1 61

	//uint8类型
	var b2 uint8
	b2 = 65
	fmt.Println(b2, reflect.TypeOf(b2)) //65 uint8
	fmt.Printf("%c\n", b2)              //字符 'A'
	fmt.Printf("%d\n", b2)              //十进制  ASCII 数字 65

	//type rune = int32
	//rune类型
	z := '哈'
	fmt.Println(z, reflect.TypeOf(z)) //21704   int32
	fmt.Printf("%c\n", z)             //'哈'

}
