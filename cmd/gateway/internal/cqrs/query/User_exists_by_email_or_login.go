package query

import (
	"context"

	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type UserExistsHandler struct {
	userClient intrvproto.UserServiceClient
}

func NewUserExistsHandler(userClient intrvproto.UserServiceClient) UserExistsHandler {
	return UserExistsHandler{userClient: userClient}
}

func (h UserExistsHandler) Handle(ctx context.Context, email string) (bool, error) {
	req := &intrvproto.StringValue{Value: email}
	resp, err := h.userClient.CheckUserExists(ctx, req)
	if err != nil {
		return false, err
	}
	return resp.GetValue(), nil
}
