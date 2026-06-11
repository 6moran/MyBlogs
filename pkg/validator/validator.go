package validator

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9]{6,20}$`)

// GetValidMsg 得到校验失败后err的具体中文错误
func GetValidMsg(err error) string {
	//将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		return validMsg(errs[0])
	}
	return "输入参数有误"
}

// 枚举所有中文错误
func validMsg(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + "是必填项"
	case "min":
		return e.Field() + "长度不能小于" + e.Param()
	case "max":
		return e.Field() + "长度不能大于" + e.Param()
	case "len":
		return e.Field() + "长度必须是" + e.Param()
	case "alphanum":
		return e.Field() + "只能包含字母和数字"
	case "containsany":
		return e.Field() + "必须包含以下字符之一: " + e.Param()
	case "gte":
		return e.Field() + "不能小于" + e.Param()
	case "lte":
		return e.Field() + "不能大于" + e.Param()
	case "email":
		return "邮箱格式不正确"
	case "contact_email":
		return "联系我们邮箱格式不正确"
	case "oneof":
		return e.Field() + "必须是" + e.Param() + "中的一个"
	case "numeric":
		return e.Field() + "必须是数字"
	case "url":
		return e.Field() + "必须是合法的URL地址"
	case "password":
		return "密码格式不正确，必须是6-20位字母和数字的组合"
	default:
		return e.Field() + "验证失败"
	}
}

// InitCustomVali 初始化自定义校验器
func InitCustomVali() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("password", passwordValid); err != nil {
			return fmt.Errorf("注册password校验器失败: %w", err)
		}
	}
	return nil
}

// 自定义密码校验器
func passwordValid(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if !passwordRegex.MatchString(password) {
		return false
	}
	hasLetter, _ := regexp.MatchString("[a-zA-Z]", password)
	hasDigit, _ := regexp.MatchString("[0-9]", password)
	return hasLetter && hasDigit
}
