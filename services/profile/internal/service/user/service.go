package user

import (
	userRepo "profile/internal/repository/postgres/user"
	adaptersRepo "profile/internal/service/user/adapters/repository"
	"time"

	"github.com/dmitryavdonin/gtools/logger"
)

type service struct {
	repository adaptersRepo.User
	logger     logger.Interface
	//tokenManager tokenManager.TokenManager
	options Options
}

type Options struct {
	TokenTTL time.Duration
}

func (uc *service) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}

func New(repository *userRepo.Repository, logger logger.Interface, options Options) (*service, error) {
	service := &service{
		repository: repository,
		logger:     logger,
	}

	service.SetOptions(options)
	return service, nil
}
