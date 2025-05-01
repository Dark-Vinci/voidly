package socket

import (
	"bytes"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan []byte
	UserID string
	logger zerolog.Logger
}

func NewClient(hub *Hub, conn *websocket.Conn, logger zerolog.Logger) *Client {
	return &Client{
		logger: logger,
		Hub:    hub,
		Conn:   conn,
		Send:   make(chan []byte, 256),
	}
}

// WritePump write to the client
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		_ = c.Conn.Close()
	}()

	//res, err := c.Hub.app.GetUserChats(context.Background(), )

	_ = c.Conn.WriteMessage(websocket.TextMessage, []byte("connected to the server"))

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// hub closed the channel
				_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			_, _ = w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(newline)
				_, _ = w.Write(<-c.Send)
			}

			if err = w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// ReadPump ReadPump reads from the client
func (c *Client) ReadPump() {
	// todo: add const
	defer func() {
		c.Hub.Unregister <- c
		_ = c.Conn.Close()
	}()

	c.Conn.SetReadLimit(readingLimit)
	_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))

		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.Err(err).Msgf("Unexpected close error: %v", err)
			}

			c.logger.Err(err).
				Any("client", c).
				Msgf("websocket read error: %v", err)

			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Hub.Broadcast <- message
	}
}
