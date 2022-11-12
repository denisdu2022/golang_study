package main

////案例1 .
//func f1() int {
//	i := 5
//	defer func() {
//		i++
//	}()
//	return i
//}

////案例2 .
//func f2() *int {
//	i := 5
//	defer func() {
//		i++
//		fmt.Printf(":::%p\n", &i) //:::0x1400001c098
//	}()
//	fmt.Printf(":::%p\n", &i) //:::0x1400001c098
//	return &i
//
//}

////案例3 .
//func f3() (result int) {
//	defer func() {
//		result++
//	}()
//	return 5 //result = 5 ;ret result(result 替换了rval)
//}

////案例4 .
//func f4() (result int) { //返回值默认值是0
//	defer func() {
//		result++
//	}()
//	return result //ret result的值
//}

////案例5
//func f5() (r int) {
//	t := 5
//	defer func() {
//		t = t + 1
//	}()
//	return t //ret r =5 (拷贝t的值5付给r)
//}

////案例6
//func f6() (r int) { //返回值默认值是0
//	fmt.Println(&r)     //0x140000aa008
//	defer func(r int) { //匿名函数的参数 默认值也是0
//		r = r + 1
//		fmt.Println(&r) //0x140000aa020
//	}(r)
//	return 5
//}

//案例7
func f7() (r int) { //0
	defer func(x int) { //0
		r = x + 1
	}(r)
	return 5
}

func main() {

	////案例1 结果
	//fmt.Println(f1()) //5

	////案例2 结果
	//fmt.Println(*f2()) //6

	////案例3 结果
	//fmt.Println(f3()) //6

	//////案例4 结果
	//fmt.Println(f4()) //1

	//////案例5 结果
	//fmt.Println(f5()) //5

	//////案例6 结果
	//fmt.Println(f6()) //5

	////////案例7 结果
	//fmt.Println(f7()) //1

}
