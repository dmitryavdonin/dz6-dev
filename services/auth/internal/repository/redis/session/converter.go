package repository

import (
	"auth/internal/domain/session"
	"auth/internal/repository/redis/session/dao"
)

func (r *Repository) toDaoSession(session *session.Session) (daoSession *dao.Session) {
	return &dao.Session{
		Id:         session.Id(),
		Login:      session.Login(),
		Token:      session.Token(),
		CreatedAt:  session.CreatedAt(),
		ModifiedAt: session.ModifiedAt(),
	}
}

func (r *Repository) toDomainSession(daoSession *dao.Session) (*session.Session, error) {
	return session.NewSessionWithId(daoSession.Id, daoSession.Login, daoSession.Token, daoSession.CreatedAt, daoSession.ModifiedAt)
}
