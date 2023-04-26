package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	//1.先读取json文件
	data, _ := ioutil.ReadFile("data.json")
	fmt.Println(string(data))

	//2.声明user对象
	var user User

	//3.反序列化把读取到的文件写入struct
	json.Unmarshal(data, &user)

	//4.操作user对象
	fmt.Println(user)
	fmt.Println("姓名: ", user.Name)
	fmt.Println("年龄: ", user.Age)

}
