package query

import (
	"context"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type SignInHandler struct {
	authClient intrvproto.AuthServiceClient
}
type SignInUser struct {
	User models.SignInUserRequest
}

func NewSignInHandler(authClient intrvproto.AuthServiceClient) SignInHandler {
	return SignInHandler{authClient: authClient}
}

func (h SignInHandler) Handle(ctx context.Context, q SignInUser) (models.UserResponse, error) {
	credentials := &intrvproto.SignInUserInput{
		Username: q.User.Username,
		Password: q.User.Password,
	}
	resp, err := h.authClient.SignInUser(ctx, credentials)
	if err != nil {
		return models.UserResponse{}, err
	}
	u := models.UserResponse{
		Username: q.User.Username,
	}
	u.FromProtoSignIn(resp)
	return u, nil
}
