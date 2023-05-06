package repository

import (
	"auth/internal/domain/session"
	"context"

	"github.com/google/uuid"
)

type Session interface {
	CreateSession(ctx context.Context, session *session.Session) (err error)
	DeleteSessionById(ctx context.Context, id uuid.UUID) (err error)
	ReadSessionById(ctx context.Context, id uuid.UUID) (session *session.Session, err error)
}

type UsersApi interface {
	ReadUserByCredetinals(ctx context.Context, params *ReadUserByCredetinalsParams) (user *User, err error)
}
