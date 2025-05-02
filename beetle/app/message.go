package app

import (
	"time"

	"github.com/google/uuid"

	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
	"github.com/dark-vinci/stripchat/beetle/utils/models/db"
)

func (a *App) CreateMessage(ctx models.CTX, payload models.CreateMessagePayload) (*db.Message, error) {
	log := a.logger.With().
		Str(utils.MethodStrHelper, "app.CreateMessage").
		Str(utils.RequestID, utils.GetRequestID(ctx.Context)).
		Logger()

	log.Info().Msg("Got request to create a message")

	message := db.Message{
		ID:         uuid.New(),
		Content:    payload.Content,
		ChatID:     payload.ChatID,
		FromUserID: ctx.UserID,
		CreatedAt:  time.Now(),
	}

	res, err := a.messageStore.Create(ctx.Context, message)
	if err != nil {
		log.Err(err).Msg("Failed to create a message")
		return nil, utils.UnableToInsert
	}

	// todo; push to message broker
	if err = a.publish(ctx.Context, *res); err != nil {
		log.Err(err).Msg("Failed to publish a message")
		return nil, utils.UnableToInsert
	}

	return res, nil
}

func (a *App) GetChatMessages(ctx models.CTX, chatID uuid.UUID, pagination models.Pagination) (*models.Paginated[[]db.Message], error) {
	log := a.logger.With().
		Str(utils.MethodStrHelper, "app.GetChatMessages").
		Str(utils.RequestID, utils.GetRequestID(ctx.Context)).
		Logger()

	log.Info().Msg("Got request to get chat messages")

	res, err := a.messageStore.GetByChatID(ctx.Context, chatID, pagination)
	if err != nil {
		log.Err(err).Msg("Failed to get chat messages")
		return nil, utils.UnableToPerformOperation
	}

	return res, nil
}
