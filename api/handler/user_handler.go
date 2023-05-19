package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/internal/mapper"
	"github.com/RafalSalwa/interview-app-srv/internal/password"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"

	"github.com/RafalSalwa/interview-app-srv/internal/services"
	"github.com/gorilla/mux"
)

type UserHandler interface {
	GetUserById() HandlerFunc
	PostUser() HandlerFunc
	PasswordChange() HandlerFunc
}

type userHandler struct {
	Router         *mux.Router
	userSqlService services.UserSqlService
	logger         *logger.Logger
}

func NewUserHandler(r *mux.Router, us services.UserSqlService, l *logger.Logger) UserHandler {
	return userHandler{r, us, l}
}

func (uh userHandler) GetUserById() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["id"]
		uh.logger.Info().Msgf("Fetching user id %s", userId)
		intId, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			uh.logger.Err(err)
			responses.RespondBadRequest(w, "Wrong paramater Type. Required int")
			return
		}

		user, err := uh.userSqlService.GetById(intId)
		if err != nil {
			uh.logger.Err(err)
			err := responses.RespondInternalServerError(w)
			if err != nil {
				return
			}
			return
		}

		if user == nil {
			responses.RespondNotFound(w)
			return
		}
		ur := mapper.MapUserDBResponseToUserResponse(user)

		responses.NewUserResponse(ur, w)
	}
}

func (uh userHandler) PostUser() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newUserRequest := &models.CreateUserRequest{}
		err := json.NewDecoder(r.Body).Decode(newUserRequest)

		if err != nil {
			uh.logger.Error().Err(err)
			responses.RespondBadRequest(w, "")
			return
		}

		err = password.Validate(newUserRequest.Password, newUserRequest.PasswordConfirm)
		if err != nil {
			uh.logger.Error().Err(err)
			responses.RespondBadRequest(w, err.Error())
			return
		}

		if uh.userSqlService.Exists(newUserRequest) {
			responses.RespondConflict(w, "username or email already exists")
			return
		}
		u, err := uh.userSqlService.CreateUser(newUserRequest)
		if err != nil {
			uh.logger.Error().Err(err)
			err := responses.RespondInternalServerError(w)
			if err != nil {
				return
			}
			return
		}
		responses.NewUserResponse(u, w)
	}
}

func (uh userHandler) PasswordChange() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		passChange := &models.ChangePasswordRequest{}
		err := json.NewDecoder(r.Body).Decode(passChange)
		if err != nil {
			uh.logger.Error().Err(err).Msg("Decode")
			responses.RespondBadRequest(w, "")
			return
		}
		err = password.Validate(passChange.Password, passChange.PasswordConfirm)
		if err != nil {
			uh.logger.Error().Err(err).Msg("Validate")
			responses.RespondBadRequest(w, err.Error())
			return
		}

		user, err := uh.userSqlService.GetById(passChange.Id)
		if err != nil {
			uh.logger.Err(err)
			err := responses.RespondInternalServerError(w)
			if err != nil {
				return
			}
			return
		}

		if user == nil {
			responses.RespondNotFound(w)
			return
		}
		if !password.CheckPasswordHash(passChange.Password, user.Password) {
			uh.logger.Error().Msg("Password cannot be the same as old one")
			responses.RespondBadRequest(w, err.Error())
			return
		}
		passHash, _ := password.HashPassword(passChange.Password)

		updateUser := &models.UpdateUserRequest{
			Id:       passChange.Id,
			Password: &passHash,
		}
		err = uh.userSqlService.UpdateUserPassword(updateUser)
		if err != nil {
			uh.logger.Err(err)
			err := responses.RespondInternalServerError(w)
			if err != nil {
				return
			}
			return
		}
		responses.RespondOk(w)
	}
}
