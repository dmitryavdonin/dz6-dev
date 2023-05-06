package user

import (
	"context"
	"errors"
	"profile/internal/domain/user"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("user name not found or incorrect password")
)

func (s *service) ReadUserById(ctx context.Context, id uuid.UUID) (user *user.User, err error) {
	return s.repository.ReadUserById(ctx, id)
}

func (s *service) ReadUserByCredetinals(ctx context.Context, login, pass string) (user *user.User, err error) {
	user, err = s.repository.ReadUserByLogin(ctx, login)
	if err != nil {
		return
	}

	if !user.Password().IsEqualTo(pass) {
		return nil, ErrUserNotFound
	}

	return
}

func (s *service) CreateUser(ctx context.Context, user *user.User) (err error) {
	return s.repository.CreateUser(ctx, user)
}

func (s *service) UpdateUser(ctx context.Context, id uuid.UUID, upFn func(oldUser *user.User) (*user.User, error)) (*user.User, error) {
	return s.repository.UpdateUser(ctx, id, upFn)
}

func (s *service) DeleteUserById(ctx context.Context, id uuid.UUID) (err error) {
	return s.repository.DeleteUserById(ctx, id)
}
