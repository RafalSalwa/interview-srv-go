package middlewares

import (
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/pkg/csrf"
	"github.com/gorilla/mux"
	"net/http"
)

func CSRFMiddleware(cfg csrf.Config) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("X-CSRF-Token")
			if token == "" {
				responses.RespondNotAuthorized(w, "CSRF token missing")
				return
			}

			if !csrf.ValidateToken(token, cfg) {
				responses.RespondNotAuthorized(w, "Wrong csrf token")
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
