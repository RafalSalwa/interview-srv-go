package http

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/RafalSalwa/interview/util/logger"
	"github.com/golang-jwt/jwt"
)

func BasicAuth(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok || len(strings.TrimSpace(u)) < 1 || len(strings.TrimSpace(p)) < 1 {
			unauthorised(w)
			return
		}
		if u != os.Getenv("AUTH_USERNAME") || p != os.Getenv("AUTH_PASSWORD") {
			unauthorised(w)
			return
		}

		handler(w, r)
	}
}

func unauthorised(rw http.ResponseWriter) {
	rw.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
	rw.WriteHeader(http.StatusUnauthorized)
}

func isAuthorized(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header["X-Access-Token"]
		if accessToken == nil {
			respondNotAuthorized(w, "Access token is missing")
			return
		}

		secretKey := os.Getenv("JWT_SECRET_KEY")
		if secretKey == "" {
			logger.Log("JWT SECRET KEY IS MISSING IN ENV FILE", logger.Error)
			respondNotAuthorized(w, "")
			return
		}

		var mySigningKey = []byte(secretKey)

		token, err := jwt.Parse(accessToken[0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("JWT token parsing error")
			}

			return mySigningKey, nil
		})

		if err != nil {
			respondNotAuthorized(w, "JWT token expired")
			return
		}

		if !token.Valid {
			respondNotAuthorized(w, "JWT token invalid")
			return
		}
	}
}
