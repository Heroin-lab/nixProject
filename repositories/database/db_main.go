package database

import (
	"database/sql"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	DB             *sql.DB
	UserRepository *UserRepos
	ProdRepository *ProductRepose
	SuppRepose     *SuppRepose
	OrderRepose    *OrderRepose
}

func New(db *sql.DB) *Storage {
	return &Storage{
		DB: db,
	}
}

func (s *Storage) Close() {
	logger.DebugMsg("DB connection was closed!")
	s.DB.Close()
}

// User repos
func (s *Storage) User() *UserRepos {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepos{
		storage: s,
	}

	return s.UserRepository
}

func (s *Storage) Product() *ProductRepose {
	if s.ProdRepository != nil {
		return s.ProdRepository
	}

	s.ProdRepository = &ProductRepose{
		storage: s,
	}

	return s.ProdRepository
}

func (s *Storage) Supplier() *SuppRepose {
	if s.SuppRepose != nil {
		return s.SuppRepose
	}

	s.SuppRepose = &SuppRepose{
		storage: s,
	}

	return s.SuppRepose
}

func (s *Storage) Order() *OrderRepose {
	if s.SuppRepose != nil {
		return s.OrderRepose
	}

	s.OrderRepose = &OrderRepose{
		storage: s,
	}

	return s.OrderRepose
}

//func (s *Storage) Open() error {
//	db, err := sql.Open("mysql", s.config.DatabaseURL)
//	if err != nil {
//		return err
//	}
//
//	if err := db.Ping(); err != nil {
//		return err
//	} else {
//		logger.DebugMsg("DB was successfully connected")
//	}
//
//	db.SetConnMaxLifetime(time.Minute * 3)
//	db.SetMaxOpenConns(10)
//	db.SetMaxIdleConns(10)
//
//	s.db = db
//	return nil
//}
