package rpc_api

import (
	"context"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"

	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (authServer *AuthServer) SignInUser(ctx context.Context, req *pb.SignInUserInput) (*pb.SignInUserResponse, error) {
	loginUser := &models.LoginUserRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}

	ur, err := authServer.authService.Load(loginUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SignInUserResponse{
		Status:       "success",
		AccessToken:  ur.Token,
		RefreshToken: ur.RefreshToken,
	}
	return res, nil
}
