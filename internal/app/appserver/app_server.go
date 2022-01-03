package appserver

import (
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type AppServer struct {
	config *Config
	router *mux.Router
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

	logger.Info("Starting app server")

	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *AppServer) configureLogger() error {
	logger.SetLogLevel(s.config.LogLevel)
	logger.DebugMsg("App logger was started in debug mod!")
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
