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
	var stuMap = map[string]interface{}{"name": "daerwen", "age": 22, "addr": "hebei"}

	//实例化学生对象
	var stuSruct = Stu{"dafenqi", 18, Addr{"hebei", "langfang"}}

	//map序列化
	jsonStuMap, _ := json.Marshal(stuMap)
	//{"addr":"hebei","age":22,"name":"daerwen"}
	fmt.Println(string(jsonStuMap))

	//结构体序列化
	jsonStuStruct, _ := json.Marshal(stuSruct)
	//{"Name":"dafenqi","Age":18,"Sheng":"hebei","Shi":"langfang"}
	fmt.Println(string(jsonStuStruct))

	//map反序列化
	var StuMap map[string]interface{}
	err := json.Unmarshal(jsonStuMap, &StuMap)
	if err != nil {
		return
	}
	//map[addr:hebei age:22 name:daerwen]
	fmt.Println(stuMap, stuMap["name"])

	//结构体反序列化
	var StuStructdata Stu
	err = json.Unmarshal(jsonStuStruct, &StuStructdata)
	if err != nil {
		return
	}
	//{dafenqi 18 {hebei langfang}} dafenqi
	fmt.Println(StuStructdata, StuStructdata.Name)

}
