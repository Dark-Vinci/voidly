package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dark-vinci/stripchat/beetle/utils"
)

func (m *Middleware) RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()

		c.Set(utils.RequestIDKey, reqID)
		c.Writer.Header().Set(utils.RequestIDKey, reqID)

		c.Next()
	}
}
