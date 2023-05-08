package middlewares

import (
	"github.com/gorilla/mux"
	"net/http"
)

func CorsMiddleware() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Connection", "keep-alive")
			w.Header().Add("Access-Control-Allow-Origin", "http://localhost:8081")
			w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
			w.Header().Add("Access-Control-Allow-Headers", "Authorization, content-type")
			w.Header().Add("Access-Control-Max-Age", "86400")

			h.ServeHTTP(w, r)
		})
	}
}
