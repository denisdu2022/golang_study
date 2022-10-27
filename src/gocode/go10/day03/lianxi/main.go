package main

import "fmt"

func main() {

	//count := 0       //初始化语句  count: 0
	//for count < 10 { //条件判断
	//	fmt.Println(count)
	//	count++ //步进语句
	//}

	////打印1-100
	//sum := 0
	//for sum <= 100 {
	//	fmt.Println(sum)
	//	sum++
	//}

	////打印100-1
	//sum := 100
	//for sum > 0 {
	//	sum--
	//	fmt.Println(sum)
	//
	//}

	//var sum = 1
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//	fmt.Println(sum)
	//}

	//for i := 0; i < 10; i++ {
	//	if i == 8 {
	//		break
	//	}
	//	fmt.Println("i 是 : ", i)
	//}
	//fmt.Println("到8跳出循环 : end")

	//for i := 0; i < 5; i++ {
	//	if i == 3 {
	//		continue
	//
	//	}
	//	fmt.Println("i 是 : ", i)
	//}

	//for i := 0; i < 5; i++ {
	//	for j := 0; j < 5; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Print("\n")
	//}
	//for i := 0; i < 5; i++ {
	//	for j := 0; j <= i; j++ {
	//		fmt.Printf("*")
	//	}
	//	fmt.Println()
	//}

	var sum = 0
	for s1 := 1; s1 <= 99; s1++ {
		if s1 == 88 {
			continue
		}
		if s1%2 == 0 {
			sum -= s1
		} else {
			sum += s1
		}

	}
	fmt.Println(sum)

}
