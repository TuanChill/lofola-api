package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

// InitValidator initializes the custom validator
func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("birthday", birthdayValidator)
		v.RegisterValidation("isNumber", isNumberValidator)
	}
}

var birthdayValidator validator.Func = func(fl validator.FieldLevel) bool {
	if birthday, ok := fl.Field().Interface().(string); ok {
		return utils.IsBirthDayValid(birthday)
	}

	return false
}

var isNumberValidator validator.Func = func(fl validator.FieldLevel) bool {
	fmt.Println(fl.Field().Interface())
	if number, ok := fl.Field().Interface().(int); ok {
		return utils.IsNumberValid(string(number))
	}

	return false
}
