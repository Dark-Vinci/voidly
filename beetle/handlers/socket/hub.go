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

		h.redis.Subscribe(context.Background(), "message", b)

		for {
			select {
			case msg := <-b:
				var c models.Message

				if err := json.Unmarshal(msg, &c); err != nil {
					h.logger.Err(err).Msg("Error unmarshalling message")
				}

				// ignore message sent by the same server
				//if c.Server != h.ServerName.String() && len(c.Server) != 0 {
				h.Broadcast <- msg
				//}
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

				// write to client
			case message := <-h.Broadcast:
				// broadcast to other servers
				go func() {
					// todo: retry -> MESSAGE MUST BE SENT, TRY AS MANY TIMES AS POSSIBLE
					_ = h.redis.Broadcast(context.Background(), "redis-key", message)
				}()

				// todo: add go routine
				for _, client := range h.Clients {
					select {
					case client.Send <- message:
						//todo: log the info
						h.logger.Info().Msg("message received from client")
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
