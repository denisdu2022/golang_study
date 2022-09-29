package model

//定义客户信息管理系统数据结构

type Customer struct {
	Cid    int    `json:"cid"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int8   `json:"age"`
	Email  string `json:"email"`
}
