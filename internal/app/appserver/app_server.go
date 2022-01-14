package appserver

import (
	"encoding/json"
	"fmt"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/repositories/database"
	"github.com/Heroin-lab/nixProject/repositories/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
	s.router.HandleFunc("/login", s.Login())

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

func (s *AppServer) Login() http.HandlerFunc {
	//type request struct {
	//	name string
	//}

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			req := new(models.LoginRequest)
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			repos := database.UserRepos{}
			user, err := repos.GetByEmail(req.Email)
			if err != nil {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}

			if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}

			//	accessString, err := GenerateToken(user.ID, accessLifetimeMinutes, accessSecret)
			//	if err != nil {
			//		http.Error(w, err.Error(), http.StatusInternalServerError)
			//		return
			//	}
			//
			//	refreshString, err := GenerateToken(user.ID, refreshLifetimeMinutes, refreshSecret)
			//	if err != nil {
			//		http.Error(w, err.Error(), http.StatusInternalServerError)
			//		return
			//	}
			//
			//	resp := LoginResponse{
			//		AccessToken:  accessString,
			//		RefreshToken: refreshString,
			//	}
			//
			//	w.WriteHeader(http.StatusOK)
			//	json.NewEncoder(w).Encode(resp)
			//default:
			//	http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			fmt.Println(user.Id)
		}
	}
}
