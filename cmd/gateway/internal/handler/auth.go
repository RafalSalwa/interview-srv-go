package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs/command"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs/query"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	grpc_codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler interface {
	SignUpUser() HandlerFunc
	SignInUser() HandlerFunc
	ForgotPassword() HandlerFunc
	NewPassword() HandlerFunc
	Verify() HandlerFunc
	GetVerificationCode() HandlerFunc
	Login() HandlerFunc
	RefreshToken() HandlerFunc
	NewToken() HandlerFunc
}

type authHandler struct {
	Router *mux.Router
	cqrs   *cqrs.Application
	logger *logger.Logger
}

func NewAuthHandler(r *mux.Router, cqrs *cqrs.Application, l *logger.Logger) AuthHandler {
	return authHandler{r, cqrs, l}
}

func (a authHandler) SignUpUser() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.GetTracerProvider().Tracer("auth-handler").Start(r.Context(), "Handler SignUpUser")
		defer span.End()

		decoder := json.NewDecoder(r.Body)
		signUpReq := models.CreateUserRequest{}

		if err := decoder.Decode(&signUpReq); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			a.logger.Error().Err(err).Msg("SignUpUser: decode")
			responses.RespondBadRequest(w, "wrong input parameters")
			return
		}

		validate := validator.New()
		if err := validate.StructCtx(ctx, signUpReq); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			a.logger.Error().Err(err)
			responses.RespondBadRequest(w, err.Error())
			return
		}
		err := a.cqrs.Commands.SignUp.Handle(ctx, command.SignUpUser{User: signUpReq})
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			a.logger.Error().Err(err).Msg("SignUpUser:create")
			if e, ok := status.FromError(err); ok {
				if e.Code() == grpc_codes.AlreadyExists {
					responses.RespondConflict(w, e.Message())
					return
				}
				responses.RespondBadRequest(w, e.Message())
				return
			}
			responses.RespondBadRequest(w, err.Error())
			return
		}
		responses.RespondCreated(w)
	}
}

func (a authHandler) GetVerificationCode() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		decoder := json.NewDecoder(r.Body)
		signIn := models.LoginUserRequest{}

		if err := decoder.Decode(&signIn); err != nil {
			a.logger.Error().Err(err).Msg("VerificationCode:decode")
			responses.RespondBadRequest(w, "wrong parameters")
			return
		}

		validate := validator.New()
		if err := validate.Struct(signIn); err != nil {
			a.logger.Error().Err(err)
			responses.RespondBadRequest(w, err.Error())
			return
		}

		user, err := a.cqrs.Queries.FetchUser.Handle(ctx, query.FetchUser{User: signIn})
		if err != nil {
			a.logger.Error().Err(err).Msg("gRPC:VerificationCode:GetUser")
		}

		if user.Id == 0 {
			responses.RespondNotFound(w)
			return
		}

		uVerification, err := a.cqrs.Queries.VerificationCode.Handle(ctx, query.VerificationCode{Email: signIn.Email})
		if err != nil {
			a.logger.Error().Err(err).Msg("gRPC:VerificationCode")
			if e, ok := status.FromError(err); ok {
				if e.Code() == grpc_codes.NotFound {
					responses.RespondNotFound(w)
					return
				}

				if e.Code() == grpc_codes.AlreadyExists {
					responses.RespondConflict(w, "user already activated")
					return
				}

				responses.RespondBadRequest(w, e.Message())
				return
			}

			responses.RespondInternalServerError(w)
			return
		}

		responses.NewUserResponse(uVerification, w)
	}
}

func (a authHandler) SignInUser() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.GetTracerProvider().Tracer("auth-handler").Start(r.Context(), "Handler SignUpUser")
		defer span.End()

		decoder := json.NewDecoder(r.Body)
		signIn := models.LoginUserRequest{}

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

		u, err := a.cqrs.Queries.SignIn.Handle(ctx, query.SignInUser{User: signIn})
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			a.logger.Error().Err(err).Msg("gRPC:SignIn")
			if e, ok := status.FromError(err); ok {
				if e.Code() == grpc_codes.NotFound {
					responses.RespondNotFound(w)
					return
				}
				responses.RespondBadRequest(w, e.Message())
				return
			}
			responses.RespondInternalServerError(w)
			return
		}
		responses.NewUserResponse(&u, w)
	}
}

func (a authHandler) Verify() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vCode := mux.Vars(r)["code"]
		if vCode == "" {
			responses.RespondBadRequest(w, "code param is missing")
		}
		err := a.cqrs.Commands.Verify.Handle(r.Context(), command.VerifyCode{VerificationCode: vCode})
		if err != nil {
			a.logger.Error().Err(err).Msg("gRPC:Verify")
			if e, ok := status.FromError(err); ok {
				if e.Code() == grpc_codes.NotFound {
					responses.RespondNotFound(w)
					return
				}
				if e.Code() == grpc_codes.AlreadyExists {
					responses.RespondConflict(w, "user already activated")
					return
				}
				responses.RespondBadRequest(w, e.Message())
				return
			}
			responses.RespondInternalServerError(w)
			return
		}
		responses.RespondOk(w)
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
