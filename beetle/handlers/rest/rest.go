package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/app"
	"github.com/dark-vinci/stripchat/beetle/handlers/rest/auth"
	"github.com/dark-vinci/stripchat/beetle/handlers/rest/chats"
	"github.com/dark-vinci/stripchat/beetle/handlers/rest/message"
	"github.com/dark-vinci/stripchat/beetle/middleware"
	"github.com/dark-vinci/stripchat/beetle/utils"
)

func Build(g *gin.RouterGroup, a app.Operation, l *zerolog.Logger, e *utils.Environment) {
	g.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":   "stripchat",
			"status": 200,
		})
	})

	m := middleware.New(l, e, a)

	auth.New(g, a, m, e, l)
	chats.New(g, a, m, e, l)
	message.New(g, a, m, e, l)
}
