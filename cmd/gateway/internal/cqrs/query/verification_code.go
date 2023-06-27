package query

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type VerificationCodeHandler struct {
	authClient intrvproto.AuthServiceClient
}
type VerificationCode struct {
	Email string
}

func NewVerificationCodeHandler(authClient intrvproto.AuthServiceClient) VerificationCodeHandler {
	return VerificationCodeHandler{authClient: authClient}
}

func (h VerificationCodeHandler) Handle(ctx context.Context, q VerificationCode) (models.UserResponse, error) {
	email := &intrvproto.VerificationCodeRequest{
		Email: q.Email,
	}
	resp, err := h.authClient.GetVerificationKey(ctx, email)
	if err != nil {
		return models.UserResponse{}, err
	}
	u := models.UserResponse{
		VerificationCode: resp.Code,
	}
	return u, nil
}
