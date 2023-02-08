package model

//登录用户信息模型

type UserInfo struct {
	ID       int `gorm:"primaryKey"`
	Account  string
	Password string
}
