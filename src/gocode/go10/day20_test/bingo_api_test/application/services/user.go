package services

import (
	. "bingotest01/application/model"
	"bingotest01/application/utils"
	"github.com/gin-gonic/gin"
)

//创建用户

func CreateUser(ctx *gin.Context) (uint, error) {
	//实例化USER对象
	user := User{}
	var err error
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return 0, err
	}
	//验证密码
	user.HashPassword, err = utils.MakeHashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	return user.Insert()
}

//获取指定ID的用户

func GetUserById(id uint) (user User) {
	user = User{}
	user.GetOneById(id)
	return user
}

//根据账户信息(用户名,手机号,邮箱)获取用户

func GetUserByAccount(account string) (user User) {
	user = User{}
	user.GetOneByAccount(account)
	return user
}

//获取所有用户

func GetAllUser() []User {
	user := User{}
	return user.GetAll()
}

//更新密码

func ChangeUserPassword(user User, RawPassword string) {
	password, _ := utils.MakeHashPassword(RawPassword)
	user.ChangePassword(password)
}
