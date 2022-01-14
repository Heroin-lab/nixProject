package database

import (
	"github.com/Heroin-lab/nixProject/repositories/models"
)

type UserRepos struct {
	storage *Storage
}

func (r *UserRepos) Create(u *models.User) (*models.User, error) {
	if err := r.storage.db.QueryRow(
		"INSERT INTO users (email, password) VALUES (?, ?)",
		u.Email,
		u.Password,
	).Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
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
