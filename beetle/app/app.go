package app

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/store/dblayer"
	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
	"github.com/dark-vinci/stripchat/beetle/utils/models/db"
)

type Operation interface {
	Dummy(ctx models.CTX, payload string) string

	// auth
	CreateAccount(ctx models.CTX, payload models.CreateAccountPayload) (*db.User, error)
	LoginToAccount(ctx models.CTX, payload models.LoginPayload) error
	DeleteAccount(ctx models.CTX, userID uuid.UUID) error

	// message
	CreateMessage(ctx models.CTX, payload models.CreateMessagePayload) (*db.Message, error)
	GetChatMessages(ctx models.CTX, chatID uuid.UUID, pagination models.Pagination) (*models.Paginated[[]db.Message], error)

	// chat
	CreateChat(ctx models.CTX, payload models.CreateChatPayload) (*db.Chat, error)
	GetUserChats(ctx models.CTX, userID uuid.UUID, pagination models.Pagination) (*models.Paginated[[]db.Chat], error)
}

type App struct {
	userStore    dblayer.UserStore
	messageStore dblayer.MessageStore
	chatStore    dblayer.ChatStore
	logger       *zerolog.Logger
	env          *utils.Environment
	redis        *utils.Client
}

func New(z *zerolog.Logger, e *utils.Environment) *Operation {
	app := &App{}

	result := Operation(app)

	return &result
}
