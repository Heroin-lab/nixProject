package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePassModel struct {
	Email   string `json:"email"`
	OldPass string `json:"old_pass"`
	NewPass string `json:"new_pass"`
}

type CategoryRequest struct {
	Category_name string `json:"category_name"`
}
