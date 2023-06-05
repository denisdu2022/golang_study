package model

import (
	"bingotest01/application/database"
	"github.com/jinzhu/gorm"
)

//用户实体

type User struct {
	//继承gorm.Model
	gorm.Model
	//用户名
	Username string `gorm:"unique_index;size:255;comment:'账户名'" json:"username" sql:"index"`
	//哈希密码
	HashPassword string `gorm:"not null;size:255;comment:'密码'" json:"_"`
	//密码
	Password string `gorm:"_" json:"password,omitempty"`
	//昵称
	Nickname string `gorm:"size:255;comment:'昵称'" json:"nickname" sql:"index"`
	//手机号
	Mobile string `gorm:"index;size:15;comment:'手机';" json:"mobile"`
	//邮箱
	Email string `valid:"email" gorm:"index;size:255;comment:'邮箱';" json:"email"`
	//头像
	Avatar string `gorm:"size:255;comment:'头像'" json:"avatar"`
	//性别
	Sex bool `gorm:"type:boolean;default:true;comment:'性别'" json:"sex"`
	//IP
	Ip string `valid:"ip" gorm:"size:255;comment:'IP地址';" json:"ip"`
	//状态
	Status bool `gorm:"type:boolean;default:true;comment:'状态'" json:"status"`
}

//设置表名

func (User) TableName() string {
	return "user_info"
}

//创建用户

func (user User) Insert() (id uint, err error) {
	//添加用户
	result := database.Orm.Create(&user)

	return user.ID, result.Error
}

//根据用户ID获取单个用户

func (user *User) GetOneById(id uint) {
	database.Orm.First(&user, id)
}

//根据账户信息(用户名,手机,邮箱获取用户)

func (user *User) GetOneByAccount(account string) {
	database.Orm.First(&user, "username = ? or mobile = ? or email = ?", account, account, account)
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
