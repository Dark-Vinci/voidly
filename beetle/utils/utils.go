package utils

import (
	"context"
	"fmt"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
	"github.com/gin-gonic/gin"
)

const packageName = "packageName"

func GetRequestID(c context.Context) string {
	fmt.Println(c.Value("X-Request-ID"))
	if val, ok := c.Value(RequestIDKey).(string); ok {
		return val
	}

	return ""
}

func GetFromContext[T any](c context.Context, key string) T {
	if val, ok := c.Value(key).(T); ok {
		return val
	}

	fmt.Println("--------------zero context---------------------")

	var zero T

	return zero
}

func GetContext(c *gin.Context) models.CTX {
	if val, ok := c.Get(CTX); ok {
		ctx := val.(models.CTX)
		return ctx
	}

	return models.CTX{}
}
