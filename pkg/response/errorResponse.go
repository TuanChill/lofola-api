package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents the error response
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Now     int64  `json:"now"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// NewErrorResponse creates a new ErrorResponse instance
func NewErrorResponse(message string, status int, code int) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Status:  status,
		Code:    code,
		Now:     time.Now().Unix(),
	}
}

// Send ErrorResponse to client
// it aborts the request and responds with the error message as JSON
func (e *ErrorResponse) Send(c *gin.Context) {
	c.AbortWithStatusJSON(e.Status, e)
}

// BadRequestError sends a 400 Bad Request error response
func BadRequestError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(code)
	}
	NewErrorResponse(message, 400, code).Send(c)
}

func BadRequestErrorWithFields(c *gin.Context, code int, fields []ErrorMsg) {
	c.AbortWithStatusJSON(400, gin.H{
		"code":    code,
		"message": GetReasonPhrase(code),
		"status":  400,
		"errors":  fields,
		"now":     time.Now().Unix(),
	})
}

// NotFoundError sends a 404 Not Found error response
func NotFoundError(c *gin.Context, code int, message ...string) {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}

	if msg == "" {
		msg = GetReasonPhrase(code)
	}
	NewErrorResponse(msg, 404, code).Send(c)
}

// ConflictError sends a 409 Conflict error response
func ConflictError(c *gin.Context, code int, message ...string) {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}

	if msg == "" {
		msg = GetReasonPhrase(code)
	}
	NewErrorResponse(msg, 409, code).Send(c)
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(c *gin.Context, code int, message ...string) {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}

	if msg == "" {
		msg = GetReasonPhrase(code)
	}
	NewErrorResponse(msg, 500, code).Send(c)
}

// UnauthorizedError sends a 401 Unauthorized error response
func UnauthorizedError(c *gin.Context, code int, message ...string) {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}

	if msg == "" {
		msg = GetReasonPhrase(code)
	}
	NewErrorResponse(msg, 401, code).Send(c)
}

// ForbiddenError sends a 403 Forbidden error response
func ForbiddenError(c *gin.Context, code int, message ...string) {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}

	if msg == "" {
		msg = GetReasonPhrase(code)
	}
	NewErrorResponse(msg, 403, code).Send(c)
}

// ServiceUnavaiable sends a 503 Service Unavailable error response
func ServiceUnavaiable(c *gin.Context, code int, message ...string) {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}

	if msg == "" {
		msg = GetReasonPhrase(code)
	}
	NewErrorResponse(msg, 503, code).Send(c)
}
