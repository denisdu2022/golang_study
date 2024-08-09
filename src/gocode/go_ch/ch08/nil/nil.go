package main

import "fmt"

type Person struct {
	name string
	age  int
	f    *int
}

func main() {
	//nil 一般代表某些数据的零值
	/*
		不同类型的数据零值不一样
		bool false
		numbers 0
		string ""
		pointer nil
		slice nil
		map nil
		channel,interface,function nil
		struct  默认值不是nil  ,默认值是具体字段的默认值

	*/

	p1 := Person{
		name: "libai",
		age:  19,
	}
	p2 := Person{
		name: "libai",
		age:  19,
	}
	if p1 == p2 {
		fmt.Println("yes")
	}

	//err 的默认值是nil  所以在判断的时候,可以直接判断err 是否 ==nil
	//if err != nil {
	//	...
	//}

	//slice的默认值
	//var ps []Person  // nil slice
	//if ps == nil {
	//	fmt.Println("ps : nil")
	//}

	//var i int
	//if i != nil{  //是不能直接判断的
	//
	//}

	var ps2 = make([]Person, 0) //empty slice
	if ps2 == nil {
		fmt.Println("ps : nil")
	}

	var m map[string]string             //nil map
	var m2 = make(map[string]string, 0) //empty map
	if m2 == nil {
		fmt.Println("map : nil")
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

	for key, value := range m2 {
		fmt.Println(key, value)
	}

	// m["name"] = "libai" //panic: assignment to entry in nil map
	m2["name"] = "lili"
	fmt.Println(m2["name"])
}
