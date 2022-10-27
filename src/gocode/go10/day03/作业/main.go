package main

func main() {
	////1.
	//s := make([]int, 5) //[0 0 0 0 0] 5 5
	//s = append(s, 1, 2, 3)
	//fmt.Println(s, len(s), cap(s)) //[0 0 0 0 0 1 2 3]  8  10

	////2.
	//p1 := 1
	//p2 := &p1
	//*p2++
	//fmt.Println(p1)  //2
	//fmt.Println(&p1) //取址0x1400012c008
	//fmt.Println(*p2) //2   取值
	//fmt.Println(&p1) //取址0x1400012c008

	////3.代码执行结果
	//a := [3]int{0, 1, 2}
	//s := a[1:2] // 1 1 2
	//
	//s[0] = 11         //s = 11   1 2
	//s = append(s, 12) //11 12  2 2
	//fmt.Println(s, len(s), cap(s))
	//s = append(s, 13) //11 12 13 3 4
	//fmt.Println(s, len(s), cap(s))
	//
	//s[0] = 21
	//
	//fmt.Println(a, len(a), cap(a)) //0 11 12 3 3
	//fmt.Println(s, len(s), cap(s)) //21 12 13  3 4

	////4.
	//var a = [5]int{1, 2, 3, 4, 5}
	//var r [5]int
	//for i, v := range &a {
	//	if i == 0 {
	//		a[1] = 12
	//		a[2] = 13
	//	}
	//	r[i] = v
	//}
	//fmt.Println("r = ", r)
	//fmt.Println("a = ", a)

	////05 切片的反转 s := []int{1,2,3,5,4}
	//s := []int{1, 2, 3, 5, 4}
	//
	////定义一个空的切片
	//s1 := []int{}
	//
	////遍历取值
	//for i := range s {
	//	//定义一个变量n,n的值,为s切片的长度-1在减去循环的次数-i
	//	n := s[len(s)-1-i]
	//	//得到值为 4 5 3 2 1
	//	fmt.Println(n)
	//	//循环追加到空的切片
	//	s1 = append(s1, n)
	//}
	//
	////打印翻转后的值
	//fmt.Println(s1)

}
