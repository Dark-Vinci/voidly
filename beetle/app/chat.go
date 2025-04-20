package app

import (
	"time"

	"github.com/google/uuid"

	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
	"github.com/dark-vinci/stripchat/beetle/utils/models/db"
)

func (a *App) CreateChat(ctx models.CTX, payload models.CreateChatPayload) (*db.Chat, error) {
	log := a.logger.With().
		Str(utils.MethodStrHelper, "app.CreateChat").
		Str(utils.RequestID, utils.GetRequestID(ctx.Context)).
		Logger()

	chat := db.Chat{
		ID:        uuid.New(),
		UserA:     payload.UserA,
		UserB:     payload.UserB,
		CreatedAt: time.Now(),
	}

	c, err := a.chatStore.Create(ctx.Context, chat)
	if err != nil {
		log.Err(utils.UnableToInsert).Msg("Unable to create chat")
		return nil, utils.NotFound
	}

	// todo; publish to redis for any interested consumer

	return c, err
}

func (a *App) GetUserChats(ctx models.CTX, userID uuid.UUID, pagination models.Pagination) (*models.Paginated[[]db.Chat], error) {
	log := a.logger.With().
		Str(utils.MethodStrHelper, "app.GetUserChats").
		Str(utils.RequestID, utils.GetRequestID(ctx.Context)).
		Logger()

	log.Info().Msg("Operation to fetch users chats")

	result, err := a.chatStore.GetByUserID(ctx.Context, userID, pagination)
	if err != nil {
		log.Err(utils.UnableToPerformOperation).Msg("Unable to get list of chats")
		return nil, utils.UnableToPerformOperation
	}

	return result, nil
}
