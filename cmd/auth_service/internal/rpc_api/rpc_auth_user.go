package rpc_api

import (
    "context"
    "github.com/RafalSalwa/interview-app-srv/pkg/models"
    pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
    "go.opentelemetry.io/otel"
    otelcodes "go.opentelemetry.io/otel/codes"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/timestamppb"
)

func (authServer *AuthServer) SignInUser(ctx context.Context, req *pb.SignInUserInput) (*pb.SignInUserResponse, error) {
    ctx, span := otel.GetTracerProvider().Tracer("auth_service-rpc").Start(ctx, "GRPC SignInUser")
    defer span.End()

    loginUser := &models.SignInUserRequest{
        Username: req.GetUsername(),
        Email:    req.GetEmail(),
        Password: req.GetPassword(),
    }

    ur, err := authServer.authService.SignInUser(loginUser)
    if err != nil {
        span.RecordError(err)
        span.SetStatus(otelcodes.Error, err.Error())
        authServer.logger.Error().Err(err).Msg("rpc:service:signin")
        
        if err.Error() == "record not found" {
            return nil, status.Errorf(codes.NotFound, err.Error())
        }
        return nil, status.Errorf(codes.Internal, err.Error())
    }

    res := &pb.SignInUserResponse{
        AccessToken:  ur.AccessToken,
        RefreshToken: ur.RefreshToken,
    }
    return res, nil
}

func (authServer *AuthServer) SignUpUser(ctx context.Context, req *pb.SignUpUserInput) (*pb.SignUpUserResponse, error) {
    ctx, span := otel.GetTracerProvider().Tracer("auth_service-rpc").Start(ctx, "GRPC SignUpUser")
    defer span.End()

    userSignUp := &models.SignUpUserRequest{
        Email:           req.GetEmail(),
        Password:        req.GetPassword(),
        PasswordConfirm: req.GetPasswordConfirm(),
    }

    um := &models.UserDBModel{}
    um.Email = req.Email
    um.Username = req.Email

    dbUser, _ := authServer.authService.Load(um)
    if dbUser != nil {
        return nil, status.Errorf(codes.AlreadyExists, "User with such credentials already exists")
    }

    ur, err := authServer.authService.SignUpUser(ctx, userSignUp)
    if err != nil {
        span.RecordError(err)
        span.SetStatus(otelcodes.Error, err.Error())
        authServer.logger.Error().Err(err).Msg("rpc:service:signup")
        return nil, status.Errorf(codes.Internal, err.Error())
    }
    res := &pb.SignUpUserResponse{
        Id:                ur.Id,
        Username:          ur.Username,
        VerificationToken: ur.VerificationCode,
        CreatedAt:         timestamppb.New(ur.CreatedAt),
    }
    return res, nil
}

func (authServer *AuthServer) GetVerificationKey(ctx context.Context, in *pb.VerificationCodeRequest) (*pb.VerificationCodeResponse, error) {
    ur, err := authServer.authService.GetVerificationKey(ctx, in.Email)
    if err != nil {
        authServer.logger.Error().Err(err).Msg("rpc:service:getkey")
        if err.Error() == "record not found" {
            return nil, status.Errorf(codes.NotFound, "user with such credentials was not found")
        }
        return nil, status.Errorf(codes.NotFound, err.Error())
    }
    return &pb.VerificationCodeResponse{Code: ur.VerificationCode}, nil
}
