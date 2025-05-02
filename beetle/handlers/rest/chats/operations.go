package chats

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
)

func (c *chatsApi) getChatList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			ct         models.CTX //todo; update this
			userID     = ctx.Query("user_id")
			err        error
			rID        = utils.GetRequestID(ctx.Request.Context())
			endpoint   = ctx.FullPath()
			pagination models.Pagination
		)

		ct.Context = ctx.Request.Context()

		requestID, _ := uuid.Parse(rID)

		log := c.z.With().Str(utils.LogEndpointLevel, endpoint).
			Str(utils.RequestID, requestID.String()).Logger()

		uuID, err := uuid.Parse(userID)
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

		chats, err := c.a.GetUserChats(ct, uuID, pagination)
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

		utils.OkResponse(ctx, http.StatusCreated, "user chats fetched", chats)
	}
}

func (c *chatsApi) createChat() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			ct       models.CTX //todo; update this
			payload  models.CreateChatPayload
			err      error
			rID      = utils.GetRequestID(ctx.Request.Context())
			endpoint = ctx.FullPath()
		)

		ct.Context = ctx.Request.Context()

		requestID, _ := uuid.Parse(rID)

		log := c.z.With().Str(utils.LogEndpointLevel, endpoint).
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

		chat, err := c.a.CreateChat(ct, payload)
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
