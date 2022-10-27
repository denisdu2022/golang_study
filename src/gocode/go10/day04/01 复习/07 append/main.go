package main

import "fmt"

func main() {

	//var s []int                       //声明一个s的切片
	//s1 := append(s, 1)                //对s这个切片扩容追加1
	//fmt.Println(s1, len(s1), cap(s1)) //[1] 1 1
	//
	//s2 := append(s1, 2, 3, 4)
	//fmt.Println(s2, len(s2), cap(s2)) // [1 2 3 4] 4 4
	//
	//var t = []int{5, 6, 7}
	//s3 := append(s2, t...)
	//fmt.Println(s3, len(s3), cap(s3), reflect.TypeOf(s3)) //[1 2 3 4 5 6 7] 7 8 []int 容量不够扩充1倍
	//
	//var s4 = make([]int, 3, 10)
	//s5 := append(s4, 100)
	//fmt.Println(s5, len(s5), cap(s5)) //[0 0 0 100] 4  10

	////append原理
	//a := []int{11, 22, 33}
	//fmt.Println(len(a), cap(a)) //3 3
	//
	//c := append(a, 44)             //原有数组容量不够,append扩容一倍后,将原有数组的值复制一份,并且将44追加到一个新的数组
	//a[0] = 100                     //改变的是原有数组的值
	//fmt.Println(a, len(a), cap(a)) //[100 22 33] 3 3
	//fmt.Println(c, len(c), cap(c)) //[11 22 33 44] 4 6

	////make
	//a := make([]int, 3, 10)
	//fmt.Println(a) //[0 0 0] 3 10
	//b := append(a, 11, 22)
	//fmt.Println(a, len(a), cap(a)) //小心a等于多少 [0 0 0] 3 10
	//fmt.Println(b, len(b), cap(b)) //有容量不扩容,还是对原数组的引用 [0 0 0 11 22] 5 10
	//a[0] = 100
	//fmt.Println(a, len(a), cap(a)) //[100  0 0] 3 10
	//fmt.Println(b, len(b), cap(b)) //[100 0 0 11 22] 5 10
	///*
	//	输出结果:
	//	[0 0 0]
	//	[0 0 0] 3 10
	//	[0 0 0 11 22] 5 10
	//	[100 0 0] 3 10
	//	[100 0 0 11 22] 5 10
	//*/

	//append 扩容面试题
	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2]
	fmt.Println(s1, len(s1), cap(s1)) //[10 20] 2 4

	s2 := s1
	fmt.Println(s2, len(s2), cap(s2)) //[10 20] 2 4
	s3 := append(append(append(s1, 1), 2), 3)
	fmt.Println(s3, len(s3), cap(s3)) //[10 20 1 2] 追加到2后容量不足,扩容1倍[10 20 1 2 3] 5 8

	s1[0] = 1000
	fmt.Println(s1)  //[1000 20]
	fmt.Println(s2)  //[1000 20]
	fmt.Println(s3)  //[10 20 1 2 3] 5 8
	fmt.Println(arr) //[1000 20 1 2] 4 4

}
