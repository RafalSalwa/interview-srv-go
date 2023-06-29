package router

import (
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/auth"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterAuthRouter(r *mux.Router, h handler.AuthHandler, cfg *config.Config) error {
	s := r.PathPrefix("/auth/").Subrouter()
	authorizer, err := auth.NewAuthorizer(h, cfg)
	if err != nil {
		return err
	}
	s.Methods(http.MethodPost).Path("/signup").HandlerFunc(authorizer.Middleware(h.SignUpUser()))
	s.Methods(http.MethodPost).Path("/signin").HandlerFunc(authorizer.Middleware(h.SignInUser()))
	s.Methods(http.MethodGet).Path("/verify/{code}").HandlerFunc(authorizer.Middleware(h.Verify()))
	s.Methods(http.MethodPost).Path("/code").HandlerFunc(authorizer.Middleware(h.GetVerificationCode()))
	s.Methods(http.MethodGet).Path("/token/refresh").HandlerFunc(authorizer.Middleware(h.RefreshToken()))
	return nil
}
