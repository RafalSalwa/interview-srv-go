package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/RafalSalwa/interview-app-srv/internal/jwt"
	"github.com/gorilla/mux"
)

func ValidateJWTAccessToken() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("authorization")
			if authorizationHeader != "" {
				bearerToken := strings.Split(authorizationHeader, " ")
				if len(bearerToken) == 2 {
					token, err := jwt.DecodeToken(bearerToken[1])
					if err != nil {
						json.NewEncoder(w).Encode(Exception{Message: error.Error()})
						return
					}
					if token.Valid {
						next.ServeHTTP(w, r)
					} else {
						json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
					}
				}
			} else {
				json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
			}
		})
	}
}
