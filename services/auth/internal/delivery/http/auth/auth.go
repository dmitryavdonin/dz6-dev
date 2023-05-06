package auth

import "github.com/google/uuid"

type SignInRequest struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type SignInResponse struct {
	SessionId uuid.UUID `json:"session_id"`
}
