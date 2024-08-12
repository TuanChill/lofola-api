package helpers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/internal/models"
	"github.com/tuanchill/lofola-api/pkg/response"
	"github.com/tuanchill/lofola-api/pkg/utils"
)

// check identifier is email or username
func CheckIdentifier(identifier string) string {
	if strings.Contains(identifier, "@") {
		return "email"
	}
	return "username"
}

// set header and value of this for response
func SetHeaderResponse(w http.ResponseWriter, header string, value string) {
	w.Header().Set(header, value)
}

// parse time string to time.Time
func ParseTime(timeStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, timeStr)
}

// validate request body
func ValidateRequestBody(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindBodyWithJSON(data); err != nil {
		if err.Error() == "EOF" {
			response.BadRequestError(c, response.ErrCodeInvalidRequest, "No data provided")
			return err
		}

		if len(utils.GetObjMessage(err)) == 0 {
			response.BadRequestError(c, response.ErrCodeInvalidRequest, err.Error())
			return err
		}
		response.BadRequestErrorWithFields(c, response.ErrCodeInvalidRequest, utils.GetObjMessage(err))
		return err
	}

	return nil
}

// ValidateRequestSearch is a function that validate request search param
func ValidateRequestSearch(c *gin.Context, data *models.SearchParam) error {
	if err := c.ShouldBindQuery(data); err != nil {
		if err == err.(*strconv.NumError) {
			response.BadRequestError(c, response.ErrCodeInvalidRequest, "Page and Limit must be a number")
			return err
		}
		response.BadRequestErrorWithFields(c, response.ErrCodeInvalidRequest, utils.GetObjMessage(err))
		return err
	}

	return nil
}
