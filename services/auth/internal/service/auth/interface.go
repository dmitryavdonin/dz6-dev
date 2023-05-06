package auth

import (
	"auth/internal/domain/session"
	"context"

	"github.com/google/uuid"
)

type Auth interface {
	SignIn(ctx context.Context, login, pass string) (sessionId uuid.UUID, err error)
	ReadSessionById(ctx context.Context, id uuid.UUID) (session *session.Session, err error)
	DeleteSessionById(ctx context.Context, id uuid.UUID) (err error)
}
