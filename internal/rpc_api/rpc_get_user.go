package rpc_api

import (
	"context"
	"strconv"

	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (userServer *UserServer) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	id := req.GetUserId()
	intId, _ := strconv.ParseInt(id, 10, 64)
	user, err := userServer.userService.GetById(intId)

	if err != nil {
		return nil, err
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
