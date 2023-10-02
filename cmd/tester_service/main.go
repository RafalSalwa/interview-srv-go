package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/RafalSalwa/interview-app-srv/cmd/tester_service/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/tester_service/internal/workers"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"

	"github.com/RafalSalwa/interview-app-srv/pkg/logger"

	"github.com/RafalSalwa/interview-app-srv/pkg/generator"
)

const (
	emailDomain = "@interview.com"
	password    = "VeryG00dPass!"
	numUsers    = 50
)

type testUser struct {
	ValidationCode string
	Username       string
	Email          string
	Password       string
}

var (
	maxNbConcurrentGoroutines = 30
	concurrentGoroutines      = make(chan struct{}, maxNbConcurrentGoroutines)
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	l := logger.NewConsole()
	workers.NewDaisyChain(ctx, cfg)
	runWorkersInOrder(ctx, cfg, l)
}

func runWorkersInOrder(ctx context.Context, cfg *config.Config, l *logger.Logger) {
	qCreatedUsers := make(chan testUser, numUsers)
	qActivatedUsers := make(chan testUser, numUsers)
	qFailedUsers := make(chan testUser, numUsers)

	done := make(chan bool)

	for i := 1; i <= numUsers; i++ {
		go createUser(ctx, cfg, qCreatedUsers, qFailedUsers)
	}
	for i := 1; i <= numUsers; i++ {
		go activateUser(ctx, cfg, qCreatedUsers, qActivatedUsers, qFailedUsers)
	}
	for i := 1; i <= numUsers; i++ {
		go tokenUser(ctx, cfg, qActivatedUsers, qFailedUsers)
	}

	go func() {
		for {
			fmt.Printf("Concurrent queue len: | %6d | testUser creation queue:  %6d | testUser activation queue: %6d \n", len(concurrentGoroutines), len(qCreatedUsers), len(qActivatedUsers))
			if len(concurrentGoroutines) == 0 {
				done <- true
				fmt.Println("Queues depleted, closing")
				break
			}
			time.Sleep(time.Second)
		}
	}()

	<-done
}

func createUser(ctx context.Context, cfg *config.Config, created, failed chan testUser) {
	concurrentGoroutines <- struct{}{}

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
		<-concurrentGoroutines
		return
	}
	created <- testUser{
		Username: pUsername,
		Email:    email,
		Password: password,
	}
	<-concurrentGoroutines
}

func activateUser(ctx context.Context, cfg *config.Config, created chan testUser, activated chan testUser, failed chan testUser) {
	concurrentGoroutines <- struct{}{}
	// ctx := context.TODO()
	user := <-created
	reqUser := &models.VerificationCodeRequest{Email: user.Email}
	marshaled, err := json.Marshal(reqUser)
	if err != nil {
		log.Fatalf("impossible to marshall: %s", err)
	}
	client := &http.Client{}
	URL := "http://" + cfg.HTTP.Addr + "/auth/code"
	req, err := http.NewRequestWithContext(ctx, "POST", URL, bytes.NewReader(marshaled))
	if err != nil {
		log.Fatalf("impossible to read all body of response: %s", err)
	}
	req.SetBasicAuth(cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	type vCode struct {
		Token string `json:"verification_token"`
	}
	type target struct {
		User vCode `json:"testUser"`
	}
	tgt := target{}
	err = json.NewDecoder(resp.Body).Decode(&tgt)
	if err != nil {
		fmt.Println(err)
		<-concurrentGoroutines
		return
	}
	defer resp.Body.Close()
	if err != nil {
		<-concurrentGoroutines
		return
	}

	client = &http.Client{}
	URL = "http://" + cfg.HTTP.Addr + "/auth/verify/" + tgt.User.Token
	req, err = http.NewRequestWithContext(ctx, "GET", URL, bytes.NewReader(marshaled))
	if err != nil {
		log.Fatalf("impossible to read all body of response: %s", err)
	}
	req.SetBasicAuth(cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password)
	resp, err = client.Do(req)
	if err != nil {
		<-concurrentGoroutines
		log.Fatal(err)
		return
	}

	if err != nil {
		<-concurrentGoroutines
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	activated <- testUser{ValidationCode: tgt.User.Token, Username: user.Username, Password: user.Password}
	<-concurrentGoroutines
}

func tokenUser(ctx context.Context, cfg *config.Config, activated, failed chan testUser) {
	concurrentGoroutines <- struct{}{}

	user := <-activated
	credentials := &models.SignInUserRequest{
		Username: user.Username,
		Password: user.Password,
	}
	marshaled, err := json.Marshal(credentials)
	if err != nil {
		log.Fatalf("impossible to marshall: %s", err)
	}
	client := &http.Client{}
	URL := "http://" + cfg.HTTP.Addr + "/auth/signin"
	req, err := http.NewRequestWithContext(ctx, "POST", URL, bytes.NewReader(marshaled))
	if err != nil {
		log.Fatalf("impossible to read all body of response: %s", err)
	}

	req.SetBasicAuth(cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("impossible to read all body of response: %s", err)
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("impossible to read all body of response: %s", err)
	}
	<-concurrentGoroutines
}
