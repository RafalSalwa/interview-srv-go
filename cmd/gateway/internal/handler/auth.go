package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/RafalSalwa/interview-app-srv/pkg/tracing"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type AuthHandler interface {
	SignUpUser() HandlerFunc
	SignInUser() HandlerFunc
	ForgotPassword() HandlerFunc
	NewPassword() HandlerFunc
	Verify() HandlerFunc
	Login() HandlerFunc
	RefreshToken() HandlerFunc
	NewToken() HandlerFunc
}

type authHandler struct {
	Router     *mux.Router
	authClient intrvproto.AuthServiceClient
	logger     *logger.Logger
}

func NewAuthHandler(r *mux.Router, authClient intrvproto.AuthServiceClient, l *logger.Logger) AuthHandler {
	return authHandler{r, authClient, l}
}

func (a authHandler) SignUpUser() HandlerFunc {
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
		//ur, err := a.authClient.SignUpUser(signUpUser)
		//
		//if err != nil {
		//	a.logger.Error().Err(err).Msg("SignUpUser: create")
		//	responses.RespondInternalServerError(w)
		//	return
		//}
		//responses.NewUserResponse(ur, w)
	}
}

func (a authHandler) SignInUser() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracing.StartHttpServerTracerSpan(r, "authHandler:SignIn")
		defer span.Finish()

		decoder := json.NewDecoder(r.Body)
		signIn := &intrvproto.SignInUserInput{}

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
		pbUser, err := a.authClient.SignInUser(ctx, signIn)
		fmt.Println(err)
		if err != nil {
			span.SetTag("error", true)
			span.LogKV("error_code", err.Error())
			a.logger.Error().Err(err)
			responses.RespondInternalServerError(w)
			return
		}
		fmt.Println("pbUser", pbUser)
		ur := models.UserResponse{}
		err = ur.FromProtoSignIn(pbUser)
		if err != nil {
			return
		}
		responses.NewUserResponse(&ur, w)
	}
}

func (a authHandler) Verify() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vCode := mux.Vars(r)["code"]
		if vCode == "" {
			responses.RespondBadRequest(w, "code param is missing")
		}
		//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		//defer cancel()

		//if err := a.service.Verify(ctx, vCode); err != nil {
		//	responses.RespondInternalServerError(w)
		//	return
		//}
		//responses.RespondOk(w)
	}
}

func (a authHandler) ForgotPassword() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a authHandler) NewPassword() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a authHandler) RefreshToken() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a authHandler) NewToken() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a authHandler) Login() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		LoginReq := &models.LoginUserRequest{}

		err := json.NewDecoder(r.Body).Decode(LoginReq)
		if err != nil {
			a.logger.Error().Err(err).Msg("Decode Request")
			responses.RespondBadRequest(w, "")
			return
		}

		//user, err := a.service.SignInUser(LoginReq)
		//if err != nil {
		//	a.logger.Err(err)
		//	responses.RespondInternalServerError(w)
		//
		//	return
		//}
		//
		//if user == nil {
		//	// it's better to hide notFound in unauthorized response to avoid user guessing
		//	responses.RespondNotAuthorized(w, "")
		//	return
		//}
		//responses.NewUserResponse(user, w)
	}
}

func (a authHandler) Logout() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a authHandler) Token() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
