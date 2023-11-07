package workers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/cmd/tester_service/config"
	"github.com/RafalSalwa/interview-app-srv/pkg/generator"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/fatih/color"
	"io"
	"log"
	"net/http"
	"time"
)

type testUser struct {
	ValidationCode string
	Username       string
	Email          string
	Password       string
}

func OneAtTime(ctx context.Context, cfg *config.Config) {
	for {
		pUsername, _ := generator.RandomString(12)
		email := pUsername + emailDomain

		user := testUser{
			Username: pUsername,
			Email:    email,
			Password: password,
		}
		newUser := &models.SignUpUserRequest{
			Email:           user.Email,
			Password:        user.Password,
			PasswordConfirm: user.Password,
		}

		marshaled, err := json.Marshal(newUser)
		if err != nil {
			log.Fatal("impossible to marshall: %+v\n", err)
		}

		client := &http.Client{}
		URL := fmt.Sprintf("http://%s/auth/signup", cfg.HTTP.Addr)
		req, err := http.NewRequestWithContext(ctx, "POST", URL, bytes.NewReader(marshaled))
		if err != nil {
			fmt.Printf("impossible to read all body of response: %s", err)
		}

		req.SetBasicAuth(cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("req do err: ", err)
		}
		if resp.StatusCode != http.StatusCreated {
			fmt.Println("Err.")

			fmt.Printf("    %s req body: %s\n", URL, string(marshaled))
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("impossible to marshall: %s\n", err)
			}
			bodyString := string(bodyBytes)
			fmt.Printf("    %s body: %s", URL, bodyString)
		} else {
			fmt.Println(color.GreenString("OK"))
		}
		resp.Body.Close()

		reqUser := &models.SignInUserRequest{Email: user.Email, Password: user.Password}

		marshaled, err = json.Marshal(reqUser)
		if err != nil {
			fmt.Printf("impossible to marshall: %s\n", err)
		}
		URL = "http://" + cfg.HTTP.Addr + "/auth/code"
		req, err = http.NewRequestWithContext(ctx, "POST", URL, bytes.NewReader(marshaled))
		if err != nil {
			fmt.Printf("impossible to read all body of response: %s\n", err)
		}
		req.SetBasicAuth(cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password)
		fmt.Print("Fetching vCode: ")
		resp, err = client.Do(req)
		if err != nil {
			fmt.Printf("impossible to marshall: %s\n", err)
		}
		if resp.StatusCode != 200 {
			fmt.Println("Err")
			fmt.Printf("%s req body: %s\n", URL, string(marshaled))
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("impossible to marshall: %s\n", err)
			}
			bodyString := string(bodyBytes)
			fmt.Printf("%s body: %s", URL, bodyString)
		} else {
			fmt.Print(color.GreenString("OK "))
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
			return
		}
		fmt.Println(tgt.User.Token)
		resp.Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

		client = &http.Client{}
		URL = "http://" + cfg.HTTP.Addr + "/auth/verify/" + tgt.User.Token
		req, err = http.NewRequestWithContext(ctx, "GET", URL, bytes.NewReader(marshaled))
		if err != nil {
			fmt.Printf("impossible to read all body of response: %s", err)
		}
		req.SetBasicAuth(cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password)

		fmt.Print("Veryfing: ")
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println("/auth/verify/", err)
			return
		}
		resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Println("Err ")
			fmt.Printf("%s req body: %s\n", URL, string(marshaled))
			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyString := string(bodyBytes)
			fmt.Printf("%s body: %s", URL, bodyString)
		} else {
			fmt.Println(color.GreenString("OK "))
		}
		if err != nil {
			fmt.Println("verify err:", err)
			return
		}

		credentials := &models.SignInUserRequest{
			Email:    user.Email,
			Password: user.Password,
		}
		marshaled, err = json.Marshal(credentials)
		if err != nil {
			fmt.Printf("impossible to marshall: %s\n", err)
		}
		URL = "http://" + cfg.HTTP.Addr + "/auth/signin"
		req, err = http.NewRequestWithContext(ctx, "POST", URL, bytes.NewReader(marshaled))
		if err != nil {
			fmt.Printf("impossible to read all body of response: %s", err)
		}

		req.SetBasicAuth(cfg.Auth.BasicAuth.Username, cfg.Auth.BasicAuth.Password)
		fmt.Print("Logging In: ")
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println("Do err: ", err)
		}

		if resp.StatusCode != 200 {
			fmt.Println("Err ")
			fmt.Printf("%s req body: %s\n", URL, string(marshaled))
			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyString := string(bodyBytes)
			fmt.Printf("%s body: %s", URL, bodyString)
		} else {
			fmt.Println(color.GreenString("OK "))
		}

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ReadAll err: ", err)
		}
		resp.Body.Close()
		fmt.Println("Token: ", string(respBody))
		time.Sleep(10 * time.Second)
		fmt.Println()
	}
}
