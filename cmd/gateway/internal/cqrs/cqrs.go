package cqrs

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs/command"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs/query"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
}

func NewCQRSService(ctx context.Context, cfg *config.Config) (*Application, error) {
	conn, err := grpc.Dial(cfg.Grpc.AuthServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	authClient := intrvproto.NewAuthServiceClient(conn)
	conn, err = grpc.Dial(cfg.Grpc.UserServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userClient := intrvproto.NewUserServiceClient(conn)
	return newApplication(ctx, authClient, userClient), nil
}

func newApplication(ctx context.Context, authClient intrvproto.AuthServiceClient, userClient intrvproto.UserServiceClient) *Application {

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
		},
	}
}
