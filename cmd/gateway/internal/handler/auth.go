package handler

import (
    "encoding/json"
    "github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs"
    "github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs/command"
    "github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs/query"
    _ "github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/router"
    "github.com/RafalSalwa/interview-app-srv/pkg/http/auth"
    "github.com/RafalSalwa/interview-app-srv/pkg/logger"
    "github.com/RafalSalwa/interview-app-srv/pkg/models"
    "github.com/RafalSalwa/interview-app-srv/pkg/responses"
    "github.com/go-playground/validator/v10"
    "github.com/gorilla/mux"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/codes"
    "go.opentelemetry.io/otel/trace"
    "google.golang.org/grpc/status"
    "net/http"
)

type AuthHandler interface {
	RegisterRoutes(r *mux.Router, cfg auth.Auth) error

	SignUpUser() http.HandlerFunc
	SignInUser() http.HandlerFunc

	Verify() http.HandlerFunc
	GetVerificationCode() http.HandlerFunc
}

type authHandler struct {
	cqrs   *cqrs.Application
	logger *logger.Logger
}

func (a authHandler) RegisterRoutes(r *mux.Router, cfg auth.Auth) error {
	authorizer, err := auth.NewAuthorizer(cfg)
	if err != nil {
		return err
	}

	sr := r.PathPrefix("/auth/").Subrouter()

	sr.Methods(http.MethodPost).Path("/signup").HandlerFunc(authorizer.Middleware(a.SignUpUser()))
	sr.Methods(http.MethodPost).Path("/signin").HandlerFunc(authorizer.Middleware(a.SignInUser()))

	sr.Methods(http.MethodGet).Path("/verify/{code}").HandlerFunc(authorizer.Middleware(a.Verify()))
	sr.Methods(http.MethodPost).Path("/code").HandlerFunc(authorizer.Middleware(a.GetVerificationCode()))
	return nil
}

func NewAuthHandler(cqrs *cqrs.Application, l *logger.Logger) AuthHandler {
	return authHandler{cqrs, l}
}

func (a authHandler) SignInUser() http.HandlerFunc {
	req := models.SignInUserRequest{}
	reqValidator := validator.New()

	res := models.UserResponse{}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.GetTracerProvider().Tracer("auth-handler").Start(r.Context(), "Handler SignUpUser")
		defer span.End()

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			span.RecordError(err, trace.WithStackTrace(true))
			span.SetStatus(codes.Error, err.Error())
			a.logger.Error().Err(err).Msg("SignInUser: decode")

			responses.RespondBadRequest(w, "wrong parameters")
			return
		}

		if err := reqValidator.Struct(req); err != nil {
			span.RecordError(err, trace.WithStackTrace(true))
			span.SetStatus(codes.Error, err.Error())
			a.logger.Error().Err(err).Msg("SignInUser: validate")

			responses.RespondBadRequest(w, err.Error())
			return
		}

		var errQuery error
		res, errQuery = a.cqrs.Queries.SignIn.Handle(ctx, query.SignInUser{User: req})

		if errQuery != nil {
			span.RecordError(errQuery)
			span.SetStatus(codes.Error, errQuery.Error())
			a.logger.Error().Err(errQuery).Msg("gRPC:SignIn")

			if e, ok := status.FromError(errQuery); ok {
				responses.FromGRPCError(e, w)
				return
			}
			responses.RespondInternalServerError(w)
			return
		}
		responses.NewUserResponse(&res, w)
	}
}

func (a authHandler) SignUpUser() http.HandlerFunc {
	req := models.SignUpUserRequest{}
	reqValidator := validator.New()

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.GetTracerProvider().Tracer("auth-handler").Start(r.Context(), "Handler SignUpUser")
		defer span.End()

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			span.RecordError(err, trace.WithStackTrace(true))
			span.SetStatus(codes.Error, err.Error())
			a.logger.Error().Err(err).Msg("SignUpUser: decode")

			responses.RespondBadRequest(w, "wrong input parameters")
			return
		}

		if err := reqValidator.StructCtx(ctx, req); err != nil {
			span.RecordError(err, trace.WithStackTrace(true))
			span.SetStatus(codes.Error, err.Error())
			a.logger.Error().Err(err).Msg("SignUpUser: validate")

			responses.RespondBadRequest(w, err.Error())
			return
		}

		err := a.cqrs.Commands.SignUp.Handle(ctx, command.SignUpUser{User: req})

		if err != nil {
			span.RecordError(err, trace.WithStackTrace(true))
			span.SetStatus(codes.Error, err.Error())
			a.logger.Error().Err(err).Msg("SignUpUser:create")

			if e, ok := status.FromError(err); ok {
				responses.FromGRPCError(e, w)
				return
			}
			responses.RespondBadRequest(w, err.Error())
			return
		}
		responses.RespondCreated(w)
	}
}

func (a authHandler) GetVerificationCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		decoder := json.NewDecoder(r.Body)
		signIn := models.SignInUserRequest{}

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
				responses.FromGRPCError(e, w)
				return
			}
			responses.RespondInternalServerError(w)
			return
		}
		responses.NewUserResponse(uVerification, w)
	}
}

func (a authHandler) Verify() http.HandlerFunc {
	var req string
	return func(w http.ResponseWriter, r *http.Request) {
		req = mux.Vars(r)["code"]
		if req == "" {
			responses.RespondBadRequest(w, "code param is missing")
		}

		err := a.cqrs.Commands.Verify.Handle(r.Context(), command.VerifyCode{VerificationCode: req})

		if err != nil {
			a.logger.Error().Err(err).Msg("gRPC:Verify")
			if e, ok := status.FromError(err); ok {
				responses.FromGRPCError(e, w)
				return
			}
			responses.RespondInternalServerError(w)
			return
		}
		responses.RespondOk(w)
	}
}
