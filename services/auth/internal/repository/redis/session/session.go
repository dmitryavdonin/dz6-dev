package repository

import (
	"auth/internal/domain/session"
	"auth/internal/repository/redis/session/dao"
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

func (s *Repository) CreateSession(ctx context.Context, session *session.Session) (err error) {
	daoSession := s.toDaoSession(session)
	return s.Redis.SetKey(ctx, session.Id().String(), *daoSession)
}

func (s *Repository) DeleteSessionById(ctx context.Context, id uuid.UUID) (err error) {
	return s.Redis.DeleteKey(ctx, id.String())
}

func (s *Repository) ReadSessionById(ctx context.Context, id uuid.UUID) (sess *session.Session, err error) {
	jsonSession, err := s.Redis.GetKey(ctx, id.String())
	if err != nil {
		return
	}
	daoSession := dao.Session{}
	err = json.Unmarshal([]byte(jsonSession.(string)), &daoSession)
	if err != nil {
		return
	}
	return s.toDomainSession(&daoSession)
}
