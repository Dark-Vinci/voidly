package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/app"
	"github.com/dark-vinci/stripchat/beetle/handlers/rest"
	"github.com/dark-vinci/stripchat/beetle/handlers/socket"
	"github.com/dark-vinci/stripchat/beetle/middleware"
	"github.com/dark-vinci/stripchat/beetle/utils"
)

const packageName = "gateway.handlers"

type Handler struct {
	log        *zerolog.Logger
	env        *utils.Environment
	app        app.Operation
	middleware *middleware.Middleware
	engine     *gin.Engine
}

func New(e *utils.Environment, log zerolog.Logger) *Handler {
	a := app.New(&log, e)

	r := gin.Default()
	mw := middleware.New(&log, e, *a)

	logger := log.With().Str(utils.PackageStrHelper, packageName).Logger()

	return &Handler{
		env:        e,
		log:        &logger,
		app:        *a,
		engine:     r,
		middleware: mw,
	}
}

func (h *Handler) Build(ctx context.Context) {
	gin.ForceConsoleColor()

	h.engine.Use(h.middleware.Cors())

	// build endpoints for REST API
	rest.Build(h.engine.Group("/api"), h.app, h.log, h.env)

	// build endpoints for websocket
	socket.New(ctx, *h.log, h.env, h.engine.Group("/socket"), h.app)

}

func (h *Handler) GetEngine() *gin.Engine {
	return h.engine
}
