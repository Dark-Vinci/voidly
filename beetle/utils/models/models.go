package models

import "github.com/google/uuid"

type CreateAccountPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateMessagePayload struct {
	Content string    `json:"content"`
	ChatID  uuid.UUID `json:"chat_id"`
	UserID  uuid.UUID `json:"user_id"`
}

type CreateChatPayload struct {
	UserA uuid.UUID `json:"user_a"`
	UserB uuid.UUID `json:"user_b"`
}

type UserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Paginated[T any] struct {
	hasNextPage bool  `json:"has_next_page"`
	hasPrevPage bool  `json:"has_prev_page"`
	totalCount  int64 `json:"total_count"`
	pageSize    int64 `json:"page_size"`
	pageNumber  int64 `json:"page_number"`
	payload     T     `json:"payload"`
}

func NewPaginated[T any](p Pagination, payload T) *Paginated[T] {
	return &Paginated[T]{}
}

type Pagination struct {
	PageSize   int
	PageNumber int
}

func (p Pagination) GetOffset() int {
	offset := (p.PageNumber - 1) * p.PageSize

	return offset
}
