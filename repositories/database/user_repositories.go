package database

import "github.com/Heroin-lab/nixProject/internal/app/models"

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
	return nil, nil
}
