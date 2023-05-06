package session

import (
	"time"

	"github.com/google/uuid"
)

type CreateSessionRequest struct {
	Id    uuid.UUID `json:"id"`
	Login string    `json:"login"`
	Role  string    `json:"role"`
}

type CreateSessionResponse struct {
	SessionId uuid.UUID `json:"session_id"`
}

type SessionResponse struct {
	Id         uuid.UUID `json:"id"`
	Login      string    `json:"login"`
	Token      string    `json:"token"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
