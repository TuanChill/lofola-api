package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/tuanchill/lofola-api/pkg/response"
)

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "email":
		return "Invalid email format"
	case "min":
		return "Minimum length is " + fe.Param()
	case "max":
		return "Maximum length is " + fe.Param()
	case "alphanum":
		return "This field must be alphanumeric"
	case "numeric":
		return "This field must be numeric"
	case "len":
		return "This field must have length " + fe.Param()
	case "birthday":
		return "Invalid birthdate format, must be YYYY-MM-DD"
	}

	return "Unknown error"
}

func GetObjMessage(err error) []response.ErrorMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]response.ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = response.ErrorMsg{Field: fe.Field(), Message: getErrorMsg(fe)}
		}

		return out
	}
	return nil
}
