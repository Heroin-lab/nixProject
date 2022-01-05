package database

import (
	"database/sql"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Storage struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepos
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	db, err := sql.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	} else {
		logger.DebugMsg("DB was successfully connected")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	s.db = db
	return nil
}

func (s *Storage) Close() {
	logger.DebugMsg("DB connection was closed!")
	s.db.Close()
}

// User repos
func (s *Storage) User() *UserRepos {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepos{
		storage: s,
	}

	return s.userRepository
}
