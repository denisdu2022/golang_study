package main

import "fmt"

func f1() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}
func f2() *int {

	i := 5
	defer func() {
		i++
		fmt.Printf(":::%p\n", &i)
	}()
	fmt.Printf(":::%p\n", &i)
	return &i
}

func f3() (result int) {
	defer func() {
		result++
	}()
	return 5 // result = 5;ret result(result替换了rval)
}

func f4() (result int) {
	defer func() {
		result++
	}()
	return result // ret result变量的值
}

func f5() (r int) {
	t := 5
	defer func() {
		t = t + 1
	}()
	return t // ret r = 5 (拷贝t的值5赋值给r)
}

func f6() (r int) {
	fmt.Println(&r)
	defer func(r int) {
		r = r + 1
		fmt.Println(&r)
	}(r)
	return 5
}

func f7() (r int) {
	defer func(x int) {
		r = x + 1
	}(r)
	return 5
}

func main() {

	// println(f1())
	fmt.Println(f1())
	// println(*f2())
	fmt.Println(*f2())
	// println(f3())
	fmt.Println(f3())
	// println(f4())
	fmt.Println(f4())
	// println(f5())
	fmt.Println(f5())
	// println(f6())
	// println(f7())

}
