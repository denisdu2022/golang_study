package model

import (
	"bingo_api/applcation/database"
	"github.com/jinzhu/gorm"
)

//用户实体

type User struct {
	//继承gorm.Model
	gorm.Model
	Username     string `gorm:"unique_index;size:255;comment:'账户名'" json:"username" sql:"index"`
	HashPassword string `gorm:"not null;size:255;comment:'密码'" json:"_"`
	Password     string `gorm:"_" json:"password,omitempty"`
	Nickname     string `gorm:"size:255;comment:'昵称'" json:"nickname" sql:"index"`
	Mobile       string `gorm:"index;size:15;comment:'手机';" json:"mobile"`
	Email        string `valid:"email" gorm:"index;size:255;comment:'邮箱';" json:"email"`
	Avatar       string `gorm:"size:255;comment:'头像'" json:"avatar"`
	Sex          bool   `gorm:"type:boolean;default:true;comment:'性别'" json:"sex"`
	IP           string `valid:"ip" gorm:"size:255;comment:'IP地址';" json:"IP"`
	Status       bool   `gorm:"type:boolean;default:true;comment:'状态'" json:"status"`
}

//设置表名

func (User) TableName() string {
	return "user_info"
}

//创建用户

func (user User) Insert() (id uint, err error) {
	//添加数据
	result := database.Orm.Create(&user)
	return user.ID, result.Error
}

//根据指定ID获取用户

func (user *User) GetOneByID(id uint) {
	database.Orm.First(&user, id)
}

//根据账户信息(用户名,手机,邮箱) 获取用户

func (user *User) GetOneByAccount(account string) {
	database.Orm.First(&user, "username =? or mobile = ? or email = ?", account, account, account)
}

//获取所有用户

func (user User) GetAll() []User {
	var users []User
	database.Orm.Find(&users)
	return users
}

//更新密码

func (user User) ChangePassword(HashPassword string) {
	user.HashPassword = HashPassword
	database.Orm.Save(&user)
}
