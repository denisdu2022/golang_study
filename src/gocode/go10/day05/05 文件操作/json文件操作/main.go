package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {
	//json序列化

	////切片序列化
	//var s = []int{1, 2, 3}
	//fmt.Println(s, reflect.TypeOf(s)) //[1 2 3] []int
	//data, _ := json.Marshal(s)
	//fmt.Println(data, reflect.TypeOf(data)) //91 49 44 50 44 51 93] []uint8
	//fmt.Printf("%#v:\n", string(data))      //"[1,2,3]":
	//fmt.Println(string(data))               //[1,2,3]
	//ioutil.WriteFile("data.json", data, 0666)

	//map序列化
	var m = make(map[string]string)
	m["name"] = "daerwen"
	m["age"] = "22"
	data1, _ := json.Marshal(m)
	//map[age:22 name:daerwen] map[string]string
	fmt.Println(data1, reflect.TypeOf(string(data1))) //string
	fmt.Printf("%#v\n", string(data1))                //"{\"age\":\"22\",\"name\":\"daerwen\"}"
	fmt.Println(string(data1))                        //{"age":"22","name":"daerwen"}
	ioutil.WriteFile("data1.json", data1, 0666)

	//反序列化
	data2, _ := ioutil.ReadFile("data1.json")
	fmt.Println(data2)
	fmt.Println(string(data2)) //{"age":"22","name":"daerwen"}
	var info = make(map[string]string)
	json.Unmarshal(data2, &info)
	//map[age:22 name:daerwen] map[string]string
	fmt.Println(info, reflect.TypeOf(info))

}
