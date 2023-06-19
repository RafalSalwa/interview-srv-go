package auth

import (
	"net/http"
	"strings"

	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	apiHandler "github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
)

type bearerTokenHandler struct {
	h     apiHandler.HandlerFunc
	token string
}

func (a *bearerTokenHandler) middleware(h apiHandler.HandlerFunc) http.HandlerFunc {
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

func newBearerTokenMiddleware(h apiHandler.HandlerFunc, token string) *bearerTokenHandler {
	return &bearerTokenHandler{
		h:     h,
		token: token,
	}
}
