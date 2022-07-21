package main //声明包的名字,main  可执行文件

import "fmt"

//函数外只能放置标识符(声明变量,常量,函数类型)

//main 程序的入口,main包没有参数也没有返回值
// func main() {
// 	fmt.Println("  $$$$$$$$  helloword      $$$$$$$$$$$$$$")
// }

//go语言的变量声明格式
//var name string //声明一个叫name的字符串类型的变量
//var age int

//批量声明变量
// var (
// 	name string // ""
// 	age  int    // 0
// 	//go语言中非全局变量,声明必须使用,不使用就无法编译
// )

// func main() {
// 	//fmt.Print("hahahhaha   hahah") //在终端输出,打印内容,不换行
// 	// fmt.Printf("sdadasd  dasdas") //打印内容,会格式化
// 	fmt.Println("dasdasd dasdsada") //打印完指定内容会增加换行符
// }

//go语言中推荐使用驼峰命名法
// var student_name string
var studentNmae string //推荐使用小驼峰
// var StudentName string

// s5 := "呵呵呵"
// 全局变量中不能使用简短变量

// func main() {
// 	fmt.Println(studentNmae)
// 	//声明变量同时赋值
// 	var s1 string = "海斌"
// 	//类型推导,根据值判断变量的类型
// 	var s2 = 20
// 	//简短变量声明,只能在函数中使用
// 	s3 := "哈哈哈"
// 	//s1 := 10  同一个作用域({})中不能重复声明同名的变量

// 	//匿名变量是一个特殊的变量

// 	fmt.Println(s1)
// 	fmt.Println(s2)
// 	fmt.Println(s3)
// }

//常量
//常量定义之后不能修改,常量是不能变化的
//在程序运行期间不会改变
const pi = 3.1415926

//批量声明常量
const (
	statusok = 200
	notfound = 404
)

//批量声明常量时,如果某一行声明后没有赋值,默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

func main() {
	fmt.Println(n1, n2, n3)
}
