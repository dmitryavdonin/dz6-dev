package repository

import (
	"context"
	"profile/internal/domain/user"

	"github.com/google/uuid"
)

type User interface {
	CreateUser(ctx context.Context, user *user.User) (err error)
	UpdateUser(ctx context.Context, id uuid.UUID, upFunc func(oldUser *user.User) (*user.User, error)) (user *user.User, err error)
	DeleteUserById(ctx context.Context, id uuid.UUID) (err error)
	ReadUserById(ctx context.Context, id uuid.UUID) (user *user.User, err error)
	ReadUserByLogin(ctx context.Context, login string) (user *user.User, err error)
}
