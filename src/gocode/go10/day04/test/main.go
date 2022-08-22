package main

import (
	"fmt"
	"strconv"
)

func main() {

	// （1）map嵌套slice
	var data = make(map[string][]string)
	data["beijing"] = []string{"朝阳", "海淀", "昌平"}
	data["shandong"] = []string{"济南", "青岛", "威海"}
	data["hebei"] = []string{"唐山", "石家庄", "保定"}
	// fmt.Println(data) // map对象
	// 1. 查询河北的第二个城市
	// fmt.Println(data["hebei"][1])
	// 2. 遍历每一个省份以及对应的城市名
	/*for proStr, citysSlice := range data {
		//fmt.Println(proStr, citysSlice)
		fmt.Println(proStr, len(citysSlice))
		for i, city := range citysSlice {
			fmt.Printf("%d.%s", i, city)
		}
		fmt.Println()

	}*/

	// 3. 删除北京的key-value
	//delete(data, "beijing")
	//fmt.Println(data)

	// （2）map嵌套map
	/*stu01 := map[string]string{"name": "rain", "age": "32"}
	stu02 := map[string]string{"name": "eric", "age": "33"}
	stu03 := map[string]string{"name": "alvin", "age": "34"}
	var stus = make(map[int]map[string]string)
	stus[1001] = stu01
	stus[1002] = stu02
	stus[1003] = stu03
	fmt.Println(stus) //map[1001:map[age:32 name:rain] 1002:map[age:33 name:eric] 1003:map[age:34 name:alvin]]
	fmt.Println(len(stus))
	// 打印学号为1002的学生的年龄
	fmt.Println(stus[1002]["age"])
	// 循环打印每个学员的学号，姓名，年龄
	for sid, stuMap := range stus {
		fmt.Println(sid, stuMap["name"], stuMap["age"])
	}
	*/
	// （3）切片嵌套map

	// 方式1
	stu01 := map[string]string{"name": "rain", "age": "32"}
	stu02 := map[string]string{"name": "eric", "age": "33"}
	stu03 := map[string]string{"name": "alvin", "age": "34"}
	// var stus = make([]map[string]string, 3)
	var stus = []map[string]string{stu01, stu02, stu03}
	// 方式2
	//var stus = []map[string]string{{"name": "rain", "age": "32"}, {"name": "eric", "age": "32"}}

	// 1. 打印第二个学生的姓名
	fmt.Println(stus[1]["name"])
	// 2. 循环打印每一个学生的姓名和年龄
	for _, stuMap := range stus {
		//fmt.Println(stuMap["name"], stuMap["age"])
		fmt.Printf("姓名：%-8s 年龄：%-8s\n", stuMap["name"], stuMap["age"])
	}

	// 3. 添加一个新的学生map
	/*var name, age string
	fmt.Print("请输入新学生的年龄和姓名：")
	fmt.Scan(&name, &age)
	newMap := map[string]string{"name": name, "age": age}
	stus = append(stus, newMap)
	fmt.Println(stus)*/

	// 4. 删除一个学生eric的map
	/*// 查询eirc的索引位置
	var deleteIndex = 0
	for index, stuMap := range stus {
		if stuMap["name"] == "rain" {
			deleteIndex = index
		}
	}
	// a = append(a[:index], a[index+1:]...)
	stus = append(stus[:deleteIndex], stus[1+deleteIndex:]...)
	fmt.Println(stus)*/

	// 将姓名为rain的学生的年龄自加一岁
	for _, stuMap := range stus {
		if stuMap["name"] == "rain" {
			// 类型转换
			age, _ := strconv.Atoi(stuMap["age"])
			stuMap["age"] = strconv.Itoa(age + 1)
		}
	}
	fmt.Println(stus)

}
