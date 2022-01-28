package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CategoryRequest struct {
	Category_name string `json:"category_name"`
}
