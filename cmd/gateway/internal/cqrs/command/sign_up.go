package command

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type SignUpUser struct {
	User models.CreateUserRequest
}

type SignUpHandler struct {
	authClient intrvproto.AuthServiceClient
}

func NewSignUpHandler(authClient intrvproto.AuthServiceClient) SignUpHandler {
	return SignUpHandler{authClient: authClient}
}

func (h SignUpHandler) Handle(ctx context.Context, cmd SignUpUser) error {
	_, err := h.authClient.SignUpUser(ctx, &intrvproto.SignUpUserInput{
		Name:            cmd.User.Username,
		Email:           cmd.User.Email,
		Password:        cmd.User.Password,
		PasswordConfirm: cmd.User.PasswordConfirm,
	})
	if err != nil {
		return err
	}
	return nil
}
