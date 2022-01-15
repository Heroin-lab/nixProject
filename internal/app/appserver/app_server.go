package appserver

import (
	"encoding/json"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	database "github.com/Heroin-lab/nixProject/repositories/database"
	"github.com/Heroin-lab/nixProject/repositories/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type AppServer struct {
	config  *Config
	router  *mux.Router
	storage *database.Storage
}

func New(config *Config) *AppServer {
	return &AppServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *AppServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configreRouter()

	if err := s.configureStorage(); err != nil {
		return err
	}

	logger.Info("Starting app server")

	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *AppServer) configureLogger() error {
	logger.SetLogLevel(s.config.LogLevel)
	logger.DebugMsg("App logger was started in debug mod!")
	return nil
}

func (s *AppServer) configreRouter() error {
	s.router.HandleFunc("/register", s.handleUsersCreate())
	s.router.HandleFunc("/login", s.handleUsersLogin())

	return nil
}

func (s *AppServer) configureStorage() error {
	st := database.New(s.config.Storage)
	if err := st.Open(); err != nil {
		return err
	}

	s.storage = st
	return nil
}

func (s *AppServer) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(models.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Server respond with bad request status!")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := s.storage.User().GetByEmail(req.Email)
		if err == nil {
			http.Error(w, "Already exists", http.StatusConflict)
			return
		}

		u := &models.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if _, err := s.storage.User().Create(u); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
	}

}

func (s *AppServer) handleUsersLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
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
			convId, _ := strconv.Atoi(user.Id)

			if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}

			accessString, err := GenerateToken(convId, s.config.AccessLifetimeMin, s.config.AccessSecretStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			refreshString, err := GenerateToken(convId, s.config.RefreshLifetimeMin, s.config.RefreshSecretStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			resp := models.LoginResponse{
				AccessToken:  accessString,
				RefreshToken: refreshString,
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
		default:
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		}
	}
}
