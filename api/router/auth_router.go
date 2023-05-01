package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/RafalSalwa/interview-app-srv/api/handler"
	_ "github.com/RafalSalwa/interview-app-srv/util/auth"
)

func RegisterAuthRouter(r *mux.Router, h handler.IAuthHandler) {
	s := r.PathPrefix("/auth/").Subrouter()
	s.Methods(http.MethodGet).Path("/login").HandlerFunc(h.Login())
}
