package main

import (
	"fmt"
	"sort"
)

func main() {

	////开头添加元素
	//var a = []int{1, 2, 3}
	//fmt.Println(a, len(a), cap(a)) //[1 2 3] 3 3
	//fmt.Printf("%p\n", &a)
	//a = append([]int{0}, a...)     //在开头添加1个元素
	//fmt.Println(a, len(a), cap(a)) //[0 1 2 3]  4 4
	//fmt.Printf("%p\n", &a)
	//a = append([]int{-3, -2, -1}, a...) //在开头添加一个切片
	//fmt.Printf("%p\n", &a)
	//fmt.Println(a, len(a), cap(a)) //[-3 -2 -1 0 1 2 3]  7  8

	//任意位置插入元素
	//var a []int
	//a = append(a[:i], append([]int{x}, a[i:]...)...) //在第i个位置插入x
	//a = append(a[:i], append([]int{1, 2, 3}, a[i:]...)...)

	//var a = []int{1, 2, 3}
	////a = append(a[:1], append([]int{12}, a[1:]...)...)
	////fmt.Println(a, len(a), cap(a)) //[1 12 2 3] 4  6
	////
	////a = append(a[1:], append([]int{-1, 1, 2, 3}, a[1:]...)...)
	////fmt.Println(a, len(a), cap(a)) //[2 3 -1 1 2 3 2 3] 8 8
	//
	//a = append(a[1:2], append([]int{-1, -2, 3}, a[1:2]...)...)
	//fmt.Println(a, len(a), cap(a))

	//var a = []int{1, 2, 3, 4}
	//s1 := a[:2]
	//s2 := a[2:]
	//fmt.Println(append(append(s1, 100)), s2...)
	//
	////从切片中删除元素
	//a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	////要删除索引为2的元素
	//a = append(a[:2], a[3:]...)
	//fmt.Println(a)

	//a := [...]int{1, 2, 3}
	//b := a[:]
	//fmt.Println(b) //[1 2 3]
	//b = append(b[:1], b[2:]...)
	//fmt.Println(a, len(a), cap(a)) //[1 3 3] 3 3
	//fmt.Println(b, len(b), cap(b)) // [1 3] 2 3

	//从小到大排序
	a := []int{10, 2, 3, 100}
	sort.Ints(a)
	fmt.Println(a) //[2 3 10 100]

	//按字母排序
	b := []string{"daerwen", "dafenqi", "apple", "bijialuo"}
	sort.Strings(b)
	fmt.Println(b)

	//从小到大排序
	c := []float64{3.14, 5.25, 1.12, 4, 78}
	sort.Float64s(c)
	fmt.Println(c)

	//如果是一个数组,需要先转成切片在排序[:]
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	sort.Sort(sort.Reverse(sort.Float64Slice(c)))
	fmt.Println(a, c)

}
