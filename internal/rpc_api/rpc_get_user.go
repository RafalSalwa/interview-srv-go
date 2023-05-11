package rpc_api

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (userServer *UserServer) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	id := req.GetUserId()
	intId, _ := strconv.ParseInt(id, 10, 64)
	fmt.Println(intId)
	user, err := userServer.userService.GetById(intId)

	if err != nil {
		return nil, err
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
