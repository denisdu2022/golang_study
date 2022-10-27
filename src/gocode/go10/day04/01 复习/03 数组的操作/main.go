package main

func main() {

	//var names = [3]string{"daerwen", "dafenqi", "bijialuo"}
	////1. 索引操作 数组[索引]
	////索引查询
	//fmt.Println(names[2]) //bijialuo
	////索引修改
	//names[2] = "Bjialuo"
	//fmt.Println(names) //[daerwen dafenqi Bjialuo]
	////2. 切片操作 数组[start索引:end索引] 左避右开(顾头不顾尾)
	//var arr = [5]int{11, 12, 13, 14, 15}
	//s := arr[1:3]
	//fmt.Println(s, reflect.TypeOf(s)) //[12 13] []int
	////从1切到尾
	//s1 := arr[1:]
	//fmt.Println(s1) //[12 13 14 15]
	////从开头切到3
	//s2 := arr[:3]
	//fmt.Println(s2) //[11,12,13]
	////元素数量判断
	//s3 := arr[2:4]  //后边减去前边 2
	//fmt.Println(s3) //[13 14]
	//
	////3. 遍历数组
	//var arr2 = [5]int{20, 21, 22, 23, 24}
	////三要素循环
	//for i := 0; i < len(arr2); i++ {
	//	fmt.Println(i, arr2[i])
	//}
	////range循环
	//for i, v := range arr2 {
	//	fmt.Println(i, v)
	//}

}
