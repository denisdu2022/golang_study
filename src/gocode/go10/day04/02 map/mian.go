package main

func main() {

	////0. map的声明
	//var stus map[string]string
	//fmt.Println(stus, reflect.TypeOf(stus)) //map[] map[string]string
	//
	////1. map 声明并初始化
	////var x = []int{1,2,3}  //切片的声明并初始化
	//var stus = map[string]string{"name": "daernwen", "age": "22"}
	////支持key查询  map[]对象
	//fmt.Println(stus) //map[age:22 name:daernwen]  是无顺序的
	////查询"age" 这个键的值
	//fmt.Println(stus["age"]) //22
	////写入key-value
	//stus["gender"] = "male"
	//fmt.Println(stus)      //map[age:22 gender:male name:daernwen]
	//fmt.Println(len(stus)) // 3
	//stus["height"] = "180cm"
	//fmt.Println(stus["height"]) // 180cm
	////修改key-value
	//stus["height"] = "175cm"
	//fmt.Println(stus["height"]) //175cm
	////删除key-value
	//delete(stus, "height")
	//fmt.Println(stus)  //map[age:22 gender:male name:daernwen]
	//
	//2. 基于make函数声明初始化  ,make引用类型 ,存储的地址
	//
	//var stus1 map[string]string
	//stus1{"name": "daerwen"} //会报错,引用类型,没有默认开辟key-value的空间,nli map
	//var stu02 = make(map[string]string)
	//stu02["name"] = "daerwen"
	//fmt.Println(stu02, reflect.TypeOf(stu02)) //map[name:daerwen] map[string]string
	//stu02["age"] = "22"
	//stu02["gender"] = "male"
	//fmt.Println(stu02) //map[age:22 gender:male name:daerwen]
	//
	////练习:
	////年龄自加1运算
	//var a = stu02["age"]
	//b, _ := strconv.Atoi(a)
	//b++
	//fmt.Println(b, reflect.TypeOf(b))
	//
	////使用空接口 interface  任意类型
	//var stu03 = make(map[string]interface{})
	//stu03["age"] = 18
	//var sum = stu03["age"]
	//fmt.Println(sum, reflect.TypeOf(sum)) //18 int
	//
	////遍历map
	////切片的遍历
	//var s2 = []int{10, 11, 12, 14, 14, 15}
	//for i, v := range s2 {
	//	fmt.Println(i, v)
	//}
	//
	////map的遍历
	////for range 无序打印
	//var stus = map[string]string{"name": "daernwen", "age": "22"}
	//for k, v := range stus {
	//	fmt.Println(k, v)
	//}
	///*
	//	name daernwen
	//	age 22
	//*/
	//
	////for range 无序打印
	//newMap := map[int]int{
	//	1: 1,
	//	2: 2,
	//	3: 3,
	//	4: 4,
	//	5: 5,
	//	6: 6,
	//}
	//
	//for k, v := range newMap {
	//	fmt.Println(k, v)
	//}
}
