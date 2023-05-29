package middlewares

import (
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/config"
	"net/http"
	"strconv"
	"strings"

	"github.com/RafalSalwa/interview-app-srv/internal/jwt"
	"github.com/gorilla/mux"
)

func ValidateJWTAccessToken(c config.ConfToken) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("authorization")
			if authorizationHeader != "" {
				bearerToken := strings.Split(authorizationHeader, " ")
				if len(bearerToken) == 2 {
					sub, err := jwt.ValidateToken(bearerToken[1], c.AccessTokenPublicKey)
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
