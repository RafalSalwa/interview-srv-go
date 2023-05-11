package middlewares

import (
	"net/http"
	"strings"

	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/internal/jwt"
	"github.com/gorilla/mux"
)

func DeserializeUser(c config.ConfToken) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var access_token string

			authorizationHeader := r.Header.Get("Authorization")
			fields := strings.Fields(authorizationHeader)

			if len(fields) != 0 && fields[0] == "Bearer" {
				access_token = fields[1]
			}

			if access_token == "" {
				responses.NewUnauthorizedResponse()
				return
			}

			_, err := jwt.ValidateToken(access_token, c.AccessTokenPublicKey)
			if err != nil {
				responses.NewUnauthorizedResponse()
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
