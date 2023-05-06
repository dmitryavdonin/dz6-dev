package repository

import (
	user "profile/internal/repository/postgres/user"

	"github.com/dmitryavdonin/gtools/psql"
)

type Repository struct {
	User *user.Repository
}

func NewRepository(pg *psql.Postgres) (*Repository, error) {
	user, err := user.New(pg, user.Options{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		User: user,
	}, nil
}
