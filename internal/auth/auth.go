package auth

import (
	"net/http"
	"os"
	"strings"

	"github.com/RafalSalwa/interview-app-srv/api/handler"
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
)

func BasicAuth(h handler.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok || len(strings.TrimSpace(u)) < 1 || len(strings.TrimSpace(p)) < 1 {
			responses.RespondNotAuthorized(w, "")
			return
		}

		if u != os.Getenv("AUTH_USERNAME") || p != os.Getenv("AUTH_PASSWORD") {
			responses.RespondNotAuthorized(w, "")
			return
		}

		h(w, r)
	}
}

func Authorization(h handler.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("x-api-key")
		if accessToken == "" {
			responses.RespondNotAuthorized(w, "")
			return
		}
		h(w, r)
	}
}
