package workers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/cmd/tester_service/config"
	"github.com/RafalSalwa/interview-app-srv/pkg/generator"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	pb "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	emailDomain = "@interview.com"
	password    = "VeryG00dPass!"
	numUsers    = 50
)

var (
	dcConn, _ = grpc.Dial("0.0.0.0:8082", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	authClient = pb.NewAuthServiceClient(dcConn)
	userClient = pb.NewUserServiceClient(dcConn)
)

type User struct {
	ValidationCode string
	Username       string
	Email          string
	Password       string
}

func NewDaisyChain(ctx context.Context, cfg *config.Config) {
	ch := make(chan User, numUsers+1)

	go Generate(ctx, ch, cfg)
	mid := make(chan User)
	go worker(ctx, ch, mid, "activate")
	out := make(chan User)
	go worker(ctx, mid, out, "token")
}

func Generate(ctx context.Context, ch chan<- User, cfg *config.Config) {
	ch <- dcCreateUser(ctx, cfg)
}

func dcCreateUser(ctx context.Context, cfg *config.Config) User {
	pUsername, _ := generator.RandomString(12)
	email := pUsername + emailDomain

	newUser := &models.SignUpUserRequest{
		Email:           email,
		Password:        password,
		PasswordConfirm: password,
	}
	marshaled, err := json.Marshal(newUser)
	if err != nil {
		log.Fatalf("impossible to marshall: %s", err)
	}
	client := &http.Client{}
	URL := "http://" + cfg.HTTP.Addr + "/auth/signup"
	// pass the values to the request's body
	req, err := http.NewRequestWithContext(ctx, "POST", URL, bytes.NewReader(marshaled))
	if err != nil {
		log.Fatalf("impossible to read all body of response: %s", err)
	}
	req.SetBasicAuth(cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	return User{
		ValidationCode: "",
		Password:       newUser.Password,
	}
}

func worker(ctx context.Context, in <-chan User, out chan<- User, task string) {
	inUser := <-in
	var outUser User
	switch task {
	case "activate":
		outUser = dcActivateUser(ctx, inUser)
	case "token":
		outUser = dcTokenUser(ctx, inUser)
	}
	out <- outUser
}

func dcActivateUser(ctx context.Context, inUser User) User {
	rVerification := &pb.VerifyUserRequest{Code: inUser.ValidationCode}
	_, _ = userClient.VerifyUser(ctx, rVerification)
	return inUser
}

func dcTokenUser(ctx context.Context, inUser User) User {
	credentials := &pb.SignInUserInput{
		Username: inUser.Username,
		Password: inUser.Password,
	}
	_, _ = authClient.SignInUser(ctx, credentials)
	return inUser
}
