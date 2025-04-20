package models

import "github.com/google/uuid"

type CreateAccountPayload struct {
	Username string
	Email    string
	Password string
}

type LoginPayload struct {
	Email    string
	Password string
}

type CreateMessagePayload struct {
	Content string
	ChatID  uuid.UUID
	UserID  uuid.UUID
}

type CreateChatPayload struct {
	UserA uuid.UUID
	UserB uuid.UUID
}

type UserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Paginated[T any] struct {
	hasNextPage bool
	hasPrevPage bool
	totalCount  int64
	pageSize    int64
	pageNumber  int64
	payload     T
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
