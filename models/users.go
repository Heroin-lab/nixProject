package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Email    string
	Password string
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := EncryptString(u.Password)
		if err != nil {
			return err
		}

		u.Password = enc
	}

	return nil
}

func EncryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
