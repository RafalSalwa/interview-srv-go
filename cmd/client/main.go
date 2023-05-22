package main

import (
	"context"
	"fmt"

	"github.com/RafalSalwa/interview-app-srv/pkg/logger"

	"github.com/RafalSalwa/interview-app-srv/internal/generator"

	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	emailDomain = "@interview.com"
	password    = "VeryG00dPass!"
)

func main() {
	l := logger.NewConsole(false)
	pUsername, _ := generator.RandomString(8)
	username := *pUsername
	email := username + emailDomain
	fmt.Println("Client starting")
	ctx := context.TODO()

	conn, err := grpc.Dial("0.0.0.0:8082", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		l.Error().Err(err)
	}
	authClient := pb.NewAuthServiceClient(conn)
	userClient := pb.NewUserServiceClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			l.Error().Err(err)
		}
	}(conn)

	newUser := &pb.SignUpUserInput{
		Name:            username,
		Email:           email,
		Password:        password,
		PasswordConfirm: password,
	}
	signUp, err := authClient.SignUpUser(ctx, newUser)
	if err != nil {
		l.Error().Err(err)
	}
	l.Info().Msgf("Created User %s with id: %s", username, signUp.Id)

	rVerification := &pb.VerifyUserRequest{Code: signUp.VerificationToken}
	vs, err := userClient.VerifyUser(ctx, rVerification)
	if err != nil {
		l.Error().Err(err)
	}
	if vs.GetSuccess() {
		l.Info().Msgf("User %s verified", username)
	}

	credentials := &pb.SignInUserInput{
		Username: username,
		Password: password,
	}
	l.Info().Msg("SignIn previously created user")
	signIn, err := authClient.SignInUser(ctx, credentials)
	if err != nil {
		fmt.Println(err)
	}

	l.Info().Msgf("Username: %s", username)
	l.Info().Msgf("AccessToken: %s", signIn.AccessToken)
	l.Info().Msgf("RefreshToken: %s", signIn.RefreshToken)

	id := &pb.GetUserRequest{
		UserId: signUp.Id,
	}
	l.Info().Msgf("GET username: %s  details. UserId: %s", username, signUp.Id)
	user, err := userClient.GetUserById(ctx, id)
	if err != nil {
		l.Error().Err(err)
	}
	l.Info().Msgf("GetUser: %s", user.GetUser().GetId())
}
