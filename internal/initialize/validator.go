package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

// InitValidator initializes the custom validator
func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("birthday", birthdayValidator)
	}
}

var birthdayValidator validator.Func = func(fl validator.FieldLevel) bool {
	if birthday, ok := fl.Field().Interface().(string); ok {
		return utils.IsBirthDayValid(birthday)
	}

	return false
}
