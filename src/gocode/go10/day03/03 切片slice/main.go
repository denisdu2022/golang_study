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

	////type Slice struct {
	////	Data uintptr //指针,指向底层数组中切片指定的开始位置
	////	Len  int     //长度,即切片的长度
	////	Cap  int     //最大长度(容量),也就是切片开始位置到数组的最后位置的长度
	////}
	//
	////示例
	//slice := []int{1, 2, 3, 4, 5}
	//newSlice := slice[1:3]
	//
	//fmt.Println(slice)         //[1 2 3 4 5]
	//fmt.Println(len(slice))    //5
	//fmt.Println(cap(slice))    //5
	//fmt.Println(newSlice)      //[2 3]
	//fmt.Println(len(newSlice)) //2
	//fmt.Println(cap(newSlice)) //4
	//
	//fmt.Println(&slice)    //&[1 2 3 4 5]
	//fmt.Println(&newSlice) //&[2 3]
	//
	////地址打印
	////a 记录的顺序表的表头,即数组指针 : 数组指针 长度 容量,所以打印a地址即打印a第一个表头地址
	//fmt.Printf("%p\n", slice)     //0xc00000a330
	//fmt.Printf("%p\n", &slice)    //0xc000004078
	//fmt.Printf("%p\n", &slice[0]) //0xc00000a330
	//fmt.Printf("%p\n", &slice[1]) //0xc00000a338
	//fmt.Printf("%p\n", &slice[2]) //0xc00000a340
	//fmt.Printf("%p\n", &slice[3]) //0xc00000a348
	//fmt.Printf("%p\n", &slice[4]) //0xc00000a350

	////练习
	//var a = [...]int{1, 2, 3, 4, 5, 6}
	//fmt.Println(a, reflect.TypeOf(a)) //[1 2 3 4 5 6] [6]int
	//
	//a1 := a[0:3] //[1 2 3]   3    6
	//a2 := a[0:5] //[1 2 3 4 5 ] 5 6
	//a3 := a[1:5] //[2 3 4 5 ] 4  5
	//fmt.Println(a3, reflect.TypeOf(a3))
	//a4 := a[1:]   //[2 3 4 5 6]  5  5
	//a5 := a[:]    //[1 2 3 4 5 6]  6 6
	//a6 := a3[1:2] //[3] 1  3
	//fmt.Println(a6, reflect.TypeOf(a6))
	//
	//fmt.Printf("a1的长度%d,容量%d\n", len(a1), cap(a1))
	//fmt.Printf("a2的长度%d,容量%d\n", len(a2), cap(a2))
	//fmt.Printf("a3的长度%d,容量%d\n", len(a3), cap(a3))
	//fmt.Printf("a4的长度%d,容量%d\n", len(a4), cap(a4))
	//fmt.Printf("a5的长度%d,容量%d\n", len(a5), cap(a5))
	//fmt.Printf("a6的长度%d,容量%d\n", len(a6), cap(a6))

	//var names = []string{"张三", "李四", "王五"}
	//fmt.Println(names, reflect.TypeOf(names)) //[张三 李四 王五] []string

}
