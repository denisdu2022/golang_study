package main

import "fmt"

func main() {
	//var birth string
	//fmt.Println("请输入您的生日(格式为:1999-3-16): ")
	//fmt.Scan(&birth)
	//nBirth := strings.Split(birth, "-")
	//
	//fmt.Printf("您输入的生日是:%s年-%s月-%s日", nBirth[0], nBirth[1], nBirth[2])

	////& 取地址
	//var a = 10
	//fmt.Println(&a)
	//
	////声明指针变量
	//var p *int
	//p = &a
	//fmt.Println(p)
	//
	////取值操作
	//fmt.Println(&p, reflect.TypeOf(p))
	//fmt.Println(*p, reflect.TypeOf(*p))
	//
	//*p = 200
	//fmt.Println(a)
	//
	////思考
	//var c = 10
	//d := c
	//c = 200
	//fmt.Println(&c, &d)
	//fmt.Println(c, d)

	////声明数组
	//var arr [3]string
	//var arr1 [5]int
	//fmt.Println(arr, arr1) //["" "" ""]  , [0 0 0 0 0 ]
	//arr1[0] = 23
	//arr1[2] = 33
	//fmt.Println(arr1, reflect.TypeOf(arr1)) //[23 0 33 0 0 ]

	//数组的声明并赋值
	//var names = [3]string{"daerwen", "jialiluo", "dafenqi"}
	//fmt.Println(names) //[daerwen  jialiluo dafenqi]
	//var ages = [5]int{33, 66, 99}
	//fmt.Println(ages) //[33 66 99 0 0]

	////省略长度赋值
	//var names1 = [...]string{"daerwen", "jialiluo", "dafenqi"}
	//fmt.Println(names1)  //[daerwen  jialiluo dafenqi]
	//
	////索引赋值
	//var names2 = [...]string{0: "bijialuo", 4: "daerwen"}
	//fmt.Println(names2)
	//fmt.Println(len(names2)) //len(5)

	////range循环
	//var arr2 = [...]int{11, 22, 33, 44, 55}
	//for _, v := range arr2 {
	//	fmt.Println(v)
	//	v = 0
	//	fmt.Println(v)
	//}

	//var a = [...]int{1, 2, 3, 4, 5, 6}
	//a1 := a[0:3]  //3  6
	//a2 := a[0:5]  //5   6
	//a3 := a[1:5]  //4  5
	//a4 := a[1:]   //5    5
	//a5 := a[:]    //6 6
	//a6 := a3[1:2] //1  4
	//fmt.Println("a1的长度%d,容量%d\n", len(a1), cap(a1))
	//fmt.Println("a1的长度%d,容量%d\n", len(a2), cap(a2))
	//fmt.Println("a1的长度%d,容量%d\n", len(a3), cap(a3))
	//fmt.Println("a1的长度%d,容量%d\n", len(a4), cap(a4))
	//fmt.Println("a1的长度%d,容量%d\n", len(a5), cap(a5))
	//fmt.Println("a1的长度%d,容量%d\n", len(a6), cap(a6))

	//var s = make([]int, 5, 10)
	//fmt.Println(len(s), cap(s))
	//fmt.Println(s) //[0,0,0,0,0]
	//s[0] = 100
	//fmt.Println(s) //[100,0,0,0,0]

	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2]                            // [10, 20]
	s2 := s1                                  //  // [10, 20]
	s3 := append(append(append(s1, 1), 2), 3) //[10,20,1,2,3]
	s1[0] = 1000
	fmt.Println(s1)  //[1000,20]
	fmt.Println(s2)  //[1000,20]
	fmt.Println(s3)  //[10,20,1,2,3]
	fmt.Println(arr) //[1000, 20, 1, 2]

}
