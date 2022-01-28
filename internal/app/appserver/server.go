package appserver

import (
	"github.com/Heroin-lab/nixProject/repositories"
	"net/http"
)

type server struct {
	router  *http.ServeMux
	storage repositories.Store
}

func NewServer(store repositories.Store) *server {
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
	//s.router.HandleFunc("/register", s.handleUsersCreate())
	//s.router.HandleFunc("/login", s.handleUsersLogin())
	//s.router.Handle("/get-items-by-category", s.handleGetProductsByCategory())
	//s.router.HandleFunc("/insert-item", s.handleInsertProduct())
}

//func (s *server) handleUsersCreate() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		req := new(models.LoginRequest)
//		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//			logger.Error("Server respond with bad request status!")
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//
//		_, err := s.storage.User().GetByEmail(req.Email)
//		if err == nil {
//			http.Error(w, "Already exists", http.StatusConflict)
//			return
//		}
//
//		u := &models.User{
//			Email:    req.Email,
//			Password: req.Password,
//		}
//
//		if err := s.storage.User().Create(u); err != nil {
//			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
//			return
//		}
//	}
//}
//
//func (s *server) handleUsersLogin() http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		switch r.Method {
//		case "POST":
//			req := new(models.LoginRequest)
//
//			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//				http.Error(w, err.Error(), http.StatusBadRequest)
//				return
//			}
//
//			user, err := s.storage.User().GetByEmail(req.Email)
//			if err != nil {
//				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
//				return
//			}
//
//			if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
//				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
//				return
//			}
//
//			accessString, err := GenerateToken(user.Id, s.config.AccessLifetimeMin, s.config.AccessSecretStr)
//			if err != nil {
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//
//			refreshString, err := GenerateToken(user.Id, s.config.RefreshLifetimeMin, s.config.RefreshSecretStr)
//			if err != nil {
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//
//			resp := models.LoginResponse{
//				AccessToken:  accessString,
//				RefreshToken: refreshString,
//			}
//
//			w.WriteHeader(http.StatusOK)
//			json.NewEncoder(w).Encode(resp)
//		default:
//			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
//		}
//	}
//}
//
//func (s *server) handleGetProductsByCategory() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		switch r.Method {
//		case "POST":
//			var req = new(models.CategoryRequest)
//
//			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//				http.Error(w, err.Error(), http.StatusBadRequest)
//				return
//			}
//
//			getItems, err := s.storage.Product().GetByCategory(req.Category_name)
//			if err != nil {
//				http.Error(w, err.Error(), http.StatusConflict)
//				return
//			}
//
//			w.WriteHeader(http.StatusOK)
//			json.NewEncoder(w).Encode(getItems)
//		default:
//			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
//		}
//	}
//}
//
//func (s *server) handleInsertProduct() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		switch r.Method {
//		case "POST":
//			var req = new(models.Products)
//
//			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//				http.Error(w, err.Error(), http.StatusBadRequest)
//				return
//			}
//
//			insertItem, err := s.storage.Product().InsertItem(req)
//			if err != nil {
//				http.Error(w, err.Error(), http.StatusConflict)
//				return
//			}
//
//			w.WriteHeader(http.StatusOK)
//			json.NewEncoder(w).Encode(insertItem)
//		default:
//			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
//		}
//	}
//}
