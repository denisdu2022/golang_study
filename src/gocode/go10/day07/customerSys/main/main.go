package main

import (
	"customerSys/model"
	"customerSys/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	//初始化保存数据
	//读取保存数据的文件
	customerJsonBytes, err := ioutil.ReadFile("db/customerData.json")
	if err != nil {
		fmt.Println("初始化客户信息失败,错误原因: ", err)
	} else {
		fmt.Println("初始化客户信息成功........")
	}
	//定义反序列化接收变量
	var customers []model.Customer
	//json反序列化
	json.Unmarshal(customerJsonBytes, &customers)

	//定义初始化自增ID initCustomerID为0
	var initCustomerID = 0
	//当customers的长度不等于0的时候让初始的自增IDinitCustomerID == customers的索引是customers的长度-1.的键Cid
	if len(customers) != 0 {
		initCustomerID = customers[len(customers)-1].Cid
	}

	//实例化CustomerService
	cs := service.CustomersSys{
		Customer:   customers,
		CustomerID: initCustomerID,
	}

	for true {
		service.Wel()
		cs.Choice()
	}

}
