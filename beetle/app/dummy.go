package app

import (
	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
)

func (a *App) Dummy(ctx models.CTX, payload string) string {
	log := a.logger.With().
		Str(utils.MethodStrHelper, "app.LoginToAccount").
		Str(utils.RequestID, utils.GetRequestID(ctx.Context)).
		Logger()

	log.Info().Msg("Got request to serve dummy")

	return "dummy world"
}
