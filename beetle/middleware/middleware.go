package middleware

import (
	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/app"
	"github.com/dark-vinci/stripchat/beetle/utils"
)

const packageName = "beetle.middleware"

type Middleware struct {
	logger zerolog.Logger
	env    *utils.Environment
	app    app.Operation
}

func New(l *zerolog.Logger, e *utils.Environment, a app.Operation) *Middleware {
	logger := l.With().Str(utils.PackageStrHelper, packageName).Logger()

	return &Middleware{
		logger: logger,
		env:    e,
		app:    a,
	}
}
