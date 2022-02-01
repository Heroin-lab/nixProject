package middleware

import "net/http"

type Middleware interface {
	PostCheck(next http.HandlerFunc) http.HandlerFunc
}

func PostCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		}
		next(w, r)
	}
}
