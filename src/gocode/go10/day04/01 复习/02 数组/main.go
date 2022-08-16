package main

import (
	"fmt"
	"reflect"
)

func main() {
	//1. 先声明在赋值
	//数组的声明
	//数组必须限制长度
	var arr [3]int
	fmt.Println(arr) //[0 0 0 ]

	//赋值  数组[索引]
	fmt.Println(arr[0]) //0
	fmt.Println(arr[1]) //0
	fmt.Println(arr[2]) //0

	//索引赋值
	arr[0] = 10
	fmt.Println(arr) //[10 0 0]
	arr[1] = 11
	arr[2] = 12
	fmt.Println(arr) //[10 11 12]

	//2. 数组的声明并赋值
	var names = [3]string{"daerwen", "dafenqi", "bijialuo"}
	fmt.Println(names) //[daerwen dafenqi bijialuo]

	var ages = [3]int{22, 23, 24}
	fmt.Println(ages) //[22 23 24]

	//3. 省略长度赋值
	var names1 = [...]string{"daerwen", "dafenqi"}
	fmt.Println(names1)                              //[daerwen dafenqi]
	fmt.Println(len(names1), reflect.TypeOf(names1)) //2 [2]string  len()计算容器数据的长度

	//4. 索引设置
	var names2 = [...]string{0: "dafenqi", 2: "dafenqi"}
	fmt.Println(names2) //[dafenqi  dafenqi]
	var ages1 = [...]int{0: 22, 3: 33}
	fmt.Println(ages1) //[22 0 0 33]

}
