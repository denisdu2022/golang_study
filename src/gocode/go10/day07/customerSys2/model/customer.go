package model

// 定义客户信息管理系统源数据结构体
type Customer struct {
	Cid    int    `json:"cid"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
}
