package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/pkg/response"
)

func BindRequest(c *gin.Context, reqBody interface{}) bool {
	if err := c.ShouldBindBodyWithJSON(&reqBody); err != nil {
		if err.Error() == "EOF" {
			response.BadRequestError(c, response.ErrCodeInvalidRequest, "No data provided")
			return false
		}
		response.BadRequestErrorWithFields(c, response.ErrCodeInvalidRequest, GetObjMessage(err))
		return false
	}

	return true
}
