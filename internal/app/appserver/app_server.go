package appserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type AppServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *AppServer {
	return &AppServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *AppServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configreRouter()

	s.logger.Info("Starting app server")

	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *AppServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *AppServer) configreRouter() error {
	s.router.HandleFunc("/hello", s.handleHello())

	return nil
}

func (s *AppServer) handleHello() http.HandlerFunc {
	//type request struct {
	//	name string
	//}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
