package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
)

func (m *Middleware) ZeroUserContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Get(utils.RequestIDKey)

		requestID, err := uuid.Parse(id.(string))

		if err != nil {
			requestID = uuid.New()
		}

		ctx := models.CTX{
			Context:   c.Request.Context(),
			UserAgent: "CHROME",
			UserID:    uuid.Nil,
			RequestID: requestID,
			Email:     "DUMMY@GMAIL.COM",
		}

		c.Set(utils.CTX, ctx)
		c.Next()
	}
}
