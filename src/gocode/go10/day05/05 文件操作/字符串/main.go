package main

import "fmt"

func main() {
	//s1 := "hello"
	//fmt.Println(s1, reflect.TypeOf(s1)) //hello string
	//
	//s2 := s1[:]
	//s3 := s1[1:]
	//fmt.Println(&s1, (*reflect.StringHeader)(unsafe.Pointer(&s1))) //0x14000104210 &{4304750566 5}
	//fmt.Println(&s2, (*reflect.StringHeader)(unsafe.Pointer(&s2))) //0x14000104240 &{4304750566 5}
	//fmt.Println(&s3, (*reflect.StringHeader)(unsafe.Pointer(&s3))) //0x14000104250 &{4304750567 4}

	////简单使用
	var s = "哈abc"
	fmt.Println(s, len(s)) //哈abc 6

	//字符串的使用
	//本质上Unicode是一个编码集,和ascii码相同,而utf8是编码规则
	//var a = '哈'
	//fmt.Printf("字符'哈'Unicode的十进制: %d\n", a)
	//fmt.Printf("字符'哈'Unicode的十六进制: %x\n", a)
	//fmt.Printf("字符'哈'Unicode的二进制: %b\n", a)
	///*	字符'哈'Unicode的十进制: 21704
	//	字符'哈'Unicode的十六进制: 54c8
	//	字符'哈'Unicode的二进制: 101010011001000*/
	//
	//var b = 101010011001000
	//fmt.Printf("字符'哈'的utf8: %x\n", b)
	//
	//var c = "哈abc"
	////10100000000000000000100010000001000010000
	//fmt.Printf("%b\n", &c)
	//
	//for i := 0; i < len(c); i++ {
	//	fmt.Printf("%d\n", c[i]) //存储字节的10进制数
	//}
	///*	229
	//	147
	//	136
	//	97
	//	98
	//	99*/
	//
	//for _, v := range c {
	//	//通过存储的utf8解析到Unicode值和对应的符号
	//	fmt.Printf("%d,%c\n", v, v)
	//}
	///*	21704,哈
	//	97,a
	//	98,b
	//	99,c*/

	for i, v := range s {
		fmt.Println(i, v, string(v))
	}

}
