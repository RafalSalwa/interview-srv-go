package rpc_api

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (userServer *UserServer) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	fmt.Printf("GetUserById: %s\n", req.GetUserId())
	id := req.GetUserId()
	intId, _ := strconv.ParseInt(id, 10, 64)
	user, err := userServer.userService.GetById(intId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, errors.New("user not found or activated").Error())
	}

	pbId := strconv.FormatInt(user.Id, 10)
	res := &pb.UserResponse{
		User: &pb.User{
			Id:        pbId,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}
	return res, nil
}

func (userServer *UserServer) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerificationResponse, error) {
	user, err := userServer.userService.GetByCode(req.GetCode())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, errors.New("user not found or activated").Error())
	}
	s := userServer.userService.Veryficate(user)
	if !s {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.VerificationResponse{Success: s}, nil
}
