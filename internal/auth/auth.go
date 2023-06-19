package auth

import (
	"handlers"
	"net/http"
)

func Authorization(h handlers.HandlerFunc) http.HandlerFunc {
	at, _ := NewAuthMethod(h, "basic")
	return at.middleware(h)
}
