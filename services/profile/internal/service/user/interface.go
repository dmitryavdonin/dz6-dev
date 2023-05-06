package user

import (
	"context"
	"profile/internal/domain/user"

	"github.com/google/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, user *user.User) (err error)
	UpdateUser(ctx context.Context, id uuid.UUID, upFn func(oldUser *user.User) (*user.User, error)) (*user.User, error)
	DeleteUserById(ctx context.Context, id uuid.UUID) (err error)

	ReadUserById(ctx context.Context, id uuid.UUID) (user *user.User, err error)
	ReadUserByCredetinals(ctx context.Context, login, pass string) (user *user.User, err error)
}
