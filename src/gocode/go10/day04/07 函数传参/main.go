package main

import "fmt"

//func func01(x int) {
//	x = 100
//	fmt.Println(x) //100
//}

//func func02(s []int) {
//	fmt.Printf("func02的s的地址: %p\n", &s)
//	s[0] = 100
//	fmt.Println(s) //[100 2 3]
//
//}

//func func02(s []int) {
//	fmt.Printf("func02的s的地址: %p\n", &s)
//	s = append(s, 1000)
//	fmt.Println(s, len(s), cap(s)) //[1 2 3 1000] 4 6
//
//}

//func func02_01(s1 []int) {
//	s1 = append(s1, 1000)
//	fmt.Println(s1, len(s1), cap(s1)) //[1 2 3 1000] 4 4
//}

func func03(p *int) {
	*p = 100
}

func main() {
	////案例1
	//var x = 10
	//func01(x)      //100
	//fmt.Println(x) //10

	//案例2
	//var s = []int{1, 2, 3}
	//fmt.Printf("main的s的地址:%p\n", &s)
	//func02(s)
	//fmt.Println(s) //[100 2 3]

	////有容量未扩容时是修改的原来的底层数组,没有容量扩容后,生成新的底层数组,容量翻倍
	//var s = []int{1, 2, 3}
	//fmt.Printf("main的s的地址:%p\n", &s)
	//func02(s)
	//fmt.Println(s) //[1 2 3]
	//
	//
	////验证
	//var s1 = make([]int, 3, 4)
	//s1[0], s1[1], s1[2] = 1, 2, 3
	//func02_01(s1)
	//fmt.Println(s1, len(s1), cap(s1)) //[1 2 3] 3 4

	//案例3
	var a = 10
	var p *int = &a
	func03(p)
	fmt.Println(a) //100
}
