package helpers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
