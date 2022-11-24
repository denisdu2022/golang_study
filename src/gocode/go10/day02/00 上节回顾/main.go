package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	//03.Parse系列函数
	//ParseInt
	/*输入:
	1.数字的字符串形式
	2.base,数字字符串的进制,比如2进制,10进制.
	3.bitsize 的含义是大小限制,如果字符串转化的整形数据类型超过了bitsize的最大值,那么输出的int64
	为bitSize的最大值,err就会显示数据超出范围
	*/

	//i1, err := strconv.ParseInt("1000", 10, 8)
	//fmt.Println(i1, err)
	//fmt.Println(reflect.TypeOf(i1), reflect.TypeOf(err))
	//
	//i2, _ := strconv.ParseInt("1000", 10, 64)
	//fmt.Println(i2, reflect.TypeOf(i2))

	////指定了bitSize也是float64 <nil> 类型是空对象
	//f2, err := strconv.ParseFloat("3.1415926", 32)
	//fmt.Println(f2, err)
	//fmt.Println(reflect.TypeOf(f2), reflect.TypeOf(err))
	//
	//f3, _ := strconv.ParseFloat("3.1415926", 64)
	//fmt.Println(f3, reflect.TypeOf(f3))

	//返回字符串表示的bool值.它接受1,0,t,f,T,F,true,false,True,False,TRUE,FALSE.否则返回错误
	b1, _ := strconv.ParseBool("true")
	fmt.Println(b1, reflect.TypeOf(b1))
	b2, _ := strconv.ParseBool("false")
	fmt.Println(b2, reflect.TypeOf(b2))
	b3, _ := strconv.ParseBool("1")
	fmt.Println(b3, reflect.TypeOf(b3))
	b4, _ := strconv.ParseBool("0")
	fmt.Println(b4, reflect.TypeOf(b4))
	b5, _ := strconv.ParseBool("t")
	fmt.Println(b5, reflect.TypeOf(b5))
	b6, _ := strconv.ParseBool("f")
	fmt.Println(b6, reflect.TypeOf(b6))

}
