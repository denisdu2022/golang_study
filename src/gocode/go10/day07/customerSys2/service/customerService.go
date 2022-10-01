package service

import (
	"customerSys2/model"
	"fmt"
	"strings"
)

// 定义客户信息管理系统数据结构
type CustomerSys struct {
	customers  []model.Customer
	customerId int
}

// 返回上一层函数功能
func isBack() bool {
	var back string
	fmt.Print("请问是否返回上一层[Y/N]: ")
	fmt.Scan(&back)
	if strings.ToUpper(back) == "Y" {
		return true
	} else {
		return false
	}
}

// 添加客户功能函数
func (cs *CustomerSys) AddCustomer() {
	for true {

		//调用返回上一层功能函数
		b := isBack()
		if b {
			break
		}
	}

}

// 查询客户信息功能函数
func (cs *CustomerSys) ListCustomer() {
	for true {

		//调用返回上一层功能函数
		b := isBack()
		if b {
			break
		}
	}

}

// 更新客户信息功能函数
func (cs *CustomerSys) UpdataCustomer() {
	for true {

		//调用返回上一层功能函数
		b := isBack()
		if b {
			break
		}
	}

}

// 删除客户信息功能函数
func (cs *CustomerSys) DelteCustomer() {
	for true {

		//调用返回上一层功能函数
		b := isBack()
		if b {
			break
		}
	}

}

// 保存客户信息功能函数
func (cs *CustomerSys) SaveCustomer() {

}
