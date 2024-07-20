package response

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message          string      `json:"message"`
	Data             interface{} `json:"data"`
	StatusCode       int         `json:"statusCode"`
	ReasonStatusCode int         `json:"reasonStatusCode"`
	MetaData         interface{} `json:"metaData,omitempty"`
	Options          interface{} `json:"options,omitempty"`
}

func NewSuccessResponse(message string, data interface{}, statusCode int, reasonStatusCode int, metaData interface{}, options interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message:          message,
		Data:             data,
		StatusCode:       statusCode,
		ReasonStatusCode: reasonStatusCode,
		MetaData:         metaData,
		Options:          options,
	}
}

func (sr *SuccessResponse) Send(c *gin.Context) {
	c.JSON(sr.StatusCode, sr)
}

func Ok(c *gin.Context, message string, metadata any) {
	if message == "" {
		message = GetReasonPhrase(StatusOK)
	}

	response := NewSuccessResponse(message, nil, StatusOK, StatusOK, metadata, nil)

	response.Send(c)
}

func OkWithData(c *gin.Context, message string, data any, metadata any) {
	if message == "" {
		message = GetReasonPhrase(StatusOK)
	}

	response := NewSuccessResponse(message, data, StatusOK, StatusOK, metadata, nil)

	response.Send(c)
}

func Created(c *gin.Context, message string, data any, metadata any) {
	if message == "" {
		message = GetReasonPhrase(StatusCreated)
	}

	response := NewSuccessResponse(message, data, StatusCreated, StatusCreated, metadata, nil)

	response.Send(c)
}

func ListDataResponse(c *gin.Context, message string, data any, metadata any) {
	if message == "" {
		message = GetReasonPhrase(StatusOK)
	}

	response := NewSuccessResponse(message, data, StatusOK, StatusOK, metadata, nil)

	response.Send(c)
}
