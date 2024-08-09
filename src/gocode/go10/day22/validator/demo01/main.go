package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// User 结构体
type User struct {
	Name string `validate:"required=1,error=姓名不能为空 (中文翻译)"`
	Age  int    `validate:"required=1,error=年龄不能为空 (中文翻译),gte=18,lte=55,error=年龄必须在18-55之间 (中文翻译)"`
}

func main() {

	//实例化结构体对象
	user := &User{
		Name: "",
		Age:  16,
	}

	//实例化中文翻译器
	zhT := zh.New()
	uni := ut.New(zhT, zhT)
	trans, _ := uni.GetTranslator("zh")

	//实例化验证器并注册中文翻译器
	validate := validator.New()
	zh.RegisterDefaultTranslations(validate, trans)

	//验证结构体对象
	err := validate.Struct(user)
	if err != nil {
		//输出中文错误信息
		for _, value := range err.(validator.ValidationErrors) {
			fmt.Println("value>>> ", value)
		}
	}

}
