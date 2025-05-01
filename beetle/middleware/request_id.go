package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "X-Request-ID"

func (m *Middleware) RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()

		fmt.Println(reqID)

		// Set the request ID in context and header
		c.Set(RequestIDKey, reqID)
		c.Writer.Header().Set(RequestIDKey, reqID)

		c.Next()
	}
}
