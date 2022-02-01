package database

import (
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/repositories/models"
)

type UserRepos struct {
	storage *Storage
}

func (r *UserRepos) Create(u *models.User) error {
	u.BeforeCreate()
	_, err := r.storage.db.Exec(
		"INSERT INTO users (email, password) VALUES (?, ?)",
		u.Email,
		u.Password,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepos) GetByEmail(email string) (*models.User, error) {
	u := &models.User{}
	if err := r.storage.db.QueryRow(
		"SELECT id, email, password FROM users WHERE email = ?", email).Scan(
		&u.Id,
		&u.Email,
		&u.Password,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepos) UpdatePassword(u *models.ChangePassModel) error {
	encPass, err := models.EncryptString(u.NewPass)
	if err != nil {
		return err
	}

	_, err = r.storage.db.Exec("UPDATE users SET password=? WHERE email=?",
		encPass,
		u.Email)
	if err != nil {
		return err
	}
	logger.Info("User with email='" + u.Email + "' was successfully change the password!")
	return nil
}
