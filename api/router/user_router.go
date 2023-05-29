package router

import (
	"github.com/RafalSalwa/interview-app-srv/api/resource/middlewares"
	"github.com/RafalSalwa/interview-app-srv/config"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/RafalSalwa/interview-app-srv/api/handler"
)

func RegisterUserRouter(r *mux.Router, h handler.UserHandler, conf *config.Conf) {
	s := r.PathPrefix("/user").Subrouter()

	s.Use(middlewares.ValidateJWTAccessToken(conf.Token))

	s.Methods(http.MethodGet).Path("/{id}").HandlerFunc(h.GetUserById())
	s.Methods(http.MethodPost).Path("").HandlerFunc(h.CreateUser())
	s.Methods(http.MethodPost).Path("/change_password").HandlerFunc(h.PasswordChange())
	s.Methods(http.MethodPost).Path("/validate/{code}").HandlerFunc(h.ValidateCode())
}
