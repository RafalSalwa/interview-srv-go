package cqrs

import (
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs/command"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs/query"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/rpc_client"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SignUp         command.SignUpHandler
	ChangePassword command.ChangePasswordHandler
	Verify         command.VerifyCodeHandler
}

type Queries struct {
	UserBasic        query.UserBasicHandler
	UserDetails      query.UserDetailsHandler
	SignIn           query.SignInHandler
	VerificationCode query.VerificationCodeHandler
	FetchUser        query.FetchUserHandler
}

func NewCQRSService(cfg config.Grpc) (*Application, error) {
	authClient, err := rpc_client.NewAuthClient(cfg.AuthServicePort)
	if err != nil {
		return nil, err
	}

	userClient, err := rpc_client.NewUserClient(cfg.UserServicePort)
	if err != nil {
		return nil, err
	}

	return newApplication(authClient, userClient), nil
}

func newApplication(authClient intrvproto.AuthServiceClient, userClient intrvproto.UserServiceClient) *Application {

	return &Application{
		Commands: Commands{
			SignUp:         command.NewSignUpHandler(authClient),
			ChangePassword: command.NewChangePasswordHandler(userClient),
			Verify:         command.NewVerifyCodeHandler(userClient),
		},
		Queries: Queries{
			SignIn:           query.NewSignInHandler(authClient),
			UserDetails:      query.NewUserDetailsHandler(userClient),
			UserBasic:        query.NewUserBasicHandler(userClient),
			VerificationCode: query.NewVerificationCodeHandler(authClient),
			FetchUser:        query.NewFetchUserHandler(userClient),
		},
	}
}
