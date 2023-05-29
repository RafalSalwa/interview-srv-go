package auth

import (
	apiHandler "github.com/RafalSalwa/interview-app-srv/api/handler"
	"net/http"
)

func Authorization(h apiHandler.HandlerFunc) http.HandlerFunc {
	at, _ := NewAuthMethod(h, "basic")
	return at.middleware(h)
}
