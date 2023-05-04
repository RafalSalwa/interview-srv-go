package middlewares

import (
	"github.com/gorilla/mux"
	"net/http"
)

func ContentTypeJson() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json;charset=utf8")
			h.ServeHTTP(w, r)
		})
	}
}
