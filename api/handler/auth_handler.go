package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/RafalSalwa/interview-app-srv/internal/services"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

type IAuthHandler interface {
	SignUpUser() HandlerFunc
	SignInUser() HandlerFunc
	ForgotPassword() HandlerFunc
	NewPassword() HandlerFunc
	Verify() HandlerFunc
	Login() HandlerFunc
	RefreshToken() HandlerFunc
	NewToken() HandlerFunc
}

type AuthHandler struct {
	Router  *mux.Router
	service services.AuthService
	logger  *logger.Logger
}

func NewAuthHandler(r *mux.Router, as services.AuthService, l *logger.Logger) IAuthHandler {
	return AuthHandler{r, as, l}
}

func (a AuthHandler) SignUpUser() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		signUpUser := &models.CreateUserRequest{}

		if err := decoder.Decode(&signUpUser); err != nil {
			fmt.Println(signUpUser)
			a.logger.Error().Err(err).Msg("SignUpUser: decode")
			responses.RespondBadRequest(w, "wrong input parameters")
			return
		}

		validate := validator.New()
		if err := validate.Struct(signUpUser); err != nil {
			a.logger.Error().Err(err)
			responses.RespondBadRequest(w, err.Error())
			return
		}

		ur, err := a.service.SignUpUser(signUpUser)

		if err != nil {
			a.logger.Error().Err(err).Msg("SignUpUser: create")
			responses.RespondInternalServerError(w)
			return
		}

		responses.NewUserResponse(ur, w)
	}
}

func (a AuthHandler) SignInUser() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		signIn := &models.LoginUserRequest{}

		if err := decoder.Decode(&signIn); err != nil {
			a.logger.Error().Err(err).Msg("SignInUser: decode")
			responses.RespondBadRequest(w, "wrong parameters")
			return
		}

		validate := validator.New()
		if err := validate.Struct(signIn); err != nil {
			a.logger.Error().Err(err)
			responses.RespondBadRequest(w, err.Error())
			return
		}
		ur, err := a.service.SignInUser(signIn)
		if err != nil {
			a.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}
		responses.NewUserResponse(ur, w)
	}
}

func (a AuthHandler) Verify() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vCode := mux.Vars(r)["code"]
		if vCode == "" {
			responses.RespondBadRequest(w, "code param is missing")
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if err := a.service.Verify(ctx, vCode); err != nil {
			responses.RespondInternalServerError(w)
			return
		}
		responses.RespondOk(w)
	}
}

func (a AuthHandler) ForgotPassword() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a AuthHandler) NewPassword() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a AuthHandler) RefreshToken() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a AuthHandler) NewToken() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a AuthHandler) Login() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		LoginReq := &models.LoginUserRequest{}

		err := json.NewDecoder(r.Body).Decode(LoginReq)
		if err != nil {
			a.logger.Error().Err(err).Msg("Decode Request")
			responses.RespondBadRequest(w, "")
			return
		}

		user, err := a.service.SignInUser(LoginReq)
		if err != nil {
			a.logger.Err(err)
			responses.RespondInternalServerError(w)

			return
		}

		if user == nil {
			// it's better to hide notFound in unauthorized response to avoid user guessing
			responses.RespondNotAuthorized(w, "")
			return
		}
		responses.NewUserResponse(user, w)
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
