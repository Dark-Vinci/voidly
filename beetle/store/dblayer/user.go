package dblayer

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/stripchat/beetle/store/connection"
	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models/db"
)

type User struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

type UserStore interface {
	Create(ctx context.Context, payload db.User) (*db.User, error)
	GetByEmail(ctx context.Context, email string) (*db.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*db.User, error)
	Delete(ctx context.Context, id uuid.UUID, now time.Time) error
}

func NewUser(conn *connection.Store) *UserStore {
	log := conn.Log.With().
		Str(utils.FunctionNameHelper, "NewUser").
		Str(utils.PackageStrHelper, packageName).
		Logger()

	user := &User{
		logger: &log,
		db:     conn.Connection,
	}

	userDB := UserStore(user)

	return &userDB
}

func (u *User) GetByEmail(ctx context.Context, email string) (*db.User, error) {
	log := u.logger.With().
		Str(utils.MethodStrHelper, "user.GetByEmail").
		Str(utils.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request to get a user by email")

	var user db.User

	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		log.Err(err).Msg("Failed to get user by email")
		return nil, err
	}

	return &user, nil
}

func (u *User) Delete(ctx context.Context, id uuid.UUID, now time.Time) error {
	log := u.logger.With().
		Str(utils.MethodStrHelper, "channel.DeleteByID").
		Str(utils.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got request to delete channel with ID %v", id)

	if err := u.db.WithContext(ctx).Model(&db.User{}).Where("id = ?", id).UpdateColumns(db.User{DeletedAt: &now}).Error; err != nil {
		log.Err(err).Msg("Failed to delete channel")

		return err
	}

	return nil
}

func (u *User) GetByID(ctx context.Context, id uuid.UUID) (*db.User, error) {
	log := u.logger.With().
		Str(utils.MethodStrHelper, "user.GetByEmail").
		Str(utils.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request to get a user by id")

	var user db.User

	if err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		log.Err(err).Msg("Failed to get user by id")
		return nil, err
	}

	return &user, nil
}

func (u *User) Create(ctx context.Context, payload db.User) (*db.User, error) {
	log := u.logger.With().
		Str(utils.MethodStrHelper, "user.Create").
		Str(utils.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request to create user")

	if err := u.db.WithContext(ctx).Create(&payload).Error; err != nil {
		log.Err(err).Msg("Failed to get create user")
		return nil, err
	}

	return &payload, nil
}
