package main

func main() {

	//00 切片的语法
	//arr [start : end] 或者 slice [start : end] //start开始索引 :  end 结束索引
	//
	//var arr = [5]int{10, 11, 12, 13, 14}
	//var s1 = arr[1:4]
	//fmt.Println(s1, reflect.TypeOf(s1)) //11,12,13 []int
	//var s2 = arr[2:5]
	//fmt.Println(s2, reflect.TypeOf(s2)) //12,13,14 []int
	//var s3 = s2[0:2]
	//fmt.Println(s3, reflect.TypeOf(s3)) //12,13 []int
	//
	//s3[0] = 1000
	//fmt.Println(":::", s1, s2, s3)
	//
	////01 值类型
	//var a int
	//var b string
	//var c bool
	//var d [2]int
	//fmt.Println(a)  //int类型默认值为 0
	//fmt.Println(b)  //string 类型默认值为 ""
	//fmt.Println(c)  //bool类型默认值为false
	//fmt.Println(d)  //数组默认值为 [0 0]
	//
	////整形赋值
	//var a = 10
	//b := a
	//b = 101
	//fmt.Printf("a: %v,a的内存地址是%p\n", a, &a) //10
	//fmt.Printf("b: %v,b的内存地址是%p\n", b, &b) //101
	//
	////数组赋值
	//var c = [3]int{1, 2, 3}
	//d := c
	//d[1] = 100
	//fmt.Printf("c: %v,c的内存地址是%p\n", c, &c) // [1 2 3]
	//fmt.Printf("d: %v,d的内存地址是%p\n", d, &d) // [1 100 3]
	//
	////引用类型
	//var a = []int{1, 2, 3}
	//b := a
	//a[0] = 100
	//fmt.Println(a) //[100 2 3]
	//fmt.Println(b) //[100 2 3]

}
