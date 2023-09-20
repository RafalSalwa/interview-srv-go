package query

import (
	"context"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type FetchUserHandler struct {
	userClient intrvproto.UserServiceClient
}
type FetchUser struct {
	User models.SignInUserRequest
}

func NewFetchUserHandler(userClient intrvproto.UserServiceClient) FetchUserHandler {
	return FetchUserHandler{userClient: userClient}
}

func (h FetchUserHandler) Handle(ctx context.Context, q FetchUser) (models.UserResponse, error) {
	credentials := &intrvproto.GetUserSignInRequest{
		Username: q.User.Username,
		Password: q.User.Password,
	}
	resp, err := h.userClient.GetUser(ctx, credentials)
	if err != nil {
		return models.UserResponse{}, err
	}
	u := models.UserResponse{
		Username: q.User.Username,
	}
	u.FromProtoUserDetails(resp)
	return u, nil
}
