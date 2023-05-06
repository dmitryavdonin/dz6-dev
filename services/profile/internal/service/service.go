package service

import (
	"profile/internal/repository"
	"profile/internal/service/user"

	"github.com/dmitryavdonin/gtools/logger"
)

type Service struct {
	User user.Service
}

func NewServices(repository *repository.Repository, logger logger.Interface) (*Service, error) {
	user, err := user.New(repository.User, logger, user.Options{})
	if err != nil {
		return nil, err
	}

	return &Service{
		User: user,
	}, nil
}
