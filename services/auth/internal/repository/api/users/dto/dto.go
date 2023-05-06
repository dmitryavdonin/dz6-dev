package dto

import (
	"time"

	"github.com/google/uuid"
)

type ReadUserByCredetinalsRequest struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type ReadUserByCredetinalsResponse struct {
	Id         uuid.UUID `json:"id"`
	Login      string    `json:"login"`
	Name       string    `json:"name"`
	MiddleName string    `json:"middleName"`
	Surname    string    `json:"surname"`
	Phone      string    `json:"phone"`
	City       string    `json:"city"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
