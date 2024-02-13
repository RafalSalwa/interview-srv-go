package query

import (
	"context"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type GetUserHandler struct {
	grpcUser intrvproto.UserServiceClient
}

func NewGetUserHandler(userClient intrvproto.UserServiceClient) GetUserHandler {
	return GetUserHandler{grpcUser: userClient}
}

func (h GetUserHandler) Handle(ctx context.Context, id int64) (models.UserResponse, error) {
	req := &intrvproto.GetUserRequest{UserId: id}
	pu, err := h.grpcUser.GetUserById(ctx, req)
	ur := models.UserResponse{}

	if err != nil {
		return ur, err
	}

	ur.FromProtoUserDetails(pu)

	return ur, nil
}
