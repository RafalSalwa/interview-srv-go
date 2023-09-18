package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/cmd/tester_service/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/tester_service/internal/workers"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/RafalSalwa/interview-app-srv/pkg/logger"

	"github.com/RafalSalwa/interview-app-srv/internal/generator"
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
	token          *Token
}

type Token struct {
	access  string
	refresh string
}

var (
	maxNbConcurrentGoroutines = 30
	concurrentGoroutines      = make(chan struct{}, maxNbConcurrentGoroutines)
)

func runWorkersInOrder(ctx context.Context, cfg *config.Config, l *logger.Logger) {

	qCreatedUsers := make(chan User, numUsers)
	qActivatedUsers := make(chan User, numUsers)
	qFailedUsers := make(chan User, numUsers)

	done := make(chan bool)

	for i := 1; i <= numUsers; i++ {
		go createUser(cfg, qCreatedUsers, qFailedUsers)
	}
	for i := 1; i <= numUsers; i++ {
		go activateUser(cfg, qCreatedUsers, qActivatedUsers, qFailedUsers)
	}
	for i := 1; i <= numUsers; i++ {
		go tokenUser(cfg, qActivatedUsers, qFailedUsers)
	}

	go func() {
		for {
			fmt.Printf("Concurrent queue len: | %6d | user creation queue:  %6d | user activation queue: %6d \n", len(concurrentGoroutines), len(qCreatedUsers), len(qActivatedUsers))
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

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	l := logger.NewConsole()
	workers.NewDaisyChain(cfg)
	runWorkersInOrder(ctx, cfg, l)
}

func createUser(cfg *config.Config, created chan User, failed chan User) {
	concurrentGoroutines <- struct{}{}

	pUsername, _ := generator.RandomString(12)
	email := *pUsername + emailDomain

	newUser := &models.SignUpUserRequest{
		Email:           email,
		Password:        password,
		PasswordConfirm: password,
	}
	marshalled, err := json.Marshal(newUser)
	if err != nil {
		log.Fatalf("impossible to marshall: %s", err)
		<-concurrentGoroutines
		return
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
	created <- User{
		Username: *pUsername,
		Email:    email,
		Password: password,
	}
	<-concurrentGoroutines
}

func activateUser(cfg *config.Config, created chan User, activated chan User, failed chan User) {
	concurrentGoroutines <- struct{}{}
	//ctx := context.TODO()
	select {
	case user := <-created:
		reqUser := &models.VerificationCodeRequest{Email: user.Email}
		marshalled, err := json.Marshal(reqUser)
		if err != nil {
			log.Fatalf("impossible to marshall: %s", err)
		}
		client := &http.Client{}
		URL := "http://" + cfg.Http.Addr + "/auth/code"
		req, err := http.NewRequest("POST", URL, bytes.NewReader(marshalled))
		req.SetBasicAuth(cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
			<-concurrentGoroutines
			return
		}
		type vCode struct {
			Token string `json:"verification_token"`
		}
		type target struct {
			User vCode `json:"user"`
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
		URL = "http://" + cfg.Http.Addr + "/auth/verify/" + tgt.User.Token
		req, err = http.NewRequest("GET", URL, bytes.NewReader(marshalled))
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

		activated <- User{ValidationCode: tgt.User.Token, Username: user.Username, Password: user.Password}
	}
	<-concurrentGoroutines
}

func tokenUser(cfg *config.Config, activated chan User, failed chan User) {
	concurrentGoroutines <- struct{}{}

	select {
	case user := <-activated:
		credentials := &models.SignInUserRequest{
			Username: user.Username,
			Password: user.Password,
		}
		marshalled, err := json.Marshal(credentials)
		if err != nil {
			log.Fatalf("impossible to marshall: %s", err)
			<-concurrentGoroutines
		}
		client := &http.Client{}
		URL := "http://" + cfg.Http.Addr + "/auth/signin"
		req, err := http.NewRequest("POST", URL, bytes.NewReader(marshalled))
		req.SetBasicAuth(cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password)
		resp, err := client.Do(req)
		_, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("impossible to read all body of response: %s", err)
			<-concurrentGoroutines
		}
	}
	<-concurrentGoroutines
}
