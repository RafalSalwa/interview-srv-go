package auth

import (
    "github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
    "net/http"
)

func Authorization(h handler.HandlerFunc) http.HandlerFunc {
	at, _ := NewAuthMethod(h, "basic")
	return at.middleware(h)
}
