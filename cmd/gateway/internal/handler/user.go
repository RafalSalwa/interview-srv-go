package handler

import (
	"encoding/json"
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs/command"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs/query"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler interface {
	GetUserById() HandlerFunc
	PasswordChange() HandlerFunc
	ValidateCode() HandlerFunc
}

type userHandler struct {
	router *mux.Router
	cqrs   *cqrs.Application
	logger *logger.Logger
}

func NewUserHandler(r *mux.Router, cqrs *cqrs.Application, l *logger.Logger) UserHandler {
	return userHandler{r, cqrs, l}
}

func (uh userHandler) GetUserById() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userId, _ := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
		user, err := uh.cqrs.Queries.UserBasic.Handle(ctx, query.UserRequest{UserId: userId})
		if err != nil {
			uh.logger.Error().Err(err).Msg("rest:handler:getUser")
			responses.RespondInternalServerError(w)
			return
		}
		if user == nil {
			responses.RespondNotFound(w)
			return
		}

		responses.NewUserResponse(user, w)
	}
}

func (uh userHandler) PasswordChange() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pcr := &models.ChangePasswordRequest{}

		if err := json.NewDecoder(r.Body).Decode(pcr); err != nil {
			uh.logger.Error().Err(err).Msg("Decode")
			responses.RespondBadRequest(w, "Invalid JSON")
			return
		}

		//if err := password.Validate(pcr.Password, pcr.PasswordConfirm); err != nil {
		//	uh.logger.Error().Err(err).Msg("Validate")
		//	responses.RespondBadRequest(w, err.Error())
		//	return
		//}
		user, err := uh.cqrs.Queries.UserDetails.Handle(ctx, query.UserRequest{UserId: pcr.Id})
		if err != nil {
			uh.logger.Error().Err(err).Msg("cqrs:user:details:get")
			responses.RespondInternalServerError(w)
			return
		}

		if user == nil {
			responses.RespondNotFound(w)
			return
		}
		err = uh.cqrs.Commands.ChangePassword.Handle(ctx, command.ChangePassword{
			Id:              pcr.Id,
			OldPassword:     user.Password,
			Password:        pcr.Password,
			PasswordConfirm: pcr.PasswordConfirm,
		})
		if err != nil {
			uh.logger.Error().Err(err).Msg("cqrs:changePassword")
			responses.RespondInternalServerError(w)
			return
		}

		responses.RespondOk(w)
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
