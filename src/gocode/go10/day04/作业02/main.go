package main

import "fmt"

////1. 常量
//const (
//	x = iota  //当iota出现的时候,为0
//	_         //匿名占位 1
//	y         //2
//	z = "zzz" //3
//	k         //声明多个变量时不赋值默认上一个值  zzz 4
//	p = iota  //5
//)

func main() {

	////1.
	//fmt.Println(x, y, z, k, p) //0 2 zzz  zzz 5

	//2.
	slice := []int{10, 11, 12, 13}
	m := make(map[int]*int)
	for key, val := range slice {
		m[key] = &val //把slice的索引当做map的键,取每个slice值的地址给m这个map的值
	}
	fmt.Println(m)
	//map[0:0xc0000160a8 1:0xc0000160a8 2:0xc0000160a8 3:0xc0000160a8]

	for k, v := range m {
		fmt.Println(k, "->", *v) //取值 map[0:10 1:11 2:12 3:13]
	}
	/*输出结果
	0 -> 13
	1 -> 13
	2 -> 13
	3 -> 13*/

}
