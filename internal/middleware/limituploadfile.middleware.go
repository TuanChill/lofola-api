package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuanchill/lofola-api/configs/common/constants"
	"github.com/tuanchill/lofola-api/pkg/response"
)

func LimitUploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20)
		// ParseMultipartForm parses a request body as multipart/form-data.
		// The whole request body is parsed and up to a total of maxMemory bytes of its file parts are stored in memory,
		// check form is multipart
		if c.Request.Header.Get("Content-Type") != "multipart/form-data" {
			c.Next()
			return
		}

		if err := c.Request.ParseMultipartForm(constants.MAX_UPLOAD_SIZE); err != nil {
			response.BadRequestError(c, response.ErrCodeInvalidRequest, "FILE_TOO_LARGE")
			c.Abort()
			return
		}
	}
}
