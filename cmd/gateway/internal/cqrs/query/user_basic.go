package query

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type UserBasic struct {
	UserId int64
}

type UserBasicHandler struct {
	grpcUser intrvproto.UserServiceClient
}

func NewUserBasicHandler(userClient intrvproto.UserServiceClient) UserBasicHandler {
	return UserBasicHandler{grpcUser: userClient}
}

func (h UserBasicHandler) Handle(ctx context.Context, query UserRequest) (*models.UserResponse, error) {
	req := &intrvproto.GetUserRequest{UserId: query.UserId}
	pu, err := h.grpcUser.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}

	ur := &models.UserResponse{}
	err = ur.FromProtoUserResponse(pu)
	if err != nil {
		return nil, err
	}

	return ur, nil
}
