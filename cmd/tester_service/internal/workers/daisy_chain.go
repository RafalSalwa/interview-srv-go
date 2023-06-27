package workers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/cmd/tester_service/config"
	"github.com/RafalSalwa/interview-app-srv/internal/generator"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var (
	dcConn, _ = grpc.Dial("0.0.0.0:8082", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	authClient = pb.NewAuthServiceClient(dcConn)
	userClient = pb.NewUserServiceClient(dcConn)
)

const (
	emailDomain = "@interview.com"
	password    = "VeryG00dPass!"
	numUsers    = 50
)

type User struct {
	ValidationCode string
	Username       string
	Email          string
	Password       string
	Token          *Token
}

type Token struct {
	access  string
	refresh string
}

func NewDaisyChain(cfg *config.Config) {
	ch := make(chan User, numUsers+1)

	go Generate(ch, cfg)
	mid := make(chan User)
	go worker(ch, mid, "activate")
	out := make(chan User)
	go worker(mid, out, "token")
	ch = out
}

func Generate(ch chan<- User, cfg *config.Config) {
	ch <- dcCreateUser(cfg)
}

func dcCreateUser(cfg *config.Config) User {
	ctx := context.TODO()
	pUsername, _ := generator.RandomString(12)
	email := *pUsername + emailDomain

	newUser := &models.CreateUserRequest{
		Username:        *pUsername,
		Email:           email,
		Password:        password,
		PasswordConfirm: password,
	}
	marshalled, err := json.Marshal(newUser)
	if err != nil {
		log.Fatalf("impossible to marshall: %s", err)
	}
	client := &http.Client{}
	URL := "http://" + cfg.Http.Addr + "/auth/signup"
	//pass the values to the request's body
	req, err := http.NewRequest("POST", URL, bytes.NewReader(marshalled))
	req.SetBasicAuth(cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password)
	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		<-concurrentGoroutines
		return
	}

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
