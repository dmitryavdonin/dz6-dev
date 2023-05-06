package repository

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID
	Login      string
	Name       string
	MiddleName string
	Surname    string
	Phone      string
	City       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type ReadUserByCredetinalsParams struct {
	Login string
	Pass  string
}
