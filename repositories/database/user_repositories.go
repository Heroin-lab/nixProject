package database

import (
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/models"
)

type UserRepos struct {
	storage *Storage
}

func (r *UserRepos) Create(u *models.User) error {
	u.BeforeCreate()
	_, err := r.storage.DB.Exec(
		"INSERT INTO users (email, password) VALUES (?, ?)",
		u.Email,
		u.Password,
	)
	if err != nil {
		return err
	}
	logger.Info("User with email: '" + u.Email + "' and password: '" + u.Password + "' was successfully created!")
	return nil
}

func (r *UserRepos) GetByEmail(email string) (*models.User, error) {
	u := &models.User{}
	if err := r.storage.DB.QueryRow(
		"SELECT users.id, email, password, role_name FROM users\n"+
			"INNER JOIN users_roles ur on users.role = ur.id\n"+
			"WHERE email = ?", email).Scan(
		&u.Id,
		&u.Email,
		&u.Password,
		&u.Role,
	); err != nil {
		return nil, err
	}
	logger.Info("A user with this email: '" + u.Email + "' has been successfully found")
	return u, nil
}

func (r *UserRepos) GetById(id int) (*models.User, error) {
	u := &models.User{}
	if err := r.storage.DB.QueryRow(
		"SELECT users.id, email, password, role_name FROM users\n"+
			"INNER JOIN users_roles ur on users.role = ur.id\n"+
			"WHERE users.id = ?", id).Scan(
		&u.Id,
		&u.Email,
		&u.Password,
		&u.Role,
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

	_, err = r.storage.DB.Exec("UPDATE users SET password=? WHERE email=?",
		encPass,
		u.Email)
	if err != nil {
		return err
	}
	logger.Info("User with email: '" + u.Email + "' was successfully change the password!")
	return nil
}
