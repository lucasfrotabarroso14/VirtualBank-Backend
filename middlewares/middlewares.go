package middlewares

import (
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/auth"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/responses"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s ", r.Method, r.RequestURI, r.Host)
		next(w, r)

	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.ValidateToken(r); erro != nil {
			responses.JSON(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
