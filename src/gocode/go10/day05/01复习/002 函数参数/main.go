package main

import (
	"fmt"
	"reflect"
)

//计算I到N的和的函数
//声明这个计算函数
//x,y 形式参数
//func sumItoN(x, y int) {
//	var ret = 0
//	for i := x; i <= y; i++ {
//		ret += i
//	}
//	fmt.Println(ret)
//}

//不限长参数
//不限位数加法器函数
func add2(num ...int) {
	fmt.Println(num, reflect.TypeOf(num)) //参数类型是一个切片[]int
	var ret = 0
	for _, v := range num {
		ret += v
	}
	fmt.Println(ret)

}

func main() {

	////调用I到N的和的函数
	////x,y实际参数
	////变量赋值,把第一个实际参数赋值给第一个形式参数,以此类推
	////位置参数,按照位置以此传参
	//sumItoN(56, 1020) //519170
	//sumItoN(2, 200)

	//不限位数加法器函数调用
	add2(1, 2, 3, 4)
	add2(13, 15, 15)

}
