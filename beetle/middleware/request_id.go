package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "X-Request-ID"

func (m *Middleware) RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()

		c.Set(RequestIDKey, reqID)
		c.Writer.Header().Set(RequestIDKey, reqID)

		c.Next()
	}
}
