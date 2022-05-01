package handlers

import (
	"encoding/json"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/configs"
	"github.com/Heroin-lab/nixProject/models"
	"github.com/Heroin-lab/nixProject/repositories/database"
	"github.com/Heroin-lab/nixProject/services"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserHandler struct {
	storage *database.Storage
}

func NewUserHandler(str *database.Storage) *UserHandler {
	return &UserHandler{
		storage: str,
	}
}

func (h *UserHandler) HandleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(models.LoginRequest)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &models.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := h.storage.User().Create(u); err != nil {
			services.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
	}
}

func (h *UserHandler) HandleUsersLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		conf := configs.Config{}
		req := new(models.LoginRequest)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
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

		accessString, err := services.GenerateToken(user.Id, user.Role, 10, conf.AccessSecretStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := services.GenerateToken(user.Id, user.Role, 60, conf.RefreshSecretStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := models.LoginResponse{
			UserId:       user.Id,
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}

		services.Respond(w, r, 200, resp)
	}
}

func (h *UserHandler) HandleChangePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(models.ChangePassModel)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			services.Error(w, r, http.StatusBadRequest, err)
			return
		}

		user, err := h.storage.User().GetByEmail(req.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusNotFound)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPass)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusConflict)
			return
		}

		if err := h.storage.User().UpdatePassword(req); err != nil {
			http.Error(w, "DB error", http.StatusUnprocessableEntity)
			return
		}

		services.Respond(w, r, 200, "Password was successfully changed!")
	}
}

func (h *UserHandler) HandleRefreshTokens() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		conf := configs.Config{}

		if token == "" {
			http.Error(w, "Token header is empty!", http.StatusMethodNotAllowed)
			return
		}

		tokenWithoutBearer, _ := services.GetTokenFromBearerString(token)

		tokenClaims, err := services.ValidateToken(tokenWithoutBearer, conf.RefreshSecretStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := h.storage.User().GetById(tokenClaims.ID)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		accessString, err := services.GenerateToken(user.Id, user.Role, 10, conf.AccessSecretStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := services.GenerateToken(user.Id, user.Role, 60, conf.RefreshSecretStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := models.LoginResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}

		services.Respond(w, r, 200, resp)
	}
}

func (h *UserHandler) HandleAdminTokenCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		services.Respond(w, r, 200, "Admin")
	}
}
