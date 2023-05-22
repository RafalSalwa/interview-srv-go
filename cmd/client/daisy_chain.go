package main

import (
	"context"
	"fmt"

	"github.com/RafalSalwa/interview-app-srv/internal/generator"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	dcConn, _ = grpc.Dial("0.0.0.0:8082", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	authClient = pb.NewAuthServiceClient(dcConn)
	userClient = pb.NewUserServiceClient(dcConn)
)

func dcCreateUser() User {
	ctx := context.TODO()

	pUsername, _ := generator.RandomString(12)
	email := *pUsername + emailDomain

	newUser := &pb.SignUpUserInput{
		Name:            *pUsername,
		Email:           email,
		Password:        password,
		PasswordConfirm: password,
	}

	signUp, _ := authClient.SignUpUser(ctx, newUser)

	return User{
		Id:             signUp.GetId(),
		ValidationCode: signUp.VerificationToken,
		Username:       *pUsername,
		Password:       password,
	}
}

func Generate(ch chan<- User) {
	ch <- dcCreateUser()
}

func worker(in <-chan User, out chan<- User, task string) {
	inUser := <-in
	var outUser User
	switch task {
	case "activate":
		outUser = dcActivateUser(inUser)
	case "token":
		outUser = dcTokenUser(inUser)
	}
	out <- outUser
}

func runWorkersInDaisyChain() {
	ch := make(chan User, numUsers+1)

	go Generate(ch)
	mid := make(chan User)
	go worker(ch, mid, "activate")
	out := make(chan User)
	go worker(mid, out, "token")
	ch = out
}

func dcActivateUser(inUser User) User {
	ctx := context.TODO()
	rVerification := &pb.VerifyUserRequest{Code: inUser.ValidationCode}
	_, _ = userClient.VerifyUser(ctx, rVerification)
	return inUser
}

func dcTokenUser(inUser User) User {
	concurrentGoroutines <- struct{}{}
	ctx := context.TODO()

	credentials := &pb.SignInUserInput{
		Username: inUser.Username,
		Password: inUser.Password,
	}
	_, _ = authClient.SignInUser(ctx, credentials)
	fmt.Println(inUser.Id, inUser.Username)
	return inUser
}
