package main

func main() {
	////案例1 :
	//foo := func() {
	//	fmt.Println("这是执行foo1功能函数")
	//}
	//defer foo()
	//foo = func() {
	//	fmt.Println("这是执行foo2功能函数")
	//}
	////输出结果
	////这是执行foo1功能函数
	//
	////案例2 :
	//x := 10
	//defer func(a int) {
	//	fmt.Println(a) //10
	//}(x) //把x这个实参传给了匿名函数的形参a,所以a = x = 10
	//x++
	//
	////案例3 :
	//x := 10
	//defer func() {
	//	fmt.Println(x) //在匿名函数内部找X,并未找到,去全局找的时候x已经自加1重新赋值了,所以 x = 11
	//	//保留X的地址
	//}()
	//x++
}
