package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CTX struct {
	Context   context.Context
	userAgent string
	userID    uuid.UUID
	requestID uuid.UUID
	timezone  time.Location
}

func NewCTX(ctx context.Context, userAgent string, requestID uuid.UUID, userID uuid.UUID) *CTX {
	return &CTX{
		Context:   ctx,
		userAgent: userAgent,
		requestID: requestID,
		userID:    userID,
	}
}
