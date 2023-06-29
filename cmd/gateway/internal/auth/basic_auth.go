package auth

import (
	"crypto/sha256"
	"crypto/subtle"
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	apiHandler "github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
	"net/http"
)

type basicAuth struct {
	h        apiHandler.AuthHandler
	username string
	password string
}

func (a *basicAuth) Middleware(h apiHandler.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(a.username))
			expectedPasswordHash := sha256.Sum256([]byte(a.password))

			usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1
			if usernameMatch && passwordMatch {
				h(w, r)
				return
			}
		}
		responses.RespondNotAuthorized(w, "")
		return
	}
}

func newBasicAuthMiddleware(h apiHandler.AuthHandler, username string, password string) *basicAuth {
	return &basicAuth{
		h:        h,
		username: username,
		password: password,
	}
}
