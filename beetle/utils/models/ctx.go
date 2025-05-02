package models

import (
	"context"
	"github.com/google/uuid"
)

type CTX struct {
	Context   context.Context
	UserAgent string
	UserID    uuid.UUID
	RequestID uuid.UUID
	Email     string
}

func NewCTX(ctx context.Context, userAgent string, requestID uuid.UUID, userID uuid.UUID) *CTX {
	return &CTX{
		Context:   ctx,
		UserAgent: userAgent,
		RequestID: requestID,
		UserID:    userID,
	}
}
