package services

import (
	"bingo_api/applcation/model"
	"bingo_api/applcation/utils"
	"github.com/gin-gonic/gin"
)

//创建用户

func CreateUser(ctx *gin.Context) (uint, error) {
	//实例化User对象
	user := model.User{}
	//定义error对象
	var err error
	//接收前端发来的数据
	if err = ctx.ShouldBindJSON(&user); err != nil {
		return 0, err
	}

	user.HashPassword, err = utils.MakeHashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	return user.Insert()
}

//获取指定ID的用户

func GetUserById(id uint) (user model.User) {
	//实例化user对象
	user = model.User{}
	//调用指定ID获取用户方法
	user.GetOneByID(id)
	//返回user对象
	return user
}

//根据账户信息(用户名,手机,邮箱)获取账户

func GetUserByAccount(account string) (user model.User) {
	//实例化user对象
	user = model.User{}
	//调用user方法
	user.GetOneByAccount(account)
	//返回user对象
	return user
}

//获取所有用户

func GetAllUser() []model.User {
	//实例化user对象
	user := model.User{}
	return user.GetAll()
}

//更新用户密码

func ChangeUserPassword(user model.User, RawPassword string) {
	password, _ := utils.MakeHashPassword(RawPassword)
	user.ChangePassword(password)
}
