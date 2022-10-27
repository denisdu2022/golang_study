package main

//func sum1to100() {
//	var sum = 0
//	for i := 1; i <= 100; i++ {
//		sum += i
//	}
//	fmt.Println(sum)
//}

//打印菱形,带参数
//func printLingxing(n int) {
//	for i := 1; i <= n; i++ {
//		for k := 1; k <= n-i; k++ {
//			fmt.Print(" ")
//		}
//		for j := 1; j <= 2*i-1; j++ {
//			fmt.Print("*")
//		}
//		fmt.Println()
//	}
//	for i := n - 1; i >= 1; i-- {
//		for k := 1; k <= n-i; k++ {
//			fmt.Print(" ")
//		}
//
//		for j := 1; j <= 2*i-1; j++ {
//			fmt.Print("*")
//		}
//		fmt.Println()
//	}
//}

////计算1到N的和,带参数函数
//func sum1toN(n int) {
//	sum := 0
//	for i := 1; i <= n; i++ {
//		sum += i
//	}
//	fmt.Println(sum)
//
//}

////函数声明 两个形参:x,y
//func addSum(x int, y int) {
//	fmt.Println(x + y)
//}

////不定长函数参数
////变参函数
//func sum(nums ...int) {
//	fmt.Println("nums", nums)
//	total := 0
//	for _, num := range nums {
//		total += num
//	}
//	fmt.Println(total)
//}

//func sum(base int, nums ...int) int {
//	fmt.Println(base, nums)
//	sum := base
//	for _, v := range nums {
//		sum = sum + v
//	}
//	return sum
//}

////加法计算函数
//func add_cal(x, y int) int {
//	return x + y
//
//}

////无返回值
//func foo() {
//	fmt.Println("hello world!")
//	return //不写return默认return空
//}

////返回多个值
//func get_name_age() (string, int) {
//	return "daerwen", 18
//}

//返回值命名
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return //return sum sub

}
func main() {

	//调用函数
	calc(3, 5)

	//name, age := get_name_age()
	//fmt.Println(name, age)

	////调用函数
	////ret := foo() //报错:无返回值不能将调用的结果作为
	//foo()

	//ret := add_cal(12, 20)
	//fmt.Println(ret)

	//ret := sum(12, 13, 12, 14)
	//fmt.Println(ret)

	//练习1:
	//计算1-100的和
	//正常写法
	//var sum = 0
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//函数写法
	//调用1-100的和的函数
	//sum1to100()

	//调用带参数菱形函数
	//printLingxing(10)

	////调用带参数的1toN的和
	//sum1toN(100)
	//sum1toN(150)
	//sum1toN(200)

	//函数调用,按顺序传参
	//addSum(2) //报错
	//addSum(1,2,3)  //报错
	//addSum(100, 20)

	////不定长参数函数调用
	//sum(1, 2, 3)
	//sum(100, 300, 100, 200)

}
