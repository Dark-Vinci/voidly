package middleware

import (
	"time"

	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/app"
	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models/db"
)

const packageName = "beetle.middleware"
const identityKey = "id"

type Middleware struct {
	logger zerolog.Logger
	env    *utils.Environment
	app    app.Operation
	jwt    *ginJwt.GinJWTMiddleware
}

func New(l *zerolog.Logger, e *utils.Environment, a app.Operation) *Middleware {
	logger := l.With().Str(utils.PackageStrHelper, packageName).Logger()
	jwt, _ := jwtMiddleware(e)

	return &Middleware{
		logger: logger,
		env:    e,
		app:    a,
		jwt:    jwt,
	}
}

func jwtMiddleware(env *utils.Environment) (*ginJwt.GinJWTMiddleware, error) {
	return ginJwt.New(&ginJwt.GinJWTMiddleware{
		Realm: "strip chat",
		Key:   []byte(env.JWTKey),
		PayloadFunc: func(data interface{}) ginJwt.MapClaims {
			if v, ok := data.(*db.User); ok {
				return ginJwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return ginJwt.MapClaims{}
		},
		IdentityKey: identityKey,
		Timeout:     time.Hour * time.Duration(env.JWTExpiration),
	})
}
