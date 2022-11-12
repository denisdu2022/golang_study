package main

func main() {

	////构建切片方式一: 通过数组切片操作获得切片对象
	//var arr = [3]string{"daerwen", "dafenqi", "bijialuo"}
	//fmt.Println(arr, reflect.TypeOf(arr)) //[daerwen dafenqi bijialuo] [3]string  数组
	//s1 := arr[0:2]
	//fmt.Println(s1, reflect.TypeOf(s1)) //[daerwen dafenqi] []string 切片
	//s2 := arr[1:]
	//fmt.Println(s2, reflect.TypeOf(s2)) //[dafenqi bijialuo] []string
	//
	////切片是对数组的引用
	//var a = [5]int{1, 2, 3, 4, 5}
	//var slice = a[:] //起始地址 长度 容量
	////长度是后边减去前边,容量是减去起始位置,或者从起始位置到最后
	//fmt.Println(len(slice), cap(slice))
	//var slice2 = a[:2]
	//fmt.Println(len(slice2), cap(slice2))
	//fmt.Println(slice) //[1 2 3 4 5]
	//newSlice := slice[1:3]
	//fmt.Println(newSlice) //[2 3]
	//fmt.Println(len(newSlice), cap(newSlice))
	//
	//newSlice[1] = 100
	//fmt.Println(a)      //[1 2 100 4 5]
	//fmt.Println(slice)  //[1 2 100 4 5]
	//fmt.Println(slice2) //[1 2]
	//
	////2. 直接声明切片
	//var s = []int{10, 11, 12, 13, 14} //底层会创建一个数组,能拿到的是切片
	//fmt.Println(s)
	//s1 := s[1:4]
	//fmt.Println(s1, len(s1), cap(s1)) //[11 12 13]  3 4
	//
	//s2 := s1[1:]
	//fmt.Println(s2, len(s2), cap(s2)) //[12 13]  2  3

}
