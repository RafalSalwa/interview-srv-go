package handler

import (
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/util/logger"
	"net/http"
	"strconv"

	"github.com/RafalSalwa/interview-app-srv/services"
	"github.com/gorilla/mux"
)

type UserHandlerFunc func(http.ResponseWriter, *http.Request)

type IUserHandler interface {
	GetUserById() UserHandlerFunc
	PostUser() UserHandlerFunc
	PasswordChange() UserHandlerFunc
	Create() UserHandlerFunc
	UserExist() UserHandlerFunc
	LogIn() UserHandlerFunc
}

type UserHandler struct {
	Router         *mux.Router
	userSqlService services.UserSqlService
	logger         *logger.Logger
}

func NewUserHandler(r *mux.Router, us services.UserSqlService, l *logger.Logger) IUserHandler {
	return UserHandler{r, us, l}
}

func (uh UserHandler) GetUserById() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["id"]
		uh.logger.Info().Msgf("Fetching user id %s", userId)
		intId, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			uh.logger.Err(err)
			responses.RespondBadRequest(w, "Wrong paramater Type. Required int")
			return
		}

		user, err := uh.userSqlService.GetUserById(intId)
		if err != nil {
			uh.logger.Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		if user == nil {
			responses.RespondNotFound(w)
			return
		}

		responses.NewUserResponse(user, w, r)
	}
}

func (uh UserHandler) PostUser() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh UserHandler) PasswordChange() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh UserHandler) Create() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh UserHandler) UserExist() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh UserHandler) LogIn() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
