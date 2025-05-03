package app

import (
	"context"
	"encoding/json"

	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models/db"
)

func (a *App) publish(ctx context.Context, message db.Message) error {
	log := a.logger.With().
		Str(utils.MethodStrHelper, "app.publish").
		Str(utils.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got request to publish a message to pub/sub server")

	bty, err := json.Marshal(message)
	if err != nil {
		log.Err(err).Msg("Failed to marshal message")
		return err
	}

	err = a.redis.Broadcast(ctx, utils.WebsocketMessageChannel, bty)
	if err != nil {
		log.Err(err).Msg("Failed to publish message to redis")
		return err
	}

	log.Info().Msg("Successfully published message to redis")

	return nil
}
