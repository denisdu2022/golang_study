package main

import (
	"fmt"
	"strings"
)

//func add(a, b int) int {
//	return a + b
//}

func add(a, b interface{}) interface{} {
	//类型断言
	//ai, ok := a.(int) //ok 代表布尔值
	//if !ok {
	//	panic("not int type")
	//}
	//bi, _ := b.(int)
	//
	//return ai + bi

	//switch 判断
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai + bi
	case float32:
		ai, _ := a.(float32)
		bi, _ := b.(float32)
		return ai + bi
	case float64:
		ai, _ := a.(float64)
		bi, _ := b.(float64)
		return ai + bi
	case string:
		as, _ := a.(string)
		bs, _ := b.(string)
		return as + bs
	default:
		panic("not supported type")
	}
}

func main() {
	//a := 1.0
	//b := 2.0
	//fmt.Println(add(a, b))

	re := add("hello:", "libai")
	//字符串分割
	//strings.Split(re,":")  //不可以直接分割,因为res的类型还是interface{}
	//断言
	res, _ := re.(string)
	ress := strings.Split(res, ":")
	fmt.Println(ress)

}
