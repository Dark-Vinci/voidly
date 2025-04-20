package dblayer

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/stripchat/beetle/store/connection"
	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
	"github.com/dark-vinci/stripchat/beetle/utils/models/db"
)

type Chat struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

type ChatStore interface {
	Create(ctx context.Context, payload db.Chat) (*db.Chat, error)
	GetByUserID(ctx context.Context, chatID uuid.UUID, pagination models.Pagination) (*models.Paginated[[]db.Chat], error)
}

func NewChat(conn *connection.Store) *ChatStore {
	log := conn.Log.With().
		Str(utils.FunctionNameHelper, "NewChat").
		Str(utils.PackageStrHelper, packageName).
		Logger()

	user := &Chat{
		logger: &log,
		db:     conn.Connection,
	}

	userDB := ChatStore(user)

	return &userDB
}

func (c *Chat) GetByUserID(ctx context.Context, userID uuid.UUID, pagination models.Pagination) (*models.Paginated[[]db.Chat], error) {
	log := c.logger.With().
		Str(utils.MethodStrHelper, "message.GetByUserID").
		Str(utils.RequestID, utils.GetRequestID(ctx)).
		Logger()

	if pagination.PageNumber < 1 {
		pagination.PageNumber = 1
	}

	offset := pagination.GetOffset()

	var messages []db.Chat

	if err := c.db.WithContext(ctx).Where("user_a = ? || user_b = ?", userID, userID).Offset(offset).Limit(pagination.PageSize).Find(&messages).Error; err != nil {
		log.Err(err).Msg("Error finding chats")
		return nil, err
	}

	result := models.NewPaginated(pagination, messages)

	return result, nil
}

func (c *Chat) Create(ctx context.Context, payload db.Chat) (*db.Chat, error) {
	log := c.logger.With().
		Str(utils.MethodStrHelper, "chat.Create").
		Str(utils.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request to create chat")

	if err := c.db.WithContext(ctx).Create(&payload).Error; err != nil {
		log.Err(err).Msg("Failed to get create chat")
		return nil, err
	}

	return &payload, nil
}
