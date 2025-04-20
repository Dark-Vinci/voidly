package chats

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/app"
	"github.com/dark-vinci/stripchat/beetle/middleware"
	"github.com/dark-vinci/stripchat/beetle/utils"
)

const packageName = "rest.chats"

type chatsApi struct {
	m *middleware.Middleware
	a app.Operation
	e *utils.Environment
	z *zerolog.Logger
}

func New(eng *gin.RouterGroup, ap app.Operation, m *middleware.Middleware, e *utils.Environment, z *zerolog.Logger) {
	logger := z.With().Str(utils.PackageStrHelper, packageName).Logger()

	a := chatsApi{
		a: ap,
		m: m,
		e: e,
		z: &logger,
	}

	chatGroup := eng.Group("/chats", a.m.Authenticate())

	chatGroup.POST("/", a.createChat())
	chatGroup.GET("/:user_id", a.getChatList())
}
