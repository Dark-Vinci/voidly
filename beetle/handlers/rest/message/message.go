package message

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/app"
	"github.com/dark-vinci/stripchat/beetle/middleware"
	"github.com/dark-vinci/stripchat/beetle/utils"
)

const packageName = "rest.message"

type messageApi struct {
	m *middleware.Middleware
	a app.Operation
	e *utils.Environment
	z *zerolog.Logger
}

func New(eng *gin.RouterGroup, ap app.Operation, m *middleware.Middleware, e *utils.Environment, z *zerolog.Logger) {
	logger := z.With().Str(utils.PackageStrHelper, packageName).Logger()

	mApi := messageApi{
		m: m,
		a: ap,
		e: e,
		z: &logger,
	}

	messageGroup := eng.Group("/message", mApi.m.Authenticate())

	messageGroup.POST("/", mApi.create())
	messageGroup.GET("/:chat_id", mApi.list())
}
