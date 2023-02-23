package model

//用户信息表

type UserInfo struct {
	//ID 主键
	ID int `gorm:"primaryKey"`
	//用户名
	Account string
	//密码
	Pwd string
}
