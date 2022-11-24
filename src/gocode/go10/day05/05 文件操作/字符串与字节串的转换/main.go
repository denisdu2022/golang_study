package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

// 反转函数
func reverseString(s []byte) []byte {
	var i, j = 0, len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		j++
		i--
	}
	return s
}

func main() {

	//1 字符串类型(string) 转为字节串类型([]byte)
	var s = "哈哈"
	fmt.Println(s, reflect.TypeOf(s)) //哈哈 string

	var b = []byte(s) //默认使用utf-8编码
	fmt.Println(b, reflect.TypeOf(b))
	//[229 147 136 229 147 136] []uint8

	//可以通过代码len([]rune(s)) 来获得字符串中字符的数量,但使用utf8.RuneCountInString(s) 效率会更高一些
	s1 := "hello,世界"
	r1 := []byte(s1)
	r2 := []rune(s1)
	fmt.Println(r1, reflect.TypeOf(r1))
	//[104 101 108 108 111 44 228 184 150 231 149 140] []uint8
	fmt.Println(r2, reflect.TypeOf(r2))
	//[104 101 108 108 111 44 19990 30028] []int32
	//统计字符个数
	fmt.Println(len(r1)) //12
	//统计字符个数
	fmt.Println(len(r2)) //8

	fmt.Println(utf8.RuneCountInString(s1)) //8

	//2 byte 转string
	fmt.Println(string(b))
	var data = []byte{229, 147, 136, 229, 147, 136}
	fmt.Println(data, string(data))
	//[229 147 136 229 147 136]   哈 哈

	//练习1
	s2 := "hello"
	c1 := []byte(s2)
	c1[0] = 'c'
	fmt.Println(c1, string(c1))
	//[99 101 108 108 111] cello

	//练习2 反转
	//调用反转函数
	//reverseString([]byte(s2))

}
