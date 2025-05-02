package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ErrorData struct {
	ID      uuid.UUID `json:"id"`
	Details string    `json:"details"`
	Status  int       `json:"status"`
	Handler string    `json:"handler" omitempty`
}

func getStringPointer(val string) *string {
	return &val
}

type GenericResponse[T any] struct {
	Code    int
	Message *string
	Data    T
	Error   *ErrorData
}

func Build[T any](code int, data T, message *string, error *ErrorData) GenericResponse[T] {
	return GenericResponse[T]{
		Code:    code,
		Message: message,
		Data:    data,
		Error:   error,
	}
}

// ErrorResponse template for error responses
func ErrorResponse(c *gin.Context, code int, error ErrorData) {
	c.JSON(code, Build[*string](
		code,
		nil,
		getStringPointer("error has occurred"),
		&error))
	c.Abort()
}

// OkResponse template for ok and successful responses
func OkResponse[T any](c *gin.Context, code int, message string, data T) {
	c.JSON(code, Build(
		code,
		data,
		getStringPointer(message),
		nil))
	c.Abort()
}
