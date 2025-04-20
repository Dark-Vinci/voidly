package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
)

func (m *messageApi) list() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			ct         models.CTX //todo; update this
			chatID     = ctx.Query("chat_id")
			err        error
			requestID  = uuid.New()
			endpoint   = ctx.FullPath()
			pagination models.Pagination
		)

		log := m.z.With().Str(utils.LogEndpointLevel, endpoint).
			Str(utils.RequestID, requestID.String()).Logger()

		cID, err := uuid.Parse(chatID)
		if err != nil {
			log.Err(err).Msg("Invalid uuid")
			utils.ErrorResponse(ctx, http.StatusUnprocessableEntity, utils.ErrorData{
				ID:      requestID,
				Details: err.Error(),
				Status:  http.StatusUnprocessableEntity,
			})
			ctx.Abort()
			return
		}

		messages, err := m.a.GetChatMessages(ct, cID, pagination)
		if err != nil {
			log.Err(err).Msg("something went wrong")
			utils.ErrorResponse(ctx, http.StatusBadGateway, utils.ErrorData{
				ID:      requestID,
				Details: err.Error(),
				Status:  http.StatusBadGateway,
			})
			ctx.Abort()
			return
		}

		utils.OkResponse(ctx, http.StatusCreated, "chat messages fetched", messages)
	}
}

func (m *messageApi) create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			ct        models.CTX //todo; update this
			payload   models.CreateMessagePayload
			err       error
			requestID = uuid.New()
			endpoint  = ctx.FullPath()
		)

		log := m.z.With().Str(utils.LogEndpointLevel, endpoint).
			Str(utils.RequestID, requestID.String()).Logger()

		if err = ctx.ShouldBind(&payload); err != nil {
			log.Err(err).Msg("bad request")
			utils.ErrorResponse(ctx, http.StatusBadRequest, utils.ErrorData{
				ID:      requestID,
				Details: err.Error(),
				Status:  http.StatusBadRequest,
			})
			ctx.Abort()
			return
		}

		chat, err := m.a.CreateMessage(ct, payload)
		if err != nil {
			log.Err(err).Msg("Invalid credentials")
			utils.ErrorResponse(ctx, http.StatusBadGateway, utils.ErrorData{
				ID:      requestID,
				Details: err.Error(),
				Status:  http.StatusBadGateway,
			})
			ctx.Abort()
			return
		}

		utils.OkResponse(ctx, http.StatusCreated, "chat created", chat)
	}
}
