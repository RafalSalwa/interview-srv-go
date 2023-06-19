package auth

import (
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
)

func Authorization(h handler.HandlerFunc) http.HandlerFunc {
	at, _ := NewAuthMethod(h, "basic")
	return at.middleware(h)
}
