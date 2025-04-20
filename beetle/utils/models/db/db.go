package db

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	ID        uuid.UUID `json:"id"`
	UserA     uuid.UUID `json:"user1"`
	UserB     uuid.UUID `json:"user2"`
	CreatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	ID        uuid.UUID `json:"id"`
	UserName  string
	Email     string
	Password  string
	CreatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

type Message struct {
	ID         uuid.UUID `json:"id"`
	Content    string    `json:"content"`
	ChatID     uuid.UUID `json:"chat_id"`
	FromUserID uuid.UUID `json:"from_user_id"`
	CreatedAt  time.Time
	DeletedAt  *time.Time
}
