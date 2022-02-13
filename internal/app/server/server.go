package server

import (
	"database/sql"
	logger "github.com/Heroin-lab/heroin-logger/v3"
	"github.com/Heroin-lab/nixProject/configs"
	"github.com/Heroin-lab/nixProject/internal/app/server/handlers"
	"github.com/Heroin-lab/nixProject/middleware"
	"github.com/Heroin-lab/nixProject/repositories/database"

	"net/http"
	"time"
)

type Server struct {
	Router           *http.ServeMux
	Storage          *database.Storage
	UserHandler      *handlers.UserHandler
	ProductHandler   *handlers.ProductHandler
	SuppliersHandler *handlers.SuppliersHandler
	OrderHandler     *handlers.OrderHandler
}

func NewServer(store *database.Storage) *Server {
	s := &Server{
		Router:           http.NewServeMux(),
		Storage:          store,
		ProductHandler:   handlers.NewProductHandler(store),
		UserHandler:      handlers.NewUserHandler(store),
		SuppliersHandler: handlers.NewSuppliersHandler(store),
		OrderHandler:     handlers.NewOrderHandler(store),
	}

	s.configureRouter()

	return s
}

func Start(config *configs.Config) error {
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

	db.SetConnMaxIdleTime(time.Minute * 3)
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

func (s *Server) configureRouter() {
	// USERS handlers
	s.Router.HandleFunc("/register", middleware.PostCheck(s.UserHandler.HandleUsersCreate()))
	s.Router.HandleFunc("/login", middleware.PostCheck(s.UserHandler.HandleUsersLogin()))
	s.Router.HandleFunc("/update-token", middleware.GetCheck(s.UserHandler.HandleRefreshTokens()))

	s.Router.HandleFunc("/change-password", middleware.PatchCheck(
		middleware.UserTokenCheck(
			s.UserHandler.HandleChangePassword())))

	// PRODUCTS handlers
	s.Router.HandleFunc("/get-items-by-category", middleware.PostCheck(
		s.ProductHandler.HandleGetProductsByCategory()))

	s.Router.HandleFunc("/insert-item", middleware.PostCheck(
		middleware.AdminTokenCheck(
			s.ProductHandler.HandleInsertProduct())))

	s.Router.HandleFunc("/delete-item", middleware.DeleteCheck(
		middleware.AdminTokenCheck(
			s.ProductHandler.HandleDeleteProduct())))

	s.Router.HandleFunc("/update-item", middleware.PatchCheck(
		middleware.AdminTokenCheck(
			s.ProductHandler.HandleUpdateProduct())))

	// SUPPLIERS handlers
	s.Router.HandleFunc("/get-suppliers-by-category", middleware.PostCheck(
		s.SuppliersHandler.HandleGetSuppliersByCategory()))

	s.Router.HandleFunc("/add-supplier", middleware.PostCheck(
		middleware.AdminTokenCheck(
			s.SuppliersHandler.HandleAddSupplier())))

	s.Router.HandleFunc("/delete-supplier", middleware.DeleteCheck(
		middleware.AdminTokenCheck(
			s.SuppliersHandler.HandleDeleteSupplier())))

	s.Router.HandleFunc("/update-supplier", middleware.PatchCheck(
		middleware.AdminTokenCheck(
			s.SuppliersHandler.HandleUpdateSupplier())))

	// ORDERS handlers
	s.Router.HandleFunc("/get-all-user-orders", middleware.PostCheck(
		middleware.UserTokenCheck(
			s.OrderHandler.HandleGetAllUserOrders())))

	s.Router.HandleFunc("/add-order", middleware.PostCheck(
		middleware.UserTokenCheck(
			s.OrderHandler.HandleAddOrder())))

	s.Router.HandleFunc("/delete-order", middleware.DeleteCheck(
		middleware.UserTokenCheck(
			s.OrderHandler.HandleDeleteOrder())))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
