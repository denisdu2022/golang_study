package validator

import (
	"bingotest01/application/model"
	"bingotest01/application/utils"
	"errors"
	"github.com/go-playground/validator/v10"
)

//登录验证器

func UserValidator(user *model.User) error {
	//获取验证器对象
	validate, trans := utils.GenValidate()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}
