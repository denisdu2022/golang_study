package main

import "fmt"

func main() {

	//客户信息管理系统的数据结构

	//01 存储客户信息的数据类型
	//var customer1 = map[string]string{}  //每一个用户的数据类型,map类型,键是字符串类型,值是字符串类型, 每一个用户可以有多个键值属性
	//var customers = []map[string]string{}  //客户信息数据的总结构,[]map[string]string 每一个客户都有一个切片的索引,方便用户信息
	//的管理,对用户的增删改查

	var customer1 = map[string]string{"name": "daerwen", "age": "22"}
	var customer2 = map[string]string{"name": "dafenqi", "age": "18"}
	var customer3 = map[string]string{"name": "bijiasuo", "age": "23"}

	var customers = []map[string]string{customer1, customer2, customer3}
	fmt.Println(customers)
	//[map[age:22 name:daerwen] map[age:18 name:dafenqi] map[age:23 name:bijiasuo]]

	//02 添加客户信息
	var newCustomer = map[string]string{"name": "hezi", "age": "26"}
	//添加客户信息到总的客户信息数据中
	customers = append(customers, newCustomer)
	fmt.Println(customers)
	//[map[age:22 name:daerwen] map[age:18 name:dafenqi] map[age:23 name:bijiasuo] map[age:26 name:hezi]]

	//03 查看客户信息
	//客户信息总的数据类型是[]map[string][string] 使用range遍历来取键和值
	for indexCustomer, valueCustomer := range customers {
		//fmt.Printf("第%-6d客户是: %-6s\n", indexCustomer, valueCustomer)
		fmt.Printf("第%-1d客户姓名是: %-6s  年龄是:%-6s\n", indexCustomer, valueCustomer["name"], valueCustomer["age"])
	}

	//03 修改客户信息,按照切片索引修改客户信息
	//要修改第三个客户的姓名为大写,年龄为21
	//newCustomer2 := customers[2]
	//newCustomer2["name"] = "BIJIASUO"
	//newCustomer2["age"] = "21"
	//fmt.Println(newCustomer2)
	//或者也可以
	customers[2]["name"] = "BIJIASUO"
	customers[2]["age"] = "21"
	fmt.Println(customers[2])
	//map[age:21 name:BIJIASUO]

	//04 删除客户信息 按照切片的索引删除客户信息
	customers = append(customers[:3], customers[3+1:]...)
	for indexCustomer, valueCustomer := range customers {
		//fmt.Printf("第%-6d客户是: %-6s\n", indexCustomer, valueCustomer)
		fmt.Printf("第%-1d客户姓名是: %-6s  年龄是:%-6s\n", indexCustomer, valueCustomer["name"], valueCustomer["age"])
	}
	/*输出结果
	第0客户姓名是: daerwen  年龄是:22
	第1客户姓名是: dafenqi  年龄是:18
	第2客户姓名是: BIJIASUO  年龄是:21
	*/
}
