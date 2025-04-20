package socket

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/app"
	"github.com/dark-vinci/stripchat/beetle/utils"
)

const packageName = "handler.websocket"

func New(ctx context.Context, log zerolog.Logger, e *utils.Environment, r *gin.RouterGroup, a app.Operation) {
	logger := log.With().Str(utils.PackageStrHelper, packageName).Logger()
	hub := NewHub(ctx, log, e, a)

	ws := WebSocket{
		upgrade: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		Hub:    hub,
		logger: logger,
	}

	ws.Build(r)
}

type WebSocket struct {
	upgrade websocket.Upgrader
	Hub     *Hub
	logger  zerolog.Logger
}

func (ws *WebSocket) Build(endpoint *gin.RouterGroup) {
	ws.Hub.Start()

	endpoint.GET("/ws", func(c *gin.Context) {
		ws.serveWs(ws.Hub, c.Writer, c.Request)
	})
}

func (ws *WebSocket) serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		ws.logger.Err(err).Msgf("error upgrading ws: %v", err)
		return
	}

	client := NewClient(hub, conn, ws.logger)

	// write to user
	go client.WritePump()
	// read from user
	go client.ReadPump()
}
