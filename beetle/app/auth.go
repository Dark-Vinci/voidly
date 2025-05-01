package app

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
	"github.com/dark-vinci/stripchat/beetle/utils/models/db"
)

func (a *App) CreateAccount(ctx models.CTX, payload models.CreateAccountPayload) (*db.User, error) {
	log := a.logger.With().
		Str(utils.MethodStrHelper, "app.CreateAccount").
		Str(utils.RequestID, utils.GetRequestID(ctx.Context)).
		Logger()

	_, err := a.userStore.GetByEmail(ctx.Context, payload.Email)
	if err == nil {
		log.Err(utils.ErrorAlreadyExist).Msg("User already exists")
		return nil, utils.ErrorAlreadyExist
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Err(utils.HashingError).Msg("hashing error")
		return nil, utils.HashingError
	}

	user := db.User{
		ID:        uuid.New(),
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	response, err := a.userStore.Create(ctx.Context, user)
	if err != nil {
		log.Err(utils.UnableToInsert).Msg("error occurred when inserting user")
		return nil, utils.UnableToInsert
	}

	return response, nil
}

func (a *App) LoginToAccount(ctx models.CTX, payload models.LoginPayload) error {
	log := a.logger.With().
		Str(utils.MethodStrHelper, "app.LoginToAccount").
		Str(utils.RequestID, utils.GetRequestID(ctx.Context)).
		Logger()

	user, err := a.userStore.GetByEmail(ctx.Context, payload.Email)
	if err != nil {
		log.Err(utils.NotFound).Msg("User not found")
		return utils.NotFound
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		log.Err(utils.InvalidCredentials).Msg("invalid credentials")
		return utils.InvalidCredentials
	}

	return nil
}

func (a *App) DeleteAccount(ctx models.CTX, userID uuid.UUID) error {
	log := a.logger.With().
		Str(utils.MethodStrHelper, "app.DeleteAccount").
		Str(utils.RequestID, utils.GetRequestID(ctx.Context)).
		Logger()

	_, err := a.userStore.GetByID(ctx.Context, userID)
	if err != nil {
		log.Err(utils.NotFound).Msg("User not found")
		return utils.NotFound
	}

	if err = a.userStore.Delete(ctx.Context, userID, time.Now()); err != nil {
		log.Err(utils.UnableToPerformOperation).Msg("error occurred when deleting user")
		return utils.UnableToPerformOperation
	}

	return nil
}
