package dao

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `db:"id"`
	Login      string    `db:"login"`
	Password   string    `db:"password"`
	Name       string    `json:"name"`
	Middlename string    `json:"middle_name"`
	Surname    string    `json:"surname"`
	Phone      string    `json:"phone"`
	City       string    `json:"city"`
	Role       string    `db:"role"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}

var UserColumns = []string{
	"id",
	"login",
	"password",
	"name",
	"middle_name",
	"surname",
	"phone",
	"city",
	"role",
	"created_at",
	"modified_at",
}
