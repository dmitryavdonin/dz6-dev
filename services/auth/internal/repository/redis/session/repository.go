package repository

import (
	"auth/pkg/redis"
)

type Repository struct {
	*redis.Redis
}

func New(redis *redis.Redis) (*Repository, error) {
	var r = &Repository{redis}
	return r, nil
}
