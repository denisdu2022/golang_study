package main

import "fmt"

func main() {

	//切片面试题1
	s1 := []int{1, 2, 3} //[1 2 3] []int 3 3
	s2 := s1[1:]         //[2 3]  2 2
	s2[1] = 4            //[2 4]
	fmt.Println(s1)      //[1 2 4]

	//切片面试题2
	var a = []int{1, 2, 3}
	b := a //引用类型接收引用类型 (引用拷贝)
	c := a[1:]
	a[0] = 100
	fmt.Println(b)                 //[100 2 3]
	fmt.Println(c, len(c), cap(c)) //[2 3]    2 2

}
