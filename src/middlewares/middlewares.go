package middlewares

import (
	"api-devbook/src/utils/auth"
	"api-devbook/src/utils/response"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func HandleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			response.Erro(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
