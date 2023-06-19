package rpc_api

import (
	"github.com/RafalSalwa/interview-app-srv/internal/services"
	"github.com/RafalSalwa/interview-app-srv/pkg/grpc"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	config      grpc.Config
	userService services.UserSqlService
}

func NewGrpcUserServer(config grpc.Config, userService services.UserSqlService) (*UserServer, error) {
	userServer := &UserServer{
		config:      config,
		userService: userService,
	}

	return userServer, nil
}
