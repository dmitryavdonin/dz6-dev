package auth

import (
	"auth/internal/domain/session"
	repo "auth/internal/service/auth/adapters/repository"
	"auth/pkg/tools/tokenManager"
	"context"

	"github.com/dmitryavdonin/gtools/logger"
	"github.com/google/uuid"
)

type auth struct {
	repositoy    repo.Session
	usersApi     repo.UsersApi
	tokenManager tokenManager.TokenManager
	logger       logger.Interface
}

func NewAuth(repositoy repo.Session, usersApi repo.UsersApi, tokenManager tokenManager.TokenManager, logger logger.Interface) (*auth, error) {
	return &auth{
		repositoy:    repositoy,
		usersApi:     usersApi,
		tokenManager: tokenManager,
		logger:       logger,
	}, nil
}

func (a *auth) SignIn(ctx context.Context, login, pass string) (sessionId uuid.UUID, err error) {
	user, err := a.usersApi.ReadUserByCredetinals(ctx, &repo.ReadUserByCredetinalsParams{Login: login, Pass: pass})
	if err != nil {
		return
	}

	token, err := a.tokenManager.NewJWT(tokenManager.AuthInfo{UserID: user.Id.String(), Login: user.Login})
	if err != nil {
		return
	}

	session, err := session.NewSession(login, token)
	if err != nil {
		return
	}
	err = a.repositoy.CreateSession(ctx, session)
	if err != nil {
		return
	}

	a.logger.Debug("auth: SignIn(): SUCCESS! login = " + login + "; pass = " + pass + "; session = " + sessionId.String())
	a.logger.Debug("auth: SignIn(): token = " + token)

	return session.Id(), nil
}

func (a *auth) ReadSessionById(ctx context.Context, id uuid.UUID) (session *session.Session, err error) {
	return a.repositoy.ReadSessionById(ctx, id)
}

func (a *auth) DeleteSessionById(ctx context.Context, id uuid.UUID) (err error) {
	return a.repositoy.DeleteSessionById(ctx, id)
}
