package app

import (
	"context"
	"encoding/json"

	"github.com/dark-vinci/stripchat/beetle/utils/models/db"
)

func (a *App) publish(ctx context.Context, message db.Message) error {
	bty, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return a.redis.Broadcast(ctx, "message", bty)
}
