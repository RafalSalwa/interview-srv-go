package handler

import (
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/RafalSalwa/interview-app-srv/util/logger"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/services"
	"github.com/gorilla/mux"
)

type AuthHandlerFunc func(http.ResponseWriter, *http.Request)

type IAuthHandler interface {
	SignUpUser(*models.SignUpInput) (*models.UserDBResponse, error)
	SignInUser(*models.SignInInput) (*models.UserDBResponse, error)
	Login() AuthHandlerFunc
	Logout() AuthHandlerFunc
}

type AuthHandler struct {
	Router         *mux.Router
	userSqlService services.AuthService
	logger         *logger.Logger
}

func NewAuthHandler(r *mux.Router, us services.AuthService, l *logger.Logger) AuthHandler {
	return AuthHandler{r, us, l}
}
