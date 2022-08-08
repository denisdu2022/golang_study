package main

func main() {
	//
	////00 数组声明语法
	//var names [5]string
	//fmt.Println(names, reflect.TypeOf(names)) // [   ]   [5]string
	//
	//var ages [3]int
	//fmt.Println(ages, reflect.TypeOf(ages)) //[0 0 0]  [3]int
	//
	////01 数组先声明在初始化
	//var names [5]string
	//var ages [5]int
	//
	//names[0] = "张三"
	//names[1] = "李四"
	//names[2] = "王五"
	//names[3] = "赵六"
	//names[4] = "孙七"
	//fmt.Println(names) //[张三 李四 王五 赵六 孙七]
	//
	//ages[0] = 22
	//ages[1] = 33
	//ages[2] = 56
	//ages[3] = 65
	//ages[4] = 77
	//fmt.Println(ages) //[22 33 56 65 77]
	//
	////02 声明并初始化
	//var names = [3]string{"daerwen", "bijialuo", "dafenqi"}
	//var ages = [3]int{22, 33, 66}
	//fmt.Println(names, reflect.TypeOf(names)) //[daerwen bijialuo dafenqi] [3]string
	//fmt.Println(ages, reflect.TypeOf(ages))   //[22 33 66] [3]int
	//
	////03 不限长度 [...]
	//var names = [...]string{"daerwen", "bijialuo", "dafenqi"}
	//var ages = [...]int{22, 33, 66}
	//fmt.Println(names, reflect.TypeOf(names)) //[daerwen bijialuo dafenqi] [3]string
	//fmt.Println(ages, reflect.TypeOf(ages))   //[22 33 66] [3]int
	//
	////04 索引设置
	//var names = [...]string{0: "daerwen", 2: "dafenqi"}
	//var ages = [...]int{0: 1, 2: 33}
	//fmt.Println(names, reflect.TypeOf(names)) //[daerwen bijialuo dafenqi] [3]string
	//fmt.Println(ages, reflect.TypeOf(ages))   //[22 33 66] [3]int
	//
	////05 修改数组元素
	//var names = [...]string{"张三", "李四", "王五", "赵六", "孙七"}
	//
	////索引取值
	//fmt.Println(names[2]) //"王五"
	//
	////修改元素值
	//names[0] = "dafenqi"
	//fmt.Println(names[0]) //"dafenqi"
	//
	////切片取值
	//fmt.Println(names[0:4]) //[dafenqi 李四 王五 赵六]
	//fmt.Println(names[0:])  //[dafenqi 李四 王五 赵六 孙七]
	//fmt.Println(names[:3])  //[dafenqi 李四 王五]
	//
	////循环取值
	////三要素循环取值
	//for i := 0; i < len(names); i++ {
	//	fmt.Println(i, names[i])
	//}
	//
	////range循环
	//for k, v := range names {
	//	fmt.Println(k, v)
	//}

}
