package router

import (
	"github.com/RafalSalwa/interview-app-srv/api/handler"
	"github.com/RafalSalwa/interview-app-srv/util/auth"
	_ "github.com/RafalSalwa/interview-app-srv/util/auth"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAuthRouter(h handler.IAuthHandler, validator *validator.Validate) http.Handler {
	router := mux.NewRouter()

	router.Methods(http.MethodGet).Path("/login").HandlerFunc(h.Login())
	router.Methods(http.MethodGet).Path("/logout").HandlerFunc(auth.BasicAuth(h.Logout()))
	router.Methods(http.MethodGet).Path("/logout").HandlerFunc(auth.BasicAuth(h.Logout()))
	router.Methods(http.MethodGet).Path("/logout").HandlerFunc(auth.BasicAuth(h.Logout()))
	return router
}

func setupAuthRoutes(r *mux.Router, h handler.IUserHandler) {

}
