package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/app"
	"github.com/dark-vinci/stripchat/beetle/middleware"
	"github.com/dark-vinci/stripchat/beetle/utils"
)

const packageName = "res.auth"

type authApi struct {
	m *middleware.Middleware
	a app.Operation
	e *utils.Environment
	z *zerolog.Logger
}

func New(eng *gin.RouterGroup, ap app.Operation, m *middleware.Middleware, e *utils.Environment, z *zerolog.Logger) {
	logger := z.With().Str(utils.PackageStrHelper, packageName).Logger()

	a := authApi{
		m: m,
		a: ap,
		e: e,
		z: &logger,
	}

	authGroup := eng.Group("/auth", m.ZeroUserContext())

	authGroup.POST("/login", a.login())
	authGroup.POST("/create", a.create())
}
