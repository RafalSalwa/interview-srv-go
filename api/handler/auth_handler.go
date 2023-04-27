package handler

import (
	"github.com/RafalSalwa/interview-app-srv/util/logger"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/services"
	"github.com/gorilla/mux"
)

type AuthHandlerFunc func(http.ResponseWriter, *http.Request)

type AuthHandler interface {
	GetUser() HandlerFunc
	PostUser() HandlerFunc
	PasswordChange() HandlerFunc
	Create() HandlerFunc
	UserExist() HandlerFunc
	LogIn() HandlerFunc
}

type AuthHandler struct {
	Router         *mux.Router
	userSqlService services.UserSqlService
	logger         *logger.Logger
}

func NewAuthHandler(r *mux.Router, us services.UserSqlService, l *logger.Logger) UserHandler {
	return AuthHandler{r, us, l}
}
