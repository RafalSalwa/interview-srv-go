package auth

import (
	"github.com/RafalSalwa/interview-app-srv/api/handler"
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"net/http"
	"os"
	"strings"
)

func BasicAuth(h handler.HandlerFunc) http.HandlerFunc {
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

		h(w, r)
	}
}

func unauthorised(rw http.ResponseWriter) {
	rw.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
	rw.WriteHeader(http.StatusUnauthorized)
}

func Authorization(h handler.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header["X-API-KEY"]
		if accessToken == nil {
			responses.RespondNotAuthorized(w, "")
			return
		}
		h(w, r)
	}
}
