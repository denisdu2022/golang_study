package main

//iota是go语言的计数器,只能在常量的表达式中使用
//类似枚举
// const (
// 	a1 = iota
// 	a2
// 	a3
// 	a4
// )

// const (
// 	a1 = iota //0
// 	a2   //1
// 	_   //2
// 	a3  //3
// )

//中间插队
// const (
// 	a1 = iota
// 	a2 = 100
// 	a3 = iota
// 	a4
// )

//多个常量声明在一行
// const (
// 	d1, d2 = iota + 1, iota + 2

// 	//每新增一行常量声明+1 ,上一行的空行不算
// 	d3, d4 = iota + 1, iota + 2
// )

func main() {
	// 	// fmt.Println(a1, a2, a3, a4)
	// 	fmt.Println("d1:", d1)
	// 	fmt.Println("d2:", d2)
	// 	fmt.Println("d3:", d3)
	// 	fmt.Println("d4:", d4)
}

//定义数量级
