package user

import (
	"time"

	"github.com/google/uuid"
)

type ReadUserByCredetinalsRequest struct {
	Login string `json:"login" default:"admin"`
	Pass  string `json:"pass" default:"Qwerty123"`
}

type CreateUserRequest struct {
	Login      string `json:"login" default:"admin"`
	Password   string `json:"password" default:"Qwerty123"`
	Name       string `json:"name" default:"Иван"`
	Middlename string `json:"middleName" default:"Иванович"`
	Surname    string `json:"surname" default:"Иванов"`
	Phone      string `json:"phone" default:"79000000000"`
	City       string `json:"city" default:"Маями"`
}

type UpdateUserRequest struct {
	Login      string `json:"login" default:"admin.upd"`
	Password   string `json:"password" default:"Qwerty123.upd"`
	Name       string `json:"name" default:"Иван.upd"`
	Middlename string `json:"middleName" default:"Иванович.upd"`
	Surname    string `json:"surname" default:"Иванов.upd"`
	Phone      string `json:"phone" default:"79111111111"`
	City       string `json:"city" default:"Лондон"`
}

type UpdateUserResponse struct {
	Result UserResponse `json:"result"`
}

type UserResponse struct {
	Id         uuid.UUID `json:"id"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	Middlename string    `json:"middleName"`
	Surname    string    `json:"surname"`
	Phone      string    `json:"phone"`
	City       string    `json:"city"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
