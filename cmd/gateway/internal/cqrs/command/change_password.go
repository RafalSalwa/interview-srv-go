package command

import (
	"context"
	"errors"
	"github.com/RafalSalwa/interview-app-srv/pkg/hashing"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type ChangePassword struct {
	Id              int64
	OldPassword     string
	Password        string
	PasswordConfirm string
}

type ChangePasswordHandler struct {
	grpcUser intrvproto.UserServiceClient
}

func NewChangePasswordHandler(grpcUser intrvproto.UserServiceClient) ChangePasswordHandler {
	return ChangePasswordHandler{grpcUser: grpcUser}
}

func (h ChangePasswordHandler) Handle(ctx context.Context, cmd ChangePassword) error {
	if !hashing.CheckPasswordHash(cmd.Password, cmd.OldPassword) {
		return errors.New("passwords are different")
	}
	passHash, err := hashing.HashPassword(cmd.Password)
	if err != nil {
		return err
	}
	_, err = h.grpcUser.ChangePassword(ctx, &intrvproto.ChangePasswordRequest{
		Id:       cmd.Id,
		Password: passHash,
	})
	if err != nil {
		return err
	}
	return nil
}
