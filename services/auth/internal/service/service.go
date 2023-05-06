package service

import (
	"auth/internal/repository"
	"auth/internal/service/auth"
	"auth/pkg/tools/tokenManager"
	"time"

	"github.com/dmitryavdonin/gtools/logger"
)

type Service struct {
	Auth auth.Auth
}

func NewServices(repository *repository.Repository, sessionTTL time.Duration, logger logger.Interface) (*Service, error) {
	tm, err := tokenManager.NewManager(sessionTTL)
	if err != nil {
		return nil, err
	}
	auth, err := auth.NewAuth(repository.Session, repository.UsersApi, tm, logger)
	if err != nil {
		return nil, err
	}

	return &Service{
		auth,
	}, nil
}
