package middleware

import (
	"net/http"
)

func GetCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}

func PostCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}

func PatchCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			http.Error(w, "Only PATCH method is allowed", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}

func DeleteCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, "Only DELETE method is allowed", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}
