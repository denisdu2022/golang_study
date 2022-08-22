package main

import "fmt"

func printLingxing() {
	var n int = 6
	for i := 1; i <= n; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := n - 1; i >= 1; i-- {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	////打印两个六层菱形
	//var n int = 6
	//for i := 1; i <= n; i++ {
	//	for k := 1; k <= n-i; k++ {
	//		fmt.Print(" ")
	//	}
	//	for j := 1; j <= 2*i-1; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
	//for i := n - 1; i >= 1; i-- {
	//	for k := 1; k <= n-i; k++ {
	//		fmt.Print(" ")
	//	}
	//
	//	for j := 1; j <= 2*i-1; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
	//
	//fmt.Println("**********************")
	//fmt.Println("功能执行开始....")
	//fmt.Println("              ")
	//fmt.Println("功能执行结束....")
	//fmt.Println("**********************")
	//
	////再次打印菱形
	//var n1 int = 6
	//for i := 1; i <= n1; i++ {
	//	for k := 1; k <= n1-i; k++ {
	//		fmt.Print(" ")
	//	}
	//	for j := 1; j <= 2*i-1; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
	//for i := n1 - 1; i >= 1; i-- {
	//	for k := 1; k <= n1-i; k++ {
	//		fmt.Print(" ")
	//	}
	//
	//	for j := 1; j <= 2*i-1; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	//上边的代码使用函数实现
	printLingxing()
	fmt.Println("**********************")
	fmt.Println("功能执行开始....")
	fmt.Println("              ")
	fmt.Println("功能执行结束....")
	fmt.Println("**********************")
	printLingxing()

}
