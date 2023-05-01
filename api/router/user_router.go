package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/RafalSalwa/interview-app-srv/api/handler"
	"github.com/RafalSalwa/interview-app-srv/util/auth"
)

func RegisterUserRouter(r *mux.Router, h handler.UserHandler) {

	s := r.PathPrefix("/user/").Subrouter()
	s.Methods(http.MethodGet).Path("/{id}").HandlerFunc(auth.BasicAuth(h.GetUserById()))
	s.Methods(http.MethodPost).Path("/{id}").HandlerFunc(auth.BasicAuth(h.PostUser()))
	s.Methods(http.MethodPost).Path("/change_password").HandlerFunc(auth.BasicAuth(h.PasswordChange()))
	s.Methods(http.MethodPost).Path("/auth").HandlerFunc(h.LogIn())
	s.Methods(http.MethodPut).Path("/registration").HandlerFunc(auth.BasicAuth(h.Create()))
	s.Methods(http.MethodPost).Path("/exist").HandlerFunc(auth.BasicAuth(h.UserExist()))

}
