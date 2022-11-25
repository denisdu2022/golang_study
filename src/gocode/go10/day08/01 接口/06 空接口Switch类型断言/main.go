package main

import "fmt"

func justifType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string类型,value is %v\n", v)
	case int:
		fmt.Printf("x is a int类型,value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool类型,value is %v\n", v)
	default:
		fmt.Println("unsupport type!")
	}
}

func main() {
	justifType(22)      //x is a int类型,value is 22
	justifType("hello") //x is a string类型,value is hello
	justifType(true)    //x is a bool类型,value is true
	justifType([]int{}) //unsupport type!

}
