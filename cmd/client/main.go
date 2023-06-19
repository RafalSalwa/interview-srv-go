package main

import (
	"context"
	"fmt"
	"time"

	"github.com/RafalSalwa/interview-app-srv/pkg/logger"

	"github.com/RafalSalwa/interview-app-srv/internal/generator"

	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	emailDomain = "@interview.com"
	password    = "VeryG00dPass!"
	numUsers    = 500
)

type User struct {
	Id             string
	ValidationCode string
	Username       string
	Password       string
	token          Token
}

type Token struct {
	access  string
	refresh string
}

var (
	maxNbConcurrentGoroutines = 100
	concurrentGoroutines      = make(chan struct{}, maxNbConcurrentGoroutines)
)

func runWorkersInOrder() {
	l := logger.NewConsole(true)
	conn, err := grpc.Dial("0.0.0.0:8089", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		l.Error().Err(err)
	}
	authClient := pb.NewAuthServiceClient(conn)
	userClient := pb.NewUserServiceClient(conn)

	qCreatedUsers := make(chan User, numUsers)
	qActivatedUsers := make(chan User, numUsers)
	qFailedUsers := make(chan User, numUsers)

	done := make(chan bool)

	for i := 1; i <= numUsers; i++ {
		go createUser(authClient, qCreatedUsers, qFailedUsers)
	}
	for i := 1; i <= numUsers; i++ {
		go activateUser(userClient, qCreatedUsers, qActivatedUsers, qFailedUsers)
	}
	for i := 1; i <= numUsers; i++ {
		go tokenUser(authClient, qActivatedUsers, qFailedUsers)
	}

	go func() {
		for {
			fmt.Printf("Concurrent queue len: | %6d | user creation queue:  %6d | user activation queue: %6d \n", len(concurrentGoroutines), len(qCreatedUsers), len(qActivatedUsers))
			if len(concurrentGoroutines) == 0 {
				done <- true
				fmt.Println("Queues depleted, closing")
			}
			time.Sleep(time.Second)
		}
	}()

	<-done
}

func main() {
	fmt.Println("starting client")
	runWorkersInOrder()
}

func createUser(authClient pb.AuthServiceClient, created chan User, failed chan User) {
	concurrentGoroutines <- struct{}{}
	ctx := context.TODO()

	pUsername, _ := generator.RandomString(12)
	email := *pUsername + emailDomain

	newUser := &pb.SignUpUserInput{
		Name:            *pUsername,
		Email:           email,
		Password:        password,
		PasswordConfirm: password,
	}

	signUp, err := authClient.SignUpUser(ctx, newUser)
	if err != nil {
		failed <- User{
			Id:       "",
			Username: *pUsername,
			Password: password,
		}
		return
	}

	created <- User{Id: signUp.GetId(),
		ValidationCode: signUp.VerificationToken,
		Username:       *pUsername,
		Password:       password,
	}
	<-concurrentGoroutines
}

func activateUser(uc pb.UserServiceClient, created chan User, activated chan User, failed chan User) {
	concurrentGoroutines <- struct{}{}
	ctx := context.TODO()
	select {
	case user := <-created:
		rVerification := &pb.VerifyUserRequest{Code: user.ValidationCode}
		_, err := uc.VerifyUser(ctx, rVerification)
		if err != nil {
			failed <- User{
				Id:       "",
				Username: user.Username,
				Password: user.Password,
			}
			return
		}
		activated <- User{Id: user.Id, ValidationCode: user.ValidationCode, Username: user.Username, Password: user.Password}
	}
	<-concurrentGoroutines
}

func tokenUser(authClient pb.AuthServiceClient, activated chan User, failed chan User) {
	concurrentGoroutines <- struct{}{}
	ctx := context.TODO()

	select {
	case user := <-activated:
		credentials := &pb.SignInUserInput{
			Username: user.Username,
			Password: user.Password,
		}
		_, err := authClient.SignInUser(ctx, credentials)
		if err != nil {
			failed <- User{
				Id:       "",
				Username: user.Username,
				Password: user.Password,
			}
		}
	}
	<-concurrentGoroutines
}

func main2() {
	l := logger.NewConsole(false)
	pUsername, _ := generator.RandomString(8)
	username := *pUsername
	email := username + emailDomain
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
