package handlers

import (
	"encoding/json"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/models"
	"github.com/Heroin-lab/nixProject/repositories/database"
	"github.com/Heroin-lab/nixProject/services"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserHandler struct {
	resService *services.RespondService
	storage    *database.Storage
}

func NewUserHandler(str *database.Storage) *UserHandler {
	return &UserHandler{
		storage: str,
	}
}

func (h *UserHandler) HandleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		req := new(models.LoginRequest)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			h.resService.Error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &models.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := h.storage.User().Create(u); err != nil {
			h.resService.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
	}
}

func (h *UserHandler) HandleUsersLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		req := new(models.LoginRequest)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := h.storage.User().GetByEmail(req.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		accessString, err := services.GenerateToken(user.Id, 10, "super_secret_key")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := services.GenerateToken(user.Id, 60, "super_secret_key(no)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := models.LoginResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}

		h.resService.Respond(w, r, 200, resp)
	}
}

func (h *UserHandler) HandleChangePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		cPM := new(models.ChangePassModel)

		if err := json.NewDecoder(r.Body).Decode(&cPM); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := h.storage.User().GetByEmail(cPM.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusNotFound)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cPM.OldPass)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusConflict)
			return
		}

		if err := h.storage.User().UpdatePassword(cPM); err != nil {
			http.Error(w, "DB error", http.StatusUnprocessableEntity)
			return
		}

		h.resService.Respond(w, r, 200, "Password was successfully changed!")
	}
}
