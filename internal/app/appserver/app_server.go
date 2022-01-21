package appserver

import (
	"encoding/json"
	"fmt"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/repositories/database"
	"github.com/gorilla/mux"
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

	s.configureRouter()

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

func (s *AppServer) configureRouter() error {
	s.router.HandleFunc("/hello", s.handleHello())
	//s.router.HandleFunc("/login", s.handleLogin())

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

func (s *AppServer) handleHello() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := new(request)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(req.Email)
	}
}
