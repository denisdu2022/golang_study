package main

import "fmt"

//func foo() {
//	x := 10 //函数内局部变量x = 10
//	fmt.Println(x)
//}

//var x = 30 //全局变量 x = 30

func swap(a int, b int) {
	c := a                //c = 10
	a = b                 //a = 20
	b = c                 // b = 10
	fmt.Println("a: ", a) //20
	fmt.Println("b: ", b) //10
	fmt.Println("c: ", c) //10
}

func main() {
	//var x = 20     //局部变量 x = 20
	//foo()          //调用函数,执行的是函数内部的局部变量10
	//fmt.Println(x) //执行的是main函数内的局部变量20

	////if的局部空间
	//if true {
	//	x := 10
	//	fmt.Println(x)
	//}
	////fmt.Println(x) 会报错,上边的x是if空间中的局部变量
	//
	////for 的局部空间
	//for i := 0; i <= 3; i++ {
	//	fmt.Println(i)
	//}
	////fmt.Println(i) 与上边if的同理

	////案例1
	//var x = 10
	//fmt.Printf("x的地址是%p\n", &x)
	//y := x //把x的值拷贝给了y,值拷贝
	//fmt.Printf("y的地址是%p\n", &y)
	//x = 100
	//fmt.Println(x, y) //100 10

	////案例2
	//var a = []int{1, 2, 3}
	//b := a //引用拷贝也是值拷贝
	//a[0] = 100
	//fmt.Println(a, b) //[100 2 3]  [100 2 3]

	////案例3
	//var m1 = map[string]string{"name": "daerwen"}
	//var m2 = m1 //引用拷贝
	//m2["age"] = "22"
	//fmt.Println(m1) //map[age:22 name:daerwen]
	//fmt.Println(m2) //map[age:22 name:daerwen]

	var x = 10
	var y = 20
	//调用函数
	swap(x, y)
	fmt.Println("x: ", x) //10
	fmt.Println("y: ", y) //20

}
