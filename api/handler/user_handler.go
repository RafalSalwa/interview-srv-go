package handler

import (
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"net/http"
	"strconv"

	"github.com/RafalSalwa/interview-app-srv/services"
	"github.com/gorilla/mux"
)

type UserHandlerFunc func(http.ResponseWriter, *http.Request)

type UserHandler interface {
	GetUserById() UserHandlerFunc
	PostUser() UserHandlerFunc
	PasswordChange() UserHandlerFunc
	Create() UserHandlerFunc
	UserExist() UserHandlerFunc
	LogIn() UserHandlerFunc
}

type userHandler struct {
	Router         *mux.Router
	userSqlService services.UserSqlService
	logger         *logger.Logger
}

func NewUserHandler(r *mux.Router, us services.UserSqlService, l *logger.Logger) UserHandler {
	return userHandler{r, us, l}
}

func (uh userHandler) GetUserById() UserHandlerFunc {
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
			fmt.Println(err)
			uh.logger.Err(err)
			responses.RespondInternalServerError(w)
			return
		}

		if user == nil {
			responses.RespondNotFound(w)
			return
		}
		ur := &models.UserResponse{
			Id:        user.Id,
			Username:  user.Username,
			Firstname: user.Firstname,
			RolesJson: user.RolesJson,
			Roles:     user.Roles,
			Verified:  user.Verified,
			Active:    user.Active,
			CreatedAt: user.CreatedAt,
		}
		responses.NewUserResponse(ur, w, r)
	}
}

func (uh userHandler) PostUser() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh userHandler) PasswordChange() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh userHandler) Create() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh userHandler) UserExist() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh userHandler) LogIn() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (u userHandler) Token() UserHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
