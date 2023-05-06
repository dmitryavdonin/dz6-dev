package repository

import (
	usersApi "auth/internal/repository/api/users"
	session "auth/internal/repository/redis/session"
	"auth/pkg/redis"
)

type Repository struct {
	Session  *session.Repository
	UsersApi *usersApi.Repository
}

func NewRepository(redis *redis.Redis, usersApiUri string) (repo *Repository, err error) {
	session, err := session.New(redis)
	if err != nil {
		return
	}

	usersApi, err := usersApi.New(usersApiUri)
	if err != nil {
		return
	}

	return &Repository{
		UsersApi: usersApi,
		Session:  session,
	}, nil
}
