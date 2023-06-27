package middlewares

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/pkg/auth"

	"github.com/RafalSalwa/interview-app-srv/pkg/jwt"
	"github.com/gorilla/mux"
)

func ValidateJWTAccessToken(c auth.JWTConfig) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("Authorization")
			if authorizationHeader != "" {
				bearerToken := strings.Split(authorizationHeader, " ")
				if len(bearerToken) == 2 {
					sub, err := jwt.ValidateToken(bearerToken[1], c.Access.PublicKey)
					if err != nil {
						responses.RespondNotAuthorized(w, "Wrong access token")
						return
					}
					r.Header.Set("x-user-id", strconv.FormatInt(sub.ID, 10))
				}
			} else {
				responses.RespondNotAuthorized(w, "An authorization header is required")
			}
			h.ServeHTTP(w, r)
		})
	}
}
