package main

import "fmt"

func main() {

	////练习1:
	//var s = []int{1, 2, 3}
	//fmt.Printf(" %p\n", &s) //0xc000114060
	//fmt.Println(s)
	//s = append(s, 4)
	//fmt.Printf(" %p\n", &s) //0xc000114060
	////应为是同一个变量地址都是同一个
	////但是s这个切片是扩容了,扩容后的值变了
	//fmt.Println(s, len(s), cap(s)) //[1 2 3 4]  4 6

	////在开头添加元素
	//var a = []int{1, 2, 3}
	////在a这个切片的前边追加一个元素0
	//fmt.Println(append([]int{0}, a...)) //[0 1 2 3]
	////追加多个元素
	//fmt.Println(append([]int{-1, -2, -3}, a...)) //[-1 -2 -3 1 2 3]
	//
	////在任意位置增加元素
	//var b = []int{1, 2, 3, 4, 5}
	//var i = 2
	//fmt.Println(append(b[:i], append([]int{10}, b[i:]...)...)) //[1 2 10 3 4 5]

	//删除元素
	var c = []int{1, 2, 3, 4, 5}
	//要删除元素的索引,要删除索引+1
	c = append(c[:3], c[3+1:]...)
	fmt.Println(c) //[1 2 3 5]

}
