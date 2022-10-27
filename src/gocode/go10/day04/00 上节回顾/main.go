package main

import "fmt"

func sum(x, y int) {
	var s = 0
	for i := x; i <= y; i++ {
		s += i
	}
	fmt.Println(s)
}

func sum1(num ...int) {
	fmt.Println("num: ", num)
	sum02 := 0
	for _, num := range num {
		sum02 += num
	}
	fmt.Println(sum02)

}

func main() {

	sum(1, 100)

	sum1(1, 2, 3, 4, 5)
	sum1(23, 25)

	////声明并赋值map
	//var students = map[string]string{"name": "daerwen", "age": "22"}
	//fmt.Println(students, reflect.TypeOf(students)) //map[age:22 name:daerwen] map[string]string
	//
	////查询map的键值
	//fmt.Println(students["name"])
	//fmt.Println(students["age"])
	//
	////新增map键值
	//students["gender"] = "male"
	//students["height"] = "170cm"
	//
	//fmt.Println(students) //map[age:22 gender:male height:170cm name:daerwen]
	//
	////修改map键值
	//students["height"] = "180cm"
	//fmt.Println(students) //map[age:22 gender:male height:180cm name:daerwen]
	//
	////删除一个map的键值
	//delete(students, "height")
	//fmt.Println(students) //map[age:22 gender:male name:daerwen]

	//案例1:
	//var data = map[string][]string{"北京": []string{"朝阳", "东城", "西城", "昌平"}, "山东": []string{"济南", "威海", "日照"}, "河北": []string{"廊坊", "香河", "大厂"}}
	//fmt.Println(data)          //map[北京:[朝阳 东城 西城 昌平] 山东:[济南 威海 日照] 河北:[廊坊 香河 大厂]]
	//fmt.Println(data["北京"][1]) //东城

	//循环打印每个省份的的名字和城市数量
	//for proStr, citys := range data {
	//	fmt.Println(proStr)
	//	for i, city := range citys {
	//		fmt.Printf("%d.%s", i, city)
	//	}
	//	fmt.Println()
	//	/*
	//		北京
	//		0.朝阳1.东城2.西城3.昌平
	//		山东
	//		0.济南1.威海2.日照
	//		河北
	//		0.廊坊1.香河2.大厂
	//	*/
	//}

	////打印河北第二个城市
	//fmt.Println(data["河北"][1]) //香河

	////添加一个新的省份和城市的key-value
	//data["河南"] = []string{"洛阳", "郑州", "开封"}
	//fmt.Println(data) //map[北京:[朝阳 东城 西城 昌平] 山东:[济南 威海 日照] 河北:[廊坊 香河 大厂] 河南:[洛阳 郑州 开封]]

	////删除北京的key-value
	//delete(data, "北京")
	//fmt.Println(data) //map[山东:[济南 威海 日照] 河北:[廊坊 香河 大厂]]

	//案例2:
	//info := map[int]map[string]string{1001: {"name": "yuan", "age": "23"}, 1002: {"name": "alvin", "age": "33"}}
	//打印学号为1002学生的年龄
	//fmt.Println(info[1002]["age"])

	////循环打印每个学员的学号,姓名,年龄
	//for k, v := range info {
	//	fmt.Printf("%d ", k)
	//	for i, b := range v {
	//		fmt.Printf("%s:", i)
	//
	//		fmt.Printf("%s ", b)
	//	}
	//
	//}

	////添加一个新的学员
	//info[1003] = map[string]string{"name": "daerwen", "age": "21"}
	//fmt.Println(info)
	////map[1001:map[age:23 name:yuan] 1002:map[age:33 name:alvin] 1003:map[age:21 name:daerwen]]

	////删除1001的学生
	//delete(info, 1001)
	//fmt.Println(info) //map[1002:map[age:33 name:alvin]]

}
