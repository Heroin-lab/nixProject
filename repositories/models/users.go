package models

<<<<<<< HEAD
import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       string
=======
import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
>>>>>>> c56ec5407024cea03fdda6c0210eab953b96d09a
	Email    string
	Password string
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.Password = enc
	}

	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
