package profile

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	id         uuid.UUID
	userId     uuid.UUID
	name       string
	middleName string
	surname    string
	phone      string
	city       string
	createdAt  time.Time
	modifiedAt time.Time
}

func NewProfileWithId(
	id uuid.UUID,
	userId uuid.UUID,
	name string,
	middleName string,
	surname string,
	phone string,
	city string,
	createdAt time.Time,
	modifiedAt time.Time,
) (*Profile, error) {
	return &Profile{
		id:         id,
		userId:     userId,
		name:       name,
		middleName: middleName,
		surname:    surname,
		phone:      phone,
		city:       city,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,
	}, nil
}

func NewProfile(
	userId uuid.UUID,
	name string,
	middleName string,
	surname string,
	phone string,
	city string,
) (*Profile, error) {
	return &Profile{
		id:         uuid.New(),
		userId:     userId,
		name:       name,
		middleName: middleName,
		surname:    surname,
		phone:      phone,
		city:       city,
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}, nil
}

func (p Profile) Id() uuid.UUID {
	return p.id
}

func (p Profile) UserId() uuid.UUID {
	return p.userId
}

func (p Profile) Name() string {
	return p.name
}

func (p Profile) MiddleName() string {
	return p.middleName
}

func (p Profile) Surname() string {
	return p.surname
}

func (p Profile) Phone() string {
	return p.phone
}

func (p Profile) City() string {
	return p.city
}

func (p Profile) CreatedAt() time.Time {
	return p.createdAt
}

func (p Profile) ModifiedAt() time.Time {
	return p.modifiedAt
}
