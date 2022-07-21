package main

import "fmt"

//常量const
//定义了常量之后不能修改
//程序运行期间不会改变的量

const pi = 3.1415926

//批量声明常量
const (
	statusOK = 200
	botFound = 404
)

//批量声明常量,如果某一行声明后没有赋值,默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

//iota:
const (
	a1 = iota //0
	a2 = iota //1
	a3
)

const (
	b1 = iota //0   当const出现iota赋值为0
	b2        //1   每增加一行计数+1
	_         //2    匿名变量
	b3        //3
)

//中间插队
const (
	c1 = iota //0
	c2 = 100  //100   iota计数+1
	c3 = iota //2
	c4        //3
)

const (
	d1, d2 = iota + 1, iota + 2 //1   2
	d3, d4 = iota + 1, iota + 2 //2   3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) //1左移10位 1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//pi = 123
	fmt.Println(n1, n2, n3)
	fmt.Println()
	fmt.Println(a1, a2, a3)
	fmt.Println()
	fmt.Println(b1, b2, b3)
	fmt.Println()
	fmt.Println(c1, c2, c3, c4)
	fmt.Println()
	fmt.Println(d1, d2, d3, d4)
}
