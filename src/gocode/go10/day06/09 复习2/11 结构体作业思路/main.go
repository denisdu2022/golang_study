package main

import (
	"encoding/json"
	"fmt"
)

//省份城市结构体
type Addr struct {
	Sheng string
	Shi   string
}

//学生数据结构体
type Stu struct {
	Name string
	Age  int
	Addr
}

func main() {
	var s1 = Stu{"s1", 18, Addr{"hebei", "langfang"}}
	var s2 = Stu{"s2", 16, Addr{"hebei", "dachang"}}
	var s3 = Stu{"s3", 17, Addr{"hebei", "xianghe"}}

	var data = []Stu{s1, s2, s3}
	//学生信息序列化
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
	/*
		[{"Name":"s1","Age":18,"Sheng":"hebei","Shi":"langfang"},{"Name":"s2","Age":16,"Sheng":"hebei","Shi":"dachang"},{"Name":"s3","Age":17,"Sheng":"hebei","Shi":"xianghe"}]
	*/

	//学生信息反序列化
	var info []Stu
	json.Unmarshal(jsonData, &info)
	//[{s1 18 {hebei langfang}} {s2 16 {hebei dachang}} {s3 17 {hebei xianghe}}] s2
	fmt.Println(info, info[1].Name)

}
