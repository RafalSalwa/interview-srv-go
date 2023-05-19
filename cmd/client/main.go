package main

import (
	"context"
	"fmt"

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
	username, _ := generator.RandomString(8)
	email := *username + emailDomain
	fmt.Println("Client starting")
	ctx := context.TODO()

	conn, err := grpc.Dial("0.0.0.0:8082", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		fmt.Println(err)
	}
	authClient := pb.NewAuthServiceClient(conn)
	userClient := pb.NewUserServiceClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	newUser := &pb.SignUpUserInput{
		Name:            *username,
		Email:           email,
		Password:        password,
		PasswordConfirm: password,
	}
	signUp, err := authClient.SignUpUser(ctx, newUser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("SignUp: %v \n\n", signUp)

	credentials := &pb.SignInUserInput{
		Username: *username,
		Password: password,
	}

	signIn, err := authClient.SignInUser(ctx, credentials)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("SignIn: %v \n\n", signIn)

	id := &pb.GetUserRequest{
		UserId: signUp.Id,
	}
	user, err := userClient.GetUserById(ctx, id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Get user %v \n\n", user)
	//
	// ids := &pb.GetUsersRequest{
	//	Users: []*pb.GetUserRequest{
	//		{UserId: "1"},
	//		{UserId: "2"},
	//		{UserId: "3"},
	//		{UserId: "100"},
	//	},
	//}
	// users, err := userClient.GetUsers(ctx, ids)
	// if err != nil {
	//	fmt.Println(err)
	//}
	// fmt.Printf("%v", users)
}
