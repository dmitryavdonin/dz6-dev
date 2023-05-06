package password

import (
	"errors"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrTooShortPass = errors.New("provided password is too short")
)

type Password string

func (p Password) String() string {
	return string(p)
}

func (p Password) IsEqualTo(input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(input))
	if err != nil {
		return false
	}
	return true
}

func NewPassword(password string) Password {
	return Password(password)
}

func EncryptPassword(notHashedPassword string) (Password, error) {
	if utf8.RuneCountInString(notHashedPassword) < 8 {
		return "", ErrTooShortPass
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(notHashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	password := Password(string(hash))
	return password, nil
}
