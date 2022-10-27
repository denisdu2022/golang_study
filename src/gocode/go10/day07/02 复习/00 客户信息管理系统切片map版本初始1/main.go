package main

import "fmt"

//客户信息可以放在map里面
//map[string]string

func main() {

	//可以存储一个客户的信息
	var customer1 = map[string]string{"name": "daerwen", "age": "22"}
	var customer2 = map[string]string{"name": "dafenqi", "age": "21"}
	var customer3 = map[string]string{"name": "bijiasuo", "age": "18"}

	//存储多个用户的信息
	var customers = []map[string]string{customer1, customer2, customer3}
	fmt.Println(len(customers), cap(customers)) //3 3

	//添加一个客户
	var newCustomer = map[string]string{"name": "hezi", "age": "21"}
	//存储多个用户的信息的customers的类型是[]map[string][string] ,所以就是对切片的扩容
	customers = append(customers, newCustomer)
	fmt.Println(customers, len(customers), cap(customers))
	//[map[age:22 name:daerwen] map[age:21 name:dafenqi] map[age:18 name:bijiasuo] map[age:21 name:hezi]] 4 6

	//查看用户信息
	//存储多个用户的信息的customers的类型是[]map[string][string] ,所以就是对切片的遍历
	for indexCustomer, valueCustomer := range customers {
		fmt.Printf("第%-1d客户的姓名是:%-6s  年龄是:%-6s\n", indexCustomer, valueCustomer["name"], valueCustomer["age"])
		//valueCustomer["age"] = "0" //把每个客户的年龄修改为0
	}
	//验证年龄修改为0,取第二个客户的年龄
	//是因为map是引用类型,所以都会被修改为0
	//fmt.Println(customers[2]["age"])

	/*输出结果
	第0客户的姓名是:daerwen  年龄是:22
	第1客户的姓名是:dafenqi  年龄是:21
	第2客户的姓名是:bijiasuo  年龄是:18
	第3客户的姓名是:hezi    年龄是:21
	*/

	//修改用户信息  按照切片的索引进行修改
	//修改第三个客户的姓名为大写,年龄为23
	customers[2]["name"] = "BIJIASUO"
	customers[2]["age"] = "23"
	fmt.Printf("第三个客户的,姓名%-8s  年龄:%-8s \n", customers[2]["name"], customers[2]["age"])
	//第三个客户的,姓名BIJIASUO  年龄:23

	//or
	//customerMod := customers[2]
	//customerMod["name"] = "BIJIASUO"
	//customerMod["age"] = "23"
	//fmt.Printf("第三个客户的,姓名%-8s  年龄:%-8s \n", customers[2]["name"], customers[2]["age"])

	//删除用户信息,按照索引进行删除
	a := []int{1, 2, 3, 4, 5, 6} //一个整形的的切片
	//要删除第三个,append的第一的元素是从0开始,到2是第三个,所以是a[:2] 取到第二个,左闭右开,3是取不到的
	//然后append第二个元素,是2+1追加到前边的第一个元素 ,加...是相当于打撒后追加4,5,6追加后就是1,2,4,5,6
	a = append(a[:2], a[2+1:]...)
	fmt.Println(a)

	//删除第四个客户的信息
	fmt.Println(customers[3]) //map[age:21 name:hezi]

	customers = append(customers[:3], customers[3+1:]...)
	for indexCustomer, valueCustomer := range customers {
		fmt.Printf("第%-1d客户的姓名是:%-6s  年龄是:%-6s\n", indexCustomer, valueCustomer["name"], valueCustomer["age"])
		//valueCustomer["age"] = "0" //把每个客户的年龄修改为0
	}
	/*输出结果
	第0客户的姓名是:daerwen  年龄是:22
	第1客户的姓名是:dafenqi  年龄是:21
	第2客户的姓名是:BIJIASUO  年龄是:23
	*/
}
