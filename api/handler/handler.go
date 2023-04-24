package handler

import (
	"github.com/RafalSalwa/interview-app-srv/util/logger"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/service"
	"github.com/gorilla/mux"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Handler interface {
	GetUser() HandlerFunc
	PostUser() HandlerFunc
	PasswordChange() HandlerFunc
	UserRegistration() HandlerFunc
	UserExist() HandlerFunc
	LogIn() HandlerFunc
}

type handler struct {
	Router         *mux.Router
	userSqlService service.UserSqlService
	logger         *logger.Logger
}

func NewHandler(r *mux.Router, us service.UserSqlService, l *logger.Logger) Handler {
	return handler{r, us, l}
}
