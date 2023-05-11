package handler

import (
	"encoding/json"
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/RafalSalwa/interview-app-srv/services"
)

type IAuthHandler interface {
	SignUpUser(request *models.CreateUserRequest) HandlerFunc
	SignInUser(request *models.LoginUserRequest) HandlerFunc
	Login() HandlerFunc
	Logout() HandlerFunc
	Token() HandlerFunc
}

type AuthHandler struct {
	Router         *mux.Router
	userORMService services.AuthService
	logger         *logger.Logger
}

func NewAuthHandler(r *mux.Router, us services.AuthService, l *logger.Logger) IAuthHandler {
	return AuthHandler{r, us, l}
}

func (a AuthHandler) SignUpUser(input *models.CreateUserRequest) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a AuthHandler) SignInUser(input *models.LoginUserRequest) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (ah AuthHandler) Login() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		LoginReq := &models.LoginUserRequest{}
		err := json.NewDecoder(r.Body).Decode(LoginReq)

		if err != nil {
			ah.logger.Error().Err(err).Msg("Decode Request")
			responses.RespondBadRequest(w, "")
			return
		}
		user, err := ah.userORMService.Load(LoginReq)
		if err != nil {
			ah.logger.Err(err)
			responses.RespondInternalServerError(w)
			return
		}
		responses.NewUserResponse(user, w, r)
	}
}

func (a AuthHandler) Logout() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a AuthHandler) Token() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
