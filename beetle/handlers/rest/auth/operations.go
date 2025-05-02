package auth

import (
	"net/http"

	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
	"github.com/gin-gonic/gin"
)

func (a *authApi) login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			payload  models.LoginPayload
			err      error
			endpoint = ctx.FullPath()
		)

		c := utils.GetFromContext[models.CTX](ctx.Request.Context(), utils.CTX)

		log := a.z.With().Str(utils.LogEndpointLevel, endpoint).
			Str(utils.RequestID, c.RequestID.String()).Logger()

		if err = ctx.ShouldBind(&payload); err != nil {
			log.Err(err).Msg("bad request")
			utils.ErrorResponse(ctx, http.StatusBadRequest, utils.ErrorData{
				ID:      c.RequestID,
				Details: err.Error(),
				Status:  http.StatusBadRequest,
			})
			ctx.Abort()
			return
		}

		if err = a.a.LoginToAccount(c, payload); err != nil {
			log.Err(err).Msg("Invalid credentials")
			utils.ErrorResponse(ctx, http.StatusUnauthorized, utils.ErrorData{
				ID:      c.RequestID,
				Details: err.Error(),
				Status:  http.StatusUnauthorized,
			})
			ctx.Abort()
			return
		}

		cred := a.m.GenerateCredentials()

		response := struct {
			Token string `json:"token"`
		}{
			Token: cred,
		}

		utils.OkResponse(ctx, http.StatusCreated, "user login token generated", response)
	}
}

func (a *authApi) create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			payload  models.CreateAccountPayload
			err      error
			endpoint = ctx.FullPath()
		)

		c := utils.GetFromContext[models.CTX](ctx.Request.Context(), utils.CTX)

		log := a.z.With().Str(utils.LogEndpointLevel, endpoint).
			Str(utils.RequestID, c.RequestID.String()).Logger()

		if err = ctx.ShouldBind(&payload); err != nil {
			log.Err(err).Msg("bad request")
			utils.ErrorResponse(ctx, http.StatusBadRequest, utils.ErrorData{
				ID:      c.RequestID,
				Details: err.Error(),
				Status:  http.StatusBadRequest,
			})
			ctx.Abort()
			return
		}

		response, err := a.a.CreateAccount(c, payload)
		if err != nil {
			log.Err(err).Msg("bad request")
			utils.ErrorResponse(ctx, http.StatusBadRequest, utils.ErrorData{
				ID:      c.RequestID,
				Details: err.Error(),
				Status:  http.StatusBadRequest,
			})
			ctx.Abort()
			return
		}

		utils.OkResponse(ctx, http.StatusCreated, "user account created", response)
	}
}
