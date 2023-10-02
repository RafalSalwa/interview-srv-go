package rpc

import (
	"context"
	"errors"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"github.com/jinzhu/copier"
	"go.opentelemetry.io/otel"
	otelcodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (us *UserServer) CheckUserExists(ctx context.Context, req *pb.StringValue) (*pb.BoolValue, error) {
	_, span := otel.GetTracerProvider().Tracer("user_service-rpc").Start(ctx, "GRPC GetUserByID")
	defer span.End()

	user := &models.UserDBModel{Email: req.GetValue()}
	exists, err := us.userService.UsernameInUse(user)
	if err != nil {
		return &pb.BoolValue{Value: false}, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.BoolValue{Value: exists}, nil
}

func (us *UserServer) GetUserByID(ctx context.Context, req *pb.GetUserRequest) (*pb.UserDetails, error) {
	ctx, span := otel.GetTracerProvider().Tracer("user_service-rpc").Start(ctx, "GRPC GetUserByID")
	defer span.End()

	udb, err := us.userService.GetByID(ctx, req.UserId)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelcodes.Error, err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if udb == nil {
		span.RecordError(err)
		span.SetStatus(otelcodes.Error, err.Error())
		return nil, status.Errorf(codes.NotFound, errors.New("user not found or activated").Error())
	}
	res := &pb.UserDetails{
		Id:        udb.Id,
		Username:  udb.Username,
		Firstname: udb.Firstname,
		Lastname:  udb.Lastname,
		Email:     udb.Email,
		Verified:  udb.Verified,
		Active:    udb.Active,
		CreatedAt: timestamppb.New(udb.CreatedAt),
	}

	return res, nil
}

func (us *UserServer) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerificationResponse, error) {
	err := us.userService.StoreVerificationData(ctx, req.GetCode())
	if err != nil {
		if err.Error() == "NotFound" {
			return nil, status.Errorf(codes.NotFound, "verification code not found")
		}
		if err.Error() == "AlreadyActivated" {
			return nil, status.Errorf(codes.AlreadyExists, "User with such code has already active account")
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.VerificationResponse{Success: true}, nil
}

func (us *UserServer) GetUser(ctx context.Context, req *pb.GetUserSignInRequest) (*pb.UserDetails, error) {
	reqUser := &models.SignInUserRequest{}
	reqUser.Username = req.Username
	reqUser.Email = req.Email

	user, err := us.userService.GetUser(ctx, reqUser)
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

func (us *UserServer) GetUserDetails(ctx context.Context, req *pb.GetUserRequest) (*pb.UserDetails, error) {
	user, err := us.userService.GetByID(ctx, req.GetUserId())
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

func (us *UserServer) ChangePassword(
	ctx context.Context,
	req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	err := us.userService.UpdateUserPassword(ctx, req.GetId(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.ChangePasswordResponse{Status: "ok"}, nil
}
