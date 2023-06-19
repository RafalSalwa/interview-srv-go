package auth

import (
	"net/http"
	"strings"

	apiHandler "github.com/RafalSalwa/interview-app-srv/api/handler"
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
)

type bearerTokenHanndler struct {
	h     apiHandler.HandlerFunc
	token string
}

func (a *bearerTokenHanndler) middleware(h apiHandler.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prefix := "Bearer "
		authHeader := r.Header.Get("Authorization")
		reqToken := strings.TrimPrefix(authHeader, prefix)
		if authHeader == "" || reqToken == authHeader {
			responses.RespondNotAuthorized(w, "Authentication header not present or malformed")
			return
		}
		h(w, r)
	}
}

func newBearerTokenMiddleware(h apiHandler.HandlerFunc, token string) *bearerTokenHanndler {
	return &bearerTokenHanndler{
		h:     h,
		token: token,
	}
}
