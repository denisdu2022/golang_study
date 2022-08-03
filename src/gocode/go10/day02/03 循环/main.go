package main

import "fmt"

func main() {

	//打印1-100中所有的偶数

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
