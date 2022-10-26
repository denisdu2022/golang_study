package main

import "fmt"

func main() {

	//1.
	var shu1 int
	var shu2 int
	fmt.Println("请输入第一个数: ")
	fmt.Scan(&shu1)
	fmt.Println("请输入第二个数: ")
	fmt.Scan(&shu2)
	//判断
	if shu1 > shu2 {
		fmt.Printf("最大数是: %d\n", shu1)
	} else {
		fmt.Printf("最大数是: %d\n", shu2)
	}

	//2.
	var shuZi int
	fmt.Println("请输入一个数字: ")
	fmt.Scan(&shuZi)
	//判断是奇数还是偶数
	if shuZi%2 == 0 {
		fmt.Println("您输入的是偶数...")
	} else {
		fmt.Println("您输入的是奇数...")
	}
}
