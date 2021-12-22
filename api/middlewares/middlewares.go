package middlewares

import (
	"errors"
	"github.com/Qgolang/golang_fullstack/api/auth"
	"github.com/Qgolang/golang_fullstack/api/responses"
	"net/http"
)

// SetMiddlewareJSON Cài đặt header request
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication Cài đặt authentication cho request
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		next(w, r)
	}
}
