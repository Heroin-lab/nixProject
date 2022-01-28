package appserver

import (
	"encoding/json"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/repositories/database"
	"github.com/Heroin-lab/nixProject/repositories/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type server struct {
	router  *http.ServeMux
	storage *database.Storage
}

func NewServer(store *database.Storage) *server {
	s := &server{
		router:  http.NewServeMux(),
		storage: store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/register", s.handleUsersCreate())
	s.router.HandleFunc("/login", s.handleUsersLogin())
	s.router.HandleFunc("/get-items-by-category", s.handleGetProductsByCategory())
	s.router.HandleFunc("/insert-item", s.handleInsertProduct())
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"Error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		req := new(models.LoginRequest)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &models.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.storage.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
	}
}

func (s *server) handleUsersLogin() http.HandlerFunc {

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

		user, err := s.storage.User().GetByEmail(req.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		accessString, err := GenerateToken(user.Id, 10, "super_secret_key")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := GenerateToken(user.Id, 60, "super_secret_key(no)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := models.LoginResponse{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}

		s.respond(w, r, 200, resp)
	}
}

func (s *server) handleGetProductsByCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var req = new(models.CategoryRequest)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		getItems, err := s.storage.Product().GetByCategory(req.Category_name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		s.respond(w, r, 200, getItems)
	}
}

func (s *server) handleInsertProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var req = new(models.Products)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		insertItem, err := s.storage.Product().InsertItem(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		s.respond(w, r, 200, insertItem)
	}
}

//func (s *server) handleDeleteProduct() http.HandlerFunc {
//	if r.Method != "POST" {
//		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
//		return
//	}
//}
