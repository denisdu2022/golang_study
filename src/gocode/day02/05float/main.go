package main

import (
	"fmt"
)

func main() {
	fmt.Println()

	//浮点数
	//math.MaxFloat32 //32位浮点数
	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认go语言中的小数都是float64类型
	f2 := float32(1.23456) //声明一个float32类型
	fmt.Printf("%T\n", f2)
	//f1 = f2  float32类型的值不能直接赋值给float64








	
}
