package socket

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/stripchat/beetle/app"
	"github.com/dark-vinci/stripchat/beetle/utils"
	"github.com/dark-vinci/stripchat/beetle/utils/models"
	"github.com/dark-vinci/stripchat/beetle/utils/models/db"
)

type Hub struct {
	app        app.Operation
	redis      utils.Redis
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	ctx        context.Context
	ctxCancel  context.CancelFunc
	mu         sync.Mutex
	ServerName uuid.UUID
	logger     zerolog.Logger
}

func NewHub(ctx context.Context, logger zerolog.Logger, e *utils.Environment, app app.Operation) *Hub {
	red := utils.NewRedis(&logger, e.RedisURL, e.RedisPassword, e.RedisUsername)

	c, cancel := context.WithCancel(ctx)

	return &Hub{
		app:        app,
		logger:     logger,
		ServerName: uuid.New(),
		redis:      *red,
		ctx:        c,
		ctxCancel:  cancel,
		mu:         sync.Mutex{},
		Clients:    make(map[string]*Client), // user_id -> client
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Start() {
	// subscribe to events posted on redis
	go func() {
		b := make(chan []byte)

		h.redis.Subscribe(context.Background(), utils.WebsocketMessageChannel, b)

		for {
			select {
			case msg := <-b:
				var c models.Message

				if err := json.Unmarshal(msg, &c); err != nil {
					h.logger.Err(err).Msg("Error unmarshalling message")
				}

				// ignore message sent by the same server
				h.Broadcast <- msg
			}
		}
	}()

	go func() {
		for {
			select {
			// register a client
			case client := <-h.Register:
				h.Clients[client.UserID] = client

			//delete a client
			case client := <-h.Unregister:
				if _, ok := h.Clients[client.UserID]; ok {
					delete(h.Clients, client.UserID)
					close(client.Send)
				}

			// write to a client
			case message := <-h.Broadcast:
				var m db.Message

				if err := json.Unmarshal(message, &m); err != nil {
					h.logger.Err(err).Msg("Error unmarshalling message")
					continue
				}

				if client, ok := h.Clients[m.ChatID.String()]; ok {
					select {
					case client.Send <- message:
						h.logger.Info().Msg("message sent to client")
					default:
						// if we cant send, close the send channel and delete client
						close(client.Send)
						delete(h.Clients, client.UserID)
					}
				}
			}
		}
	}()
}
