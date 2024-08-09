package utils

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	. "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

//生成验证器

func GenValidate() (*validator.Validate, ut.Translator) {
	//获取翻译对象
	zhCh := zh.New()
	//获取validator对象
	validate := validator.New()

	//注册一个函数,获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	//翻译
	UniversalTranslator := ut.New(zhCh)
	trans, _ := UniversalTranslator.GetTranslator("zh")
	_ = RegisterDefaultTranslations(validate, trans)
	return validate, trans
}
