package db

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	ID        uuid.UUID  `json:"id"`
	UserA     uuid.UUID  `json:"user1"`
	UserB     uuid.UUID  `json:"user2"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type User struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Message struct {
	ID         uuid.UUID  `json:"id"`
	Content    string     `json:"content"`
	ChatID     uuid.UUID  `json:"chat_id"`
	FromUserID uuid.UUID  `json:"from_user_id"`
	CreatedAt  time.Time  `json:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
