package utils

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
)

//const packageName = "util.redis"

type Client struct {
	Val *redis.Client
}

//go:generate mockgen -source redis.go -destination ./mock/redis_mock.go -package mock RedisOps
type Redis interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
	Broadcast(ctx context.Context, key string, value []byte) error
	Subscribe(ctx context.Context, key string, messageChannel chan<- []byte)
	Close() error
}

func NewRedis(z *zerolog.Logger, addr, password, username string) *Redis {
	log := z.With().Str(PackageStrHelper, packageName).Logger()

	r := redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis server address
		Password: password, // No password set
		DB:       0,        // Use default DB
		Username: username,
	})

	log.Info().Msg("connected to redis db")

	red := &Client{
		Val: r,
	}

	redOps := Redis(red)

	return &redOps
}

func (r *Client) Get(ctx context.Context, key string) ([]byte, error) {
	return r.Val.WithContext(ctx).Get(ctx, key).Bytes()
}

func (r *Client) Close() error {
	return r.Val.Close()
}

func (r *Client) Set(ctx context.Context, key string, value []byte) error {
	return r.Val.WithContext(ctx).Set(ctx, key, value, 0).Err()
}

func (r *Client) Broadcast(ctx context.Context, key string, value []byte) error {
	if publish := r.Val.Publish(ctx, key, value); publish.Err() != nil {
		return publish.Err()
	}

	return nil
}

func (r *Client) Subscribe(ctx context.Context, key string, messageChannel chan<- []byte) {
	subscription := r.Val.Subscribe(ctx, key)
	ch := subscription.Channel()

	go func() {
		for msg := range ch {
			messageChannel <- []byte(msg.Payload)
		}
	}()
}
