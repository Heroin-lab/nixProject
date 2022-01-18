package database

import (
	"github.com/Heroin-lab/nixProject/repositories/models"
)

type UserRepos struct {
	storage *Storage
}

func (r *UserRepos) Create(u *models.User) (*models.User, error) {
	u.BeforeCreate()
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

//func GetProfile(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "GET":
//		claims, err := ValidateToken(GetTokenFromBearerString(r.Header.Get("Authorization")), refreshSecret)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusUnauthorized)
//			return
//		}
//
//		user, err := NewUserRepository().GetUserByID(claims.ID)
//		if err != nil {
//			http.Error(w, "User does not exist", http.StatusBadRequest)
//			return
//		}
//
//		resp := UserResponse{
//			ID:    user.ID,
//			Name:  user.Name,
//			Email: user.Email,
//		}
//
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(resp)
//	default:
//		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
//	}
//}
