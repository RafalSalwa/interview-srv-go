package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/RafalSalwa/interview-app-srv/internal/repository"

	apiHandler "github.com/RafalSalwa/interview-app-srv/api/handler"
	apiRouter "github.com/RafalSalwa/interview-app-srv/api/router"
	apiServer "github.com/RafalSalwa/interview-app-srv/api/server"
	"github.com/RafalSalwa/interview-app-srv/internal/services"

	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/sql"

	"github.com/stretchr/testify/assert"
)

var (
	conf   *config.Conf
	server *http.Server
	ctx    context.Context

	userHandler apiHandler.UserHandler
	authHandler apiHandler.IAuthHandler

	userService services.UserSqlService
	authService services.AuthService
)

func runTestServer() {
	ctx = context.TODO()
	conf = config.New()
	l := logger.NewConsole(conf.App.Debug)

	db := sql.NewUsersDB(conf.DB, l)
	ormDB := sql.NewUsersDBGorm(conf.DB, l)
	userRepository := repository.NewUserAdapter(ormDB)
	userService = services.NewMySqlService(db, l)
	authService = services.NewAuthService(ctx, userRepository, l, conf.Token)

	r := apiRouter.NewApiRouter(l, conf.Token)
	userHandler = apiHandler.NewUserHandler(r, userService, l)
	authHandler = apiHandler.NewAuthHandler(r, authService, l)

	apiRouter.RegisterUserRouter(r, userHandler)
	apiRouter.RegisterAuthRouter(r, authHandler)

	server = apiServer.NewServer(conf, r)
	l.Info().Msgf("Starting REST server %v", server.Addr)
	apiServer.Run(server, conf)
}

func Test_get_api_health(t *testing.T) {
	runTestServer()

	t.Run("Health should return 200", func(t *testing.T) {
		parsedUrl, err := url.Parse("http://" + conf.Server.Addr)
		resp, err := http.Get(fmt.Sprintf("%s/health", parsedUrl))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Every request should have correlation id", func(t *testing.T) {
		parsedUrl, err := url.Parse("http://" + conf.Server.Addr)
		resp, err := http.Get(fmt.Sprintf("%s/health", parsedUrl))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		headers := resp.Header
		_, ok := headers["X-Correlation-Id"]
		if !ok {
			t.Fatal("Correlation ID missing")
		}
		_, ok = headers["Access-Control-Allow-Headers"]
		if !ok {
			t.Fatal("CORS Allow Headers missing")
		}

		assert.Equal(t, 200, resp.StatusCode)
	})
}

func Test_get_api_user(t *testing.T) {
	parsedUrl, _ := url.Parse("http://" + conf.Server.Addr)

	t.Run("it should return unauthenticated", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/user/1", parsedUrl))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
	t.Run("it should return validation error while store when request miss required parameters", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/1", parsedUrl), strings.NewReader(""))
		req.Header.Set("Authorization", "Basic aW50ZXJ2aWV3OmludGVydmlldw==")
		client := &http.Client{}
		resp, _ := client.Do(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		respBody, err := io.ReadAll(resp.Body)
		defer func() {
			_ = resp.Body.Close()
		}()
		contentType := resp.Header.Get("Content-Type")
		if contentType != "application/json;charset=utf8" {
			t.Fatal("Wront Content Type, Expected json")
		}
		fmt.Println(string(respBody), http.DetectContentType(respBody))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		//assert.Equal(t, http.StatusOK, resp.C)
		//assert.Equal(t, `{"statusMessage":"validation error"}`, respBody)
	})
}
