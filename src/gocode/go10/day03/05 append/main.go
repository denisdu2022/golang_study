package main

import "fmt"

func main() {

	////扩容机制
	////示例1:
	//a := []int{12, 23, 34}
	//fmt.Println(len(a), cap(a)) //3 3
	//
	//c := append(a, 45)
	//a[0] = 100
	//fmt.Println(a, len(a), cap(a)) //[100 23 34]  3 3
	//fmt.Println(c, len(c), cap(c)) //[12 23 34 45] 4 6
	//fmt.Printf("a:的地址是%p\n", &a)
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Printf("b:的地址是%p\n", &c)
	//fmt.Println(reflect.TypeOf(c))

	////示例2:
	//a := make([]int, 3, 10)
	//fmt.Println(a)              //[0 0 0]
	//fmt.Println(len(a), cap(a)) //3 10
	//
	//b := append(a, 45, 36)
	//fmt.Println(b) //[0 0 0 45 36]
	//a[0] = 100
	//fmt.Println(a)              //[100 0 0]
	//fmt.Println(len(a), cap(a)) //3 10
	//fmt.Println(b)              //[100 0 0 45 35]
	//fmt.Println(len(b), cap(b)) //5 10
	//
	////示例3:
	//l := make([]int, 5, 10)
	//v1 := append(l, 1)
	//fmt.Printf("%p\n", &l)
	//fmt.Println(l, len(l), cap(l))    // [0 0 0 0 0] 5 10
	//fmt.Println(v1, len(v1), cap(v1)) //[0 0 0 0 0 1]  6 10
	//fmt.Printf("%p\n", &v1)
	//
	//v2 := append(l, 2)
	//fmt.Println(v2, len(v2), cap(v2)) //[0 0 0 0 0 2]  6 10
	//fmt.Printf("%p\n", &v2)
	//fmt.Println(l)  //[0 0 0 0 0]  5 10
	//fmt.Println(v1) //[0 0 0 0 0 2]  6 10  why?
	//
	////验证:是否以最后追加的为准
	//v3 := append(l, 10)
	//fmt.Println(v3) //[0 0 0 0 0 10] 6 10
	//fmt.Println(v2)
	//fmt.Println(v1)
	//fmt.Println(l)

	//面试题
	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2]                            //[10 20]  2 4
	s2 := s1                                  //[10 20]
	s3 := append(append(append(s1, 1), 2), 3) //[10 20 1 2 3]  //5 8
	s1[0] = 1000                              //[1000 20]
	fmt.Println(s2[0], s2)                    //[1000 20]
	fmt.Println(s3[0], s3)                    //[10 20 1 2 3]

}
