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

type Message struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

type MessageStore interface {
	Create(ctx context.Context, payload db.Message) (*db.Message, error)
	GetByChatID(ctx context.Context, chatID uuid.UUID, pagination models.Pagination) (*models.Paginated[[]db.Message], error)
}

func NewMessage(conn *connection.Store) *MessageStore {
	log := conn.Log.With().
		Str(utils.FunctionNameHelper, "NewMessage").
		Str(utils.PackageStrHelper, packageName).
		Logger()

	user := &Message{
		logger: &log,
		db:     conn.Connection,
	}

	userDB := MessageStore(user)

	return &userDB
}

func (m *Message) GetByChatID(ctx context.Context, chatID uuid.UUID, pagination models.Pagination) (*models.Paginated[[]db.Message], error) {
	log := m.logger.With().
		Str(utils.MethodStrHelper, "message.GetByChatID").
		Str(utils.RequestID, utils.GetRequestID(ctx)).
		Logger()

	if pagination.PageNumber < 1 {
		pagination.PageNumber = 1
	}

	offset := pagination.GetOffset()

	var messages []db.Message

	if err := m.db.WithContext(ctx).Where("chat_id = ?", chatID).Offset(offset).Limit(pagination.PageSize).Find(&messages).Error; err != nil {
		log.Err(err).Msg("Error finding messages")
		return nil, err
	}

	result := models.NewPaginated(pagination, messages)

	return result, nil
}

func (m *Message) Create(ctx context.Context, payload db.Message) (*db.Message, error) {
	log := m.logger.With().
		Str(utils.MethodStrHelper, "message.Create").
		Str(utils.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request to create message")

	if err := m.db.WithContext(ctx).Create(&payload).Error; err != nil {
		log.Err(err).Msg("Failed to get create message")
		return nil, err
	}

	return &payload, nil
}
