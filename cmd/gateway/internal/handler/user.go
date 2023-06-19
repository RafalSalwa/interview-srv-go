package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/internal/password"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/go-playground/validator/v10"

	"github.com/gorilla/mux"
)

type UserHandler interface {
	GetUserById() HandlerFunc
	CreateUser() HandlerFunc
	PasswordChange() HandlerFunc
	ValidateCode() HandlerFunc
}

type userHandler struct {
	router *mux.Router
	logger *logger.Logger
}

func NewUserHandler(r *mux.Router, l *logger.Logger) UserHandler {
	return userHandler{r, l}
}

func (uh userHandler) GetUserById() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//userId, err := strconv.Atoi(mux.Vars(r)["id"])
		//if err != nil {
		//	uh.logger.Err(err)
		//	responses.RespondBadRequest(w, "Wrong parameter Type. Required int")
		//	return
		//}
		responses.Respond(w, 200, []byte("123"))
		//user, err := uh.service.GetById(userId)
		//if err != nil {
		//	uh.logger.Err(err)
		//	responses.RespondInternalServerError(w)
		//	return
		//}
		//
		//if user == nil {
		//	responses.RespondNotFound(w)
		//	return
		//}
		//
		//ur := &models.UserResponse{}
		//if err = ur.FromDBResponse(user); err != nil {
		//	uh.logger.Err(err)
		//	responses.RespondInternalServerError(w)
		//	return
		//}
		//
		//responses.NewUserResponse(ur, w)
	}
}

func (uh userHandler) CreateUser() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newUserRequest := &models.CreateUserRequest{}
		if err := json.NewDecoder(r.Body).Decode(newUserRequest); err != nil {
			uh.logger.Error().Err(err)
			responses.RespondBadRequest(w, "Invalid JSON request")
			return
		}

		validate := validator.New()
		if err := validate.Struct(newUserRequest); err != nil {
			uh.logger.Error().Err(err)
			responses.RespondBadRequest(w, err.Error())
			return
		}

		//newUser, err := uh.service.CreateUser(newUserRequest)
		//if err != nil {
		//	uh.logger.Error().Err(err)
		//	responses.RespondInternalServerError(w)
		//	return
		//}
		//
		//responses.NewUserResponse(newUser, w)
	}
}

func (uh userHandler) PasswordChange() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pcr := &models.ChangePasswordRequest{}

		if err := json.NewDecoder(r.Body).Decode(pcr); err != nil {
			uh.logger.Error().Err(err).Msg("Decode")
			responses.RespondBadRequest(w, "Invalid JSON")
			return
		}

		if err := password.Validate(pcr.Password, pcr.PasswordConfirm); err != nil {
			uh.logger.Error().Err(err).Msg("Validate")
			responses.RespondBadRequest(w, err.Error())
			return
		}

		//user, err := uh.service.GetById(pcr.Id)
		//if err != nil {
		//	uh.logger.Error().Err(err)
		//	responses.RespondInternalServerError(w)
		//	return
		//}
		//
		//if user == nil {
		//	responses.RespondNotFound(w)
		//	return
		//}
		//
		//if !password.CheckPasswordHash(pcr.Password, user.Password) {
		//	uh.logger.Error().Msg("New Password cannot be the same as old one")
		//	responses.RespondBadRequest(w, err.Error())
		//	return
		//}
		//
		//passHash, err := password.HashPassword(pcr.Password)
		//if err != nil {
		//	uh.logger.Error().Err(err).Msg("New Password cannot be the same as old one")
		//	responses.RespondInternalServerError(w)
		//	return
		//}
		//updateUser := &models.UpdateUserRequest{
		//	Id:       pcr.Id,
		//	Password: &passHash,
		//}
		//
		//if errUpdate := uh.service.UpdateUserPassword(updateUser); errUpdate != nil {
		//	uh.logger.Error().Err(errUpdate)
		//	responses.RespondInternalServerError(w)
		//	return
		//}
		//responses.RespondOk(w) // Instead we can also send redirect to login page
	}
}

func (uh userHandler) ValidateCode() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//code := mux.Vars(r)["code"]
		//um := models.UserDBModel{VerificationCode: code}
		//status := uh.service.StoreVerificationData(&um)
		//if !status {
		//	responses.RespondInternalServerError(w)
		//}
		//responses.RespondOk(w)
	}
}
