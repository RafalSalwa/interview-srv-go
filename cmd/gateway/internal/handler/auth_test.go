package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/cqrs"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestSignUpUser(t *testing.T) {
	os.Setenv("APP_ENV", "staging")
	cfg := config.InitConfig()

	service, _ := cqrs.NewCQRSService(cfg.Grpc)
	l := logger.NewConsole()

	signUpHandler := NewAuthHandler(service, l)
	handler := http.HandlerFunc(signUpHandler.SignUpUser())
	req := models.SignUpUserRequest{
		Email:           fmt.Sprintf("interview%d@interview.com", rand.Int63()),
		Password:        "VeryG00dPass!",
		PasswordConfirm: "VeryG00dPass!",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(req)
	assert.NoError(t, err)

	r := httptest.NewRequest(http.MethodPost, "/auth/signup", &buf)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	if status := w.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	// Check the response body is what we expect.
	expected := `{"status":"created"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}

func TestSignInUser(t *testing.T) {
	os.Setenv("APP_ENV", "staging")
	cfg := config.InitConfig()
	service, _ := cqrs.NewCQRSService(cfg.Grpc)
	l := logger.NewConsole()

	authHandler := NewAuthHandler(service, l)
	handler := http.HandlerFunc(authHandler.SignInUser())
	req := models.SignInUserRequest{
		Username: "user1",
		Password: "password",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(req)
	assert.NoError(t, err)

	r := httptest.NewRequest(http.MethodPost, "/auth/signup", &buf)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if !strings.Contains(w.Body.String(), `"token"`) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}
