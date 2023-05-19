package rpc_api

import (
	"context"
	"strconv"

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

func (authServer *AuthServer) SignUpUser(ctx context.Context, req *pb.SignUpUserInput) (*pb.SignUpUserResponse, error) {
	signUpUser := &models.CreateUserRequest{
		Username:        req.GetName(),
		Email:           req.GetEmail(),
		Password:        req.GetPassword(),
		PasswordConfirm: req.GetPasswordConfirm(),
	}

	ur, err := authServer.authService.SignUpUser(signUpUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SignUpUserResponse{
		Id:                strconv.FormatInt(ur.Id, 10),
		Username:          ur.Username,
		VerificationToken: ur.VerificationToken,
		CreatedAt:         nil,
	}
	return res, nil
}
func (authServer *AuthServer) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerificationResponse, error) {
	return nil, nil
}
