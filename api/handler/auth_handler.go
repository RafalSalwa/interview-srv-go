package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/RafalSalwa/interview-app-srv/services"
	"github.com/RafalSalwa/interview-app-srv/util/logger"
)

type AuthHandlerFunc func(http.ResponseWriter, *http.Request)

type IAuthHandler interface {
	SignUpUser(*models.SignUpInput) AuthHandlerFunc
	SignInUser(*models.SignInInput) AuthHandlerFunc
	Login() AuthHandlerFunc
	Logout() AuthHandlerFunc
	Token() AuthHandlerFunc
}

type AuthHandler struct {
	Router         *mux.Router
	userSqlService services.AuthService
	logger         *logger.Logger
}

func NewAuthHandler(r *mux.Router, us services.AuthService, l *logger.Logger) IAuthHandler {
	return AuthHandler{r, us, l}
}

func (a AuthHandler) SignUpUser(input *models.SignUpInput) AuthHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a AuthHandler) SignInUser(input *models.SignInInput) AuthHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a AuthHandler) Login() AuthHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a AuthHandler) Logout() AuthHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a AuthHandler) Token() AuthHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
