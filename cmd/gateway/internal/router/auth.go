package router

import (
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/auth"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterAuthRouter(r *mux.Router, h handler.AuthHandler) {
	s := r.PathPrefix("/auth/").Subrouter()
	s.Methods(http.MethodPost).Path("/signup").HandlerFunc(auth.Authorization(h.SignUpUser()))
	s.Methods(http.MethodPost).Path("/signin").HandlerFunc(auth.Authorization(h.SignInUser()))
	s.Methods(http.MethodGet).Path("/verify/{code}").HandlerFunc(auth.Authorization(h.Verify()))

	s.Methods(http.MethodGet).Path("/token/refresh").HandlerFunc(auth.Authorization(h.RefreshToken()))
}
