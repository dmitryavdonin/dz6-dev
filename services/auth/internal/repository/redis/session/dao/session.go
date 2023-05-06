package dao

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id         uuid.UUID `json:"id"`
	Login      string    `json:"login"`
	Token      string    `json:"token"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (s Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}
