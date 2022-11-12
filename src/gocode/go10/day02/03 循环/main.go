package main

import "fmt"

func main() {

	////for循环
	//
	////01 打印1到100
	//
	//var cont = 1
	//for cont < 100 {
	//	fmt.Println(cont)
	//	cont++
	//}

	//var cont = 1
	//for cont < 100 {
	//	cont++
	//	fmt.Println(cont)
	//
	//}
	//
	//for i := 1; i <= 100; i++ {
	//	fmt.Println(i)
	//}

	////打印100-1
	//var i = 100
	//for i > 0 {
	//	i--
	//	fmt.Println(i)
	//}

	////打印1-100中所有的偶数
	//
	//for i := 1; i <= 100; i++ {
	//	if i%2 == 0 {
	//		fmt.Println(i)
	//	}
	//}

	//计算1+到100的和
	var s = 1
	for i := 1; i <= 100; i++ {
		s += i
		fmt.Println(s)
	}

}
