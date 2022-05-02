package middleware

import (
	"github.com/Heroin-lab/nixProject/configs"
	"github.com/Heroin-lab/nixProject/services"
	"net/http"
)

func UserTokenCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secretAccessStr := configs.Config{}
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Token header is empty!", http.StatusMethodNotAllowed)
			return
		}

		tokenWithoutBearer, _ := services.GetTokenFromBearerString(token)

		_, err := services.ValidateToken(tokenWithoutBearer, secretAccessStr.AccessSecretStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func AdminTokenCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secretAccessStr := configs.Config{}
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Token header is empty!", http.StatusMethodNotAllowed)
			return
		}

		tokenWithoutBearer, _ := services.GetTokenFromBearerString(token)

		tokenClaims, err := services.ValidateToken(tokenWithoutBearer, secretAccessStr.AccessSecretStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if tokenClaims.UserRole != "admin" {
			http.Error(w, "You have not enough rights!", http.StatusNotAcceptable)
			return
		}

		next(w, r)
	}
}
