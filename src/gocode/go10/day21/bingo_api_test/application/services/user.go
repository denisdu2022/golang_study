package services

import (
	"bingotest01/application/constants"
	. "bingotest01/application/model"
	"bingotest01/application/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

//用户登录

func UserLogin(ctx *gin.Context) (user User, err error) {
	//实例化user对象
	user = User{}
	//获取user
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return user, err
	}
	//数据库查询用户名
	user.GetOneByAccount(user.Username)
	//判断
	if user.ID < 1 {
		err = errors.New(constants.NoSuchUser)
		//记录日志
		zap.S().Error(constants.NoSuchUser)
		return
	}
	//判断:如果用户密码错误,则给用户返回密码错误
	checkPass := utils.CheckPassword(user.HashPassword, user.Password)
	if !checkPass {
		err = errors.New(constants.PasswordError)
		zap.S().Error(constants.PasswordError)
		return
	}
	return
}
