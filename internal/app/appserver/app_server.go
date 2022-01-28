package appserver

import (
	"database/sql"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/repositories/database"
	"net/http"
	"time"
)

func Start(config *Config) error {
	if err := configureLogger(config.LogLevel); err != nil {
		return err
	}

	sqlDB, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	store := database.New(sqlDB)
	srv := NewServer(store)

	logger.Info("Starting app server")

	return http.ListenAndServe(config.BindAddress, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	logger.DebugMsg("DB was successfully connected")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func configureLogger(logLevel string) error {
	logger.SetLogLevel(logLevel)
	logger.DebugMsg("App logger was started in debug mod!")
	return nil
}
