package rpc_api

import (
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/internal/services"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	config      config.ConfGRPC
	userService services.UserSqlService
}

func NewGrpcUserServer(config config.ConfGRPC, userService services.UserSqlService) (*UserServer, error) {
	userServer := &UserServer{
		config:      config,
		userService: userService,
	}

	return userServer, nil
}
