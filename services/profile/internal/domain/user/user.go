package user

import (
	"time"

	"profile/internal/domain/user/password"

	"github.com/google/uuid"
)

type User struct {
	id         uuid.UUID
	login      string
	password   password.Password
	name       string
	middleName string
	surname    string
	phone      string
	city       string
	role       string
	createdAt  time.Time
	modifiedAt time.Time
}

func NewUser(
	login string,
	pass string,
	name string,
	middleName string,
	surname string,
	phone string,
	city string,
	role string,
) (*User, error) {
	encryptesPass, err := password.EncryptPassword(pass)
	if err != nil {
		return nil, err
	}
	return &User{
		id:         uuid.New(),
		login:      login,
		password:   encryptesPass,
		name:       name,
		middleName: middleName,
		surname:    surname,
		phone:      phone,
		city:       city,
		role:       role,
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}, nil
}

func NewUserWithId(
	id uuid.UUID,
	login string,
	pass string,
	name string,
	middleName string,
	surname string,
	phone string,
	city string,
	role string,
	createdAt time.Time,
	modifiedAt time.Time,
) *User {
	password := password.NewPassword(pass)
	return &User{
		id:         id,
		login:      login,
		password:   password,
		name:       name,
		middleName: middleName,
		surname:    surname,
		phone:      phone,
		city:       city,
		role:       role,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,
	}
}

func (u User) Id() uuid.UUID {
	return u.id
}

func (u User) Login() string {
	return u.login
}

func (u User) Password() password.Password {
	return u.password
}

func (p User) Name() string {
	return p.name
}

func (p User) MiddleName() string {
	return p.middleName
}

func (p User) Surname() string {
	return p.surname
}

func (p User) Phone() string {
	return p.phone
}

func (p User) City() string {
	return p.city
}

func (u User) Role() string {
	return u.role
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) ModifiedAt() time.Time {
	return u.modifiedAt
}
