package auth

import (
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	simpleHandler "github.com/RafalSalwa/interview-app-srv/pkg/simple_handler"
	"net/http"
	"strings"
)

type bearerTokenHanndler struct {
	h     simpleHandler.HandlerFunc
	token string
}

func (a *bearerTokenHanndler) middleware(h simpleHandler.HandlerFunc) http.HandlerFunc {
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

func newBearerTokenMiddleware(h simpleHandler.HandlerFunc, token string) *bearerTokenHanndler {
	return &bearerTokenHanndler{
		h:     h,
		token: token,
	}
}
