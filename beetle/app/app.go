package app

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/store/connection"
	"github.com/dark-vinci/stripchat/beetle/store/dblayer"
	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
	"github.com/dark-vinci/stripchat/beetle/utils/models/db"
)

const packageName = "beetle.app"

type Operation interface {
	Dummy(ctx models.CTX, payload string) string

	// auth & user
	CreateAccount(ctx models.CTX, payload models.CreateAccountPayload) (*db.User, error)
	LoginToAccount(ctx models.CTX, payload models.LoginPayload) (*uuid.UUID, error)
	DeleteAccount(ctx models.CTX, userID uuid.UUID) error
	GetUserByID(ctx models.CTX, userID uuid.UUID) (*db.User, error)

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
	redis        utils.Redis
	store        *connection.Store
}

func New(z *zerolog.Logger, e *utils.Environment) *Operation {
	log := z.With().Str(utils.PackageStrHelper, packageName).Logger()

	store := connection.NewStore(log, e)

	messageStore := dblayer.NewMessage(store)
	chatStore := dblayer.NewChat(store)
	userStore := dblayer.NewUser(store)

	red := utils.NewRedis(&log, e.RedisURL, e.RedisPassword, e.RedisPassword)

	app := &App{
		logger:       &log,
		store:        store,
		messageStore: *messageStore,
		chatStore:    *chatStore,
		userStore:    *userStore,
		redis:        *red,
	}

	result := Operation(app)

	return &result
}
