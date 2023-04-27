package handler

import (
	"github.com/RafalSalwa/interview-app-srv/util/logger"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/services"
	"github.com/gorilla/mux"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type UserHandler interface {
	GetUser() HandlerFunc
	PostUser() HandlerFunc
	PasswordChange() HandlerFunc
	Create() HandlerFunc
	UserExist() HandlerFunc
	LogIn() HandlerFunc
}

type handler struct {
	Router         *mux.Router
	userSqlService services.UserSqlService
	logger         *logger.Logger
}

func NewUserHandler(r *mux.Router, us services.UserSqlService, l *logger.Logger) UserHandler {
	return handler{r, us, l}
}
