package main

import "fmt"

//type Student struct {
//	string //string string  光写一个类型,匿名字段就是把类型当做类型名,不写,直接写字段类型
//	int    //int  int
//}

// 学生住址信息结构体
type Addr struct {
	shen string
	shi  string
	xian string
}

// 学生信息结构体
type Student struct {
	name string
	age  int
	Addr //Addr == Addr Addr
}

func main() {
	//s1 := Student{"daerwen", 22}
	//fmt.Println(s1.string) //daerwen

	//实例化学生对象
	s1 := Student{"dafenqi", 22, Addr{"山东", "济南", "历城"}}
	fmt.Println(s1.Addr)     //{山东 济南 历城}
	fmt.Println(s1.Addr.shi) //济南
	fmt.Println(s1.shi)      //济南
}
