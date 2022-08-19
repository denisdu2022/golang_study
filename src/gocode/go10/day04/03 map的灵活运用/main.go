package main

import "fmt"

func main() {
	////案例1: map[string][]string
	//data := map[string][]string{"hebei": []string{"廊坊", "石家庄", "香河"}, "beijing": []string{"朝阳", "海淀", "丰台"}, "shandong": []string{"济南", "德州", "泰安"}}
	////fmt.Println(data)
	////map[beijing:[朝阳 海淀 丰台] hebei:[廊坊 石家庄 香河] shandong:[济南 德州 泰安]]
	////打印河北第二个城市
	////fmt.Println(data["hebei"][1]) //石家庄
	////循环打印每个省份的名字和城市数量
	////for i, v := range data {
	////	fmt.Println(i)
	////	for k, v1 := range v {
	////		fmt.Println(k, v1)
	////	}
	////}
	///*	hebei
	//	0 廊坊
	//	1 石家庄
	//	2 香河
	//	beijing
	//	0 朝阳
	//	1 海淀
	//	2 丰台
	//	shandong
	//	0 济南
	//	1 德州
	//	2 泰安*/
	////添加一个新的省份和城市的key-value
	//data["henan"] = []string{"新乡", "郑州", "开封"}
	//data["shanxi"] = []string{"大同", "太原", "晋中"}
	//fmt.Println(data)
	////map[beijing:[朝阳 海淀 丰台] hebei:[廊坊 石家庄 香河] henan:[新乡 郑州 开封] shandong:[济南 德州 泰安] shanxi:[大同 太原 晋中]]
	////删除北京的key-value
	//delete(data, "beijing")
	//fmt.Println(data)
	////map[hebei:[廊坊 石家庄 香河] henan:[新乡 郑州 开封] shandong:[济南 德州 泰安] shanxi:[大同 太原 晋中]]

	////案例2: map[int]map[string]string
	//info := map[int]map[string]string{1001: {"name": "daerwen", "age": "22"}, 1002: {"name": "dafenqi", "age": "18"},
	//1003: {"name": "bijieluo", "age": "21"}}
	////打印学号为1002的学生的年龄
	////fmt.Println(info[1002]["age"])  //18
	////循环打印每个学员的学号,姓名,年龄
	//for i, v := range info {
	//	fmt.Println(i)
	//	for i1, v1 := range v {
	//		fmt.Println(i1, v1)
	//	}
	//}
	//添加一个新的学员
	//info[1004] = map[string]string{"name": "eice", "age": "22"}
	//fmt.Println(info)
	////删除1003的学生
	//delete(info, 1003)
	//fmt.Println(info)
	////map[1001:map[age:22 name:daerwen] 1002:map[age:18 name:dafenqi]]

	//info := map[int]map[string]string{1001: {"name": "daerwen", "age": "22", "gender": "male", "height": "180cm"},
	//1002: {"name": "dafenqi", "age": "18", "gender": "male", "height": "175cm"}, 1003: {"name": "bijieluo", "age": "21", "gender": "male", "height": "178"}}

	//案例3: []map[string]string
	//stus := []map[string]string{{"name": "daerwen", "age": "22"}, {"name": "bijiasuo", "age": "21"}, {"name": "bijialuo", "age": "33"}}
	//打印第二个学生的姓名
	//fmt.Println(stus[1]["name"]) //bijiasuo
	//循环打印每个学生的姓名和年龄
	//for _, v := range stus {
	//	for i1, v1 := range v {
	//		fmt.Println(i1, v1)
	//
	//	}
	//}
	///*	name daerwen
	//	age 22
	//	name bijiasuo
	//	age 21
	//	name bijialuo
	//	age 33*/
	//添加一个新学生的map
	//stus = append(stus, map[string]string{"name": "eieck", "age": "22"})
	//fmt.Println(stus)
	//删除一个学生map
	//stus = append(stus[:2], stus[3:]...)
	//fmt.Println(stus)
	//[map[age:22 name:daerwen] map[age:21 name:bijiasuo]]
	////将姓名为bijiesuo的学生自加一岁
	//fmt.Println(stus[1]["age"])
	//var ages = stus[1]["age"]
	//var ageNum, _ = strconv.Atoi(ages)
	//ageNum++
	//fmt.Println(ageNum) //22

	//联系:
	//根据age的大小重新排序
	stus := []map[string]int{map[string]int{"age": 23}, map[string]int{"age": 33}, map[string]int{"age": 18}}
	fmt.Println(stus)
}
