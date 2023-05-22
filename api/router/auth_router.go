package router

import (
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/internal/auth"

	"github.com/gorilla/mux"

	"github.com/RafalSalwa/interview-app-srv/api/handler"
)

func RegisterAuthRouter(r *mux.Router, h handler.IAuthHandler) {
	s := r.PathPrefix("/auth/").Subrouter()
	s.Methods(http.MethodGet).Path("/login").HandlerFunc(auth.Authorization(h.Login()))
	s.Methods(http.MethodGet).Path("/token").HandlerFunc(auth.Authorization(h.Login()))
	s.Methods(http.MethodGet).Path("/token/refresh").HandlerFunc(auth.Authorization(h.Login()))
}
