package main

import "fmt"

// 想要放多个person的信息到一个集合中
// var persons [][]string
var persons [][]interface{}

//个人信息的元素: name , age ,address ,height

//类型集合

type Person struct {
	name    string
	age     int
	address string
	height  float32
}

func main() {
	//实例化具体的数据
	//persons := append(persons,[]string{"lili","18","beijing","1.80"})
	//persons := append(persons,[]interface{}{"lili",18,"beijign","1.80"})

	//如何初始化结构体
	////方法1: 省略写法,每个值必填
	//p1 := Person{"lili",18,"beijing",1.80}
	////方法2: 灵活性强,可以只写需要的值
	//p2 := Person{
	//	name: "libai",
	//	age: 18,
	//}
	//
	////方法3:
	//var persons []Person
	//persons = append(persons,p1)
	//persons = append(persons,Person{
	//	name: "libai",
	//	age: 18,
	//	address: "beijing",
	//})
	//
	//persons2 := []Person{
	//	{"lili",18,"beijing",1.80},
	//	{
	//		age: 19,
	//	},
	//}

	////也可以直接赋值
	//var p Person
	//p.age = 18
	//fmt.Println(p.age)

	//匿名结构体
	address := struct {
		province string
		city     string
		address  string
	}{
		"北京市",
		"通州区",
		"xxxxx",
	}

	fmt.Println(address.city)

}
