package main

func main() {

	////arr := []int{}
	//var arr []int
	//arr[0] = 1
	//fmt.Println(arr, reflect.TypeOf(arr))

	////如果是 var arr [2]int 默认值 0
	//var arr [2]int
	//arr[0] = 1
	//fmt.Println(arr, reflect.TypeOf(arr))

	//a := make([]int, 2)     //len 2 cap 2
	//b := make([]int, 2, 10) //len 2 cap 10
	//fmt.Println(a, b)       //[0 0]  [0 0]
	//fmt.Println(len(a), len(b))
	//fmt.Println(cap(a), cap(b))

	//a := make([]int, 5) //[0 0 0 0 0]
	//b := a[0:3]         //[0 0 0]
	//a[0] = 100  //不会发生内存分配操作.就是不会分配新的内存空间
	//fmt.Println(a) //[100 0 0 0 0]
	//fmt.Println(b) //[100 0 0]

	//var s []int
	////可以追加空切片一个值
	//s1 := append(s, 0)
	//fmt.Println(s1) //[0]
	//
	////可以追加多个值
	//s2 := append(s1, 1, 2, 3)
	//fmt.Println(s2) //[0 1 2 3]
	//
	////可以追加一个切片
	//s3 := append(s2, []int{4, 5, 6}...)
	//fmt.Println(s3) // [0 1 2 3 4 5 6]

}
