package service

import (
	"customerSys2/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

//客户信息管理系统service结构体
type CustomerSys struct {
	customer   []model.Customer
	customerId int
}

//返回上一层函数
func isBack() bool {
	var back string
	fmt.Print("请问是否返回上一层[Y/N]: ")
	fmt.Scanln(&back)
	if strings.ToUpper(back) == "Y" {
		return true
	} else {
		return false
	}

}

//客户信息管理系统[]customer索引函数
func (cs *CustomerSys) CustomerIndex(cid int) int {
	var indexCustome = -1
	for handleIndex, customerValue := range cs.customer {
		if customerValue.Cid == cid {
			indexCustome = handleIndex
		}
	}
	return indexCustome
}

//初始化客户信息数据函数
func (cs *CustomerSys) InitCusomers() {
	fmt.Println("-------------------初始化客户信息开始-------------------")
	customerJsonByts, err := ioutil.ReadFile("db/cusomerJsonData.json")
	if err != nil {
		fmt.Println("初始化客户信息数据失败,错误信息是: ", err)
	} else {
		fmt.Println("初始化客户信息数据成功***")
	}
	//反序列化
	newInitCustomer := cs.customer
	json.Unmarshal(customerJsonByts, &newInitCustomer)
	fmt.Println("-------------------初始化客户信息结束-------------------")

}

//添加客户信息功能函数
func (cs *CustomerSys) AddCustome() {
	for true {
		fmt.Println("-------------------添加客户信息开始-------------------")
		//引导用户输入客户信息
		name, gender, age, email := model.InputCustomer()
		//每输入一个客户,客户信息编号自增1
		cs.customerId++
		//将输入的客户信息赋值给model.customer的每个成员变量
		newCustomer := model.Customer{
			Cid:    cs.customerId,
			Name:   name,
			Gender: gender,
			Age:    age,
			Email:  email,
		}
		//将新添加的客户信息,扩容添加到新的客户信息切片中
		cs.customer = append(cs.customer, newCustomer)
		fmt.Println("-------------------添加客户信息成功-------------------")
		fmt.Println("-------------------添加客户信息结束-------------------")
		//判断是否返回上一层
		b := isBack()
		if b {
			break
		}
	}

}

//查看客户信息功能函数
func (cs *CustomerSys) ListCustome() {
	for true {
		fmt.Println("-------------------查询客户信息开始-------------------")
		//for循环遍历每一个customerValue
		for _, customerValue := range cs.customer {
			fmt.Printf("客户编号:%-6d 姓名:%-6s 性别:%-6s 年龄:%-6d 邮箱:%-6s\n", customerValue.Cid,
				customerValue.Name, customerValue.Gender, customerValue.Age, customerValue.Email)
		}
		fmt.Println("-------------------查询客户信息成功-------------------")
		fmt.Println("-------------------查询客户信息结束-------------------")
		//判断是否返回上一层
		b := isBack()
		if b {
			break
		}
	}

}

//更新客户信息功能函数
func (cs *CustomerSys) UpdataCustome() {
	for true {
		//判断是否返回上一层
		b := isBack()
		if b {
			break
		}
	}

}

//删除客户信息功能函数
func (cs *CustomerSys) DelCustome() {
	for true {
		var cid int
		fmt.Println("-------------------删除客户信息开始-------------------")
		//引导用户输入要删除的客户信息编号
		fmt.Print("请输入您要删除的客户信息编号: ")
		fmt.Scanln(&cid)
		//传入要删除的客户信息编号,并接收返回的[]modle.customer的索引值
		delCustomerIndex := cs.CustomerIndex(cid)
		//判断索引是否存在,如果不存在则需要继续输入
		if delCustomerIndex == -1 {
			fmt.Println("您输入的客户信息编号不存在,请您重新输入...")
			continue
		}
		cs.customer = append(cs.customer[:delCustomerIndex], cs.customer[delCustomerIndex+1:]...)
		fmt.Println("-------------------删除客户信息成功-------------------")
		fmt.Println("-------------------删除客户信息结束-------------------")
		//判断是否返回上一层
		b := isBack()
		if b {
			break
		}
	}

}

//保存客户信息功能函数
func (cs *CustomerSys) SaveCustome() {
	fmt.Println("-------------------保存客户信息开始-------------------")
	//定义文件保存接收信息和err
	customerData, err := json.Marshal(cs.customer)
	//判断文件是否保存成功
	if err != nil {
		fmt.Println("客户信息保存失败,失败原因是: ", err)
	} else {
		fmt.Println("文件保存成功***")
	}
	//将json数据写入文件
	ioutil.WriteFile("db/customerJsonData.json", customerData, 0666)
	fmt.Println("-------------------保存客户信息结束-------------------")
}
