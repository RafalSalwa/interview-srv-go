package router

import (
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/api/resource/middlewares"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
	"github.com/RafalSalwa/interview-app-srv/pkg/auth"
	"github.com/gorilla/mux"
)

func RegisterUserRouter(r *mux.Router, h handler.UserHandler, cfg auth.JWTConfig) {
	s := r.PathPrefix("/user").Subrouter()
	s.Use(middlewares.ValidateJWTAccessToken(cfg))

	s.Methods(http.MethodGet).Path("/{id}").HandlerFunc(h.GetUserById())
	s.Methods(http.MethodPost).Path("").HandlerFunc(h.CreateUser())
	s.Methods(http.MethodPost).Path("/change_password").HandlerFunc(h.PasswordChange())
	s.Methods(http.MethodPost).Path("/validate/{code}").HandlerFunc(h.ValidateCode())
}
