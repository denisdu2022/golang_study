package main

import "fmt"

//go语言之后推荐使用驼峰命名法
//声明变量
//var name string
//var age int
//var isok bool

//批量声明
var (
	name string //""
	age  int    //0
	isOK bool   //false
)

func main() {
	name = "理想"
	age = 18
	//isOk = true go语言中变量声明必须使用,不使用就编译不过去

	//%s占位符 使用name这个变量值去替换占位符
	fmt.Printf("name:%s\n", name) //\n是换行符
	fmt.Println(age)              //Println 打印完指定的内容之后会在后面加一个换行符
	fmt.Println()                 //不写值,快速打印空行
	fmt.Println(isOK)

	fmt.Println()
	//声明变量同时赋值
	var s1 string = "pikaqiu"
	fmt.Println(s1)
	var s2 = "20"
	fmt.Println(s2)

	fmt.Println()
	//简短变量声明,只能在函数中使用
	s3 := "哈哈哈哈哈"
	fmt.Println(s3)
	fmt.Println()

	//匿名变量
	//匿名变量是一个特殊变量:函数场景中使用

	//注意:变量不能重复声明(同一个作用域中不能重复声明同名的变量)

	











}
































