package rpc_api

import (
	"context"
	"errors"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (us *UserServer) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := us.userService.GetById(ctx, req.UserId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, errors.New("user not found or activated").Error())
	}

	res := &pb.UserResponse{
		User: &pb.User{
			Id:        user.Id,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}
	return res, nil
}

func (us *UserServer) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerificationResponse, error) {
	err := us.userService.StoreVerificationData(ctx, req.GetCode())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.VerificationResponse{Success: true}, nil
}

func (us *UserServer) GetUserDetails(ctx context.Context, req *pb.GetUserRequest) (*pb.UserDetails, error) {
	user, err := us.userService.GetById(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, errors.New("user not found or activated").Error())
	}

	ud := &pb.UserDetails{}
	err = copier.Copy(ud, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return ud, nil
}

func (us *UserServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	err := us.userService.UpdateUserPassword(req.GetId(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.ChangePasswordResponse{Status: "ok"}, nil
}
